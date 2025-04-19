package controller

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"

	"example.com/review/v2/config"
	"example.com/review/v2/db"
	"example.com/review/v2/db/model"
	"example.com/review/v2/storage"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"gorm.io/gorm"
)

type App struct {
	Config *config.Config
	DB     *gorm.DB
	S3     *s3.Client
}

func New(cfg *config.Config) *App {
	return &App{Config: cfg}
}

func (a *App) JobApiHandler(w http.ResponseWriter, r *http.Request) {
	// process new file
	a.ProcessDump()

	// send response
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Message: "Reviews Processed",
		Status:  "success",
	}
	json.NewEncoder(w).Encode(response)
}

func (a *App) ReviewsApiHandler(w http.ResponseWriter, r *http.Request) {
	// get page number
	page := r.URL.Query().Get("page")
	pageNum := 1
	if page != "" {
		var err error
		pageNum, err = strconv.Atoi(page)
		if err != nil {
			response := Response{
				Message: "Invalid page number",
				Status:  "failed",
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		} else if pageNum <= 0 {
			response := Response{
				Message: "page should be positive",
				Status:  "failed",
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	// get 100 latest reviews at requested page
	pageSize := 100
	offset := (pageNum - 1) * pageSize
	var reviews []model.HotelReview
	a.DB.Order("created_at desc").Offset(offset).Limit(100).Find(&reviews)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reviews)
}

func (a *App) StartServer() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/process/", a.JobApiHandler).Methods("GET")
	r.HandleFunc("/reviews/", a.ReviewsApiHandler).Methods("GET")
	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed: %v\n", err)
	}
}

func (a *App) Run() error {
	// connect to DB
	dbConn, err := db.InitDB(a.Config)
	if err != nil {
		return fmt.Errorf("DB connection failed: %w", err)
	}
	a.DB = dbConn

	// connect to S3
	s3client, err := storage.NewClient(a.Config)
	if err != nil {
		return fmt.Errorf("S3 client failed: %w", err)
	}
	a.S3 = s3client

	// setup api
	go a.StartServer()

	// process reviews now
	go a.ProcessDump()

	// also run job every hour
	timer := time.NewTicker(time.Hour)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			go a.ProcessDump()
		default:
			log.Println("heart beat")
			time.Sleep(10 * time.Second)
		}
	}
}

func (a *App) ProcessDump() {
	// fetch review file
	filepath, err := storage.DownloadFile(a.S3, a.Config)
	if err != nil {
		log.Printf("S3 download failed: %v", err)
		return
	}

	log.Printf("downloaded data in file: %v", filepath)

	// read from file and print
	inputFile, err := os.Open(filepath)
	if err != nil {
		log.Printf("Unable to open file %v: %v", filepath, err)
		return
	}
	defer inputFile.Close()

	var wg sync.WaitGroup

	// buffered scanner to read jl file line by line
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()

		wg.Add(1)
		go a.ProcessReview(line, &wg)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	// wait for all goroutines
	wg.Wait()
}

func (a *App) ProcessReview(line string, wg *sync.WaitGroup) {
	defer wg.Done()

	// parse json
	var review model.HotelReview
	err := json.Unmarshal([]byte(line), &review)
	if err != nil {
		log.Print("Error unmashalling JSON:", err)
		return
	}

	// validate input
	if err := review.Validate(); err != nil {
		fmt.Print("validation failed:", err)
		return
	}

	// check duplicate
	duplicate, err := review.CheckDuplicate(a.DB)
	if err != nil {
		fmt.Printf("DB error: %v", err)
		return
	}
	if !duplicate {
		// insert new record
		if err := a.DB.Create(&review).Error; err != nil {
			fmt.Printf("failed to insert HotelReview: %v", err)
			return
		}
	} else {
		fmt.Printf("skipped duplicate review\n")
	}
}
