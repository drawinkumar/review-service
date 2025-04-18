package storage

import (
	"context"
	"fmt"
	"os"
	"time"

	"example.com/review/v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func DownloadFile(client *s3.Client, cfg *config.Config) (string, error) {
	reviewFilePath := time.Now().Format("2006_01_02.jl")
	localFilePath := fmt.Sprintf("./tmp/%s", reviewFilePath)

	// get s3 object
	resp, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: &cfg.Bucket,
		Key:    &reviewFilePath,
	})
	if err != nil {
		fmt.Println("Bucket: ", cfg.Bucket)
		return "", fmt.Errorf("S3 get failure for %s: %w", reviewFilePath, err)
	}
	defer resp.Body.Close()

	// create tmp file
	localFile, err := os.Create(localFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %v, %v", localFilePath, err)
	}
	defer localFile.Close()

	// copy data
	_, err = localFile.ReadFrom(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to copy data to file %v, %v", localFilePath, err)
	}
	return localFilePath, nil
}
