package config

import (
	"errors"
	"os"
)

type Config struct {
	// MinIO/S3
	Provider  string
	Endpoint  string
	AccessKey string
	SecretKey string
	Region    string
	Bucket    string
	Key       string

	// Mysql
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

func Load() (*Config, error) {
	cfg := &Config{
		// S3/MinIO
		Provider:  os.Getenv("STORAGE_PROVIDER"),
		Endpoint:  os.Getenv("S3_ENDPOINT"),
		AccessKey: os.Getenv("AWS_ACCESS_KEY_ID"),
		SecretKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
		Region:    os.Getenv("S3_REGION"),
		Bucket:    os.Getenv("S3_BUCKET"),

		// DB
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
	}

	if cfg.Provider != "aws" && cfg.Provider != "minio" {
		return nil, errors.New("STORAGE_PROVIDER must be 'aws' or 'minio'")
	}

	if cfg.Bucket == "" || cfg.Region == "" {
		return nil, errors.New("S3_BUCKET and S3_REGION must be set")
	}

	if cfg.Provider == "minio" && cfg.Endpoint == "" {
		return nil, errors.New("S3_ENDPOINT must be set for MinIO")
	}

	return cfg, nil
}
