package controller

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"example.com/review/v2/config"
	"example.com/review/v2/db"
	"example.com/review/v2/db/model"
	"example.com/review/v2/storage"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"gorm.io/gorm"
)

type App struct {
	Config *config.Config
}

func New(cfg *config.Config) *App {
	return &App{Config: cfg}
}

func (a *App) Run() error {
	// connect to DB
	dbConn, err := db.InitDB(a.Config)
	if err != nil {
		return fmt.Errorf("DB connection failed: %w", err)
	}

	// connect to S3
	s3client, err := storage.NewClient(a.Config)
	if err != nil {
		return fmt.Errorf("S3 client failed: %w", err)
	}

	_ = dbConn

	// process reviews now
	go a.ProcessDump(s3client, dbConn)

	// also run job every hour
	timer := time.NewTicker(time.Hour)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			go a.ProcessDump(s3client, dbConn)
		default:
			log.Println("heart beat")
			time.Sleep(2 * time.Second)
		}
	}
}

func (a *App) ProcessDump(s3client *s3.Client, dbConn *gorm.DB) {
	// fetch review file
	filepath, err := storage.DownloadFile(s3client, a.Config)
	if err != nil {
		log.Printf("S3 download failed: %v", err)
		return
	}

	log.Printf("Got data in file: %v", filepath)

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
		go a.ProcessReview(line, dbConn, &wg)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	log.Printf("ProcessDump done")

	// wait for all goroutines
	wg.Wait()
}

func (a *App) ProcessReview(line string, dbConn *gorm.DB, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Printf("ProcessReview: %s\n", line)
	// parse json
	var review model.HotelReview
	err := json.Unmarshal([]byte(line), &review)
	if err != nil {
		log.Print("Error unmashalling JSON:", err)
		return
	}

	log.Println("Unmashalling completed")
	log.Println(review)

	jsonBytes, err := json.MarshalIndent(review, "", "  ")
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return
	}
	fmt.Println(string(jsonBytes))

	// validate input
	if err := review.Validate(); err != nil {
		fmt.Print("validation failed:", err)
		return
	}

	// check duplicate
	duplicate, err := review.CheckDuplicate(dbConn)
	if err != nil {
		fmt.Printf("DB error: %v", err)
		return
	}
	if !duplicate {
		// insert new record
		if err := dbConn.Create(&review).Error; err != nil {
			fmt.Printf("failed to insert HotelReview: %v", err)
			return
		}
	}
}
