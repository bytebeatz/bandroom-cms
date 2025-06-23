package storage

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"time"

	"cloud.google.com/go/storage"
	"github.com/bytebeatz/bandroom-cms/config"
)

// UploadFile uploads a file to the configured GCS bucket.
func UploadFile(
	ctx context.Context,
	file multipart.File,
	fileHeader *multipart.FileHeader,
	destPath string,
) (string, error) {
	client := config.GCSClient
	bucketName := config.AppConfig.GCSBucket
	objectName := destPath

	wc := client.Bucket(bucketName).Object(objectName).NewWriter(ctx)
	wc.ContentType = fileHeader.Header.Get("Content-Type")
	wc.CacheControl = "public, max-age=86400"
	wc.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}

	_, err := io.Copy(wc, file)
	if err != nil {
		return "", fmt.Errorf("upload failed: %w", err)
	}

	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("finalizing upload failed: %w", err)
	}

	url := fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, objectName)
	return url, nil
}

// DownloadFile streams a file from GCS.
func DownloadFile(ctx context.Context, objectPath string) (io.ReadCloser, error) {
	client := config.GCSClient
	bucket := config.AppConfig.GCSBucket

	reader, err := client.Bucket(bucket).Object(objectPath).NewReader(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to download file: %w", err)
	}
	return reader, nil
}

// DeleteFile deletes an object from GCS.
func DeleteFile(ctx context.Context, objectPath string) error {
	client := config.GCSClient
	bucket := config.AppConfig.GCSBucket

	obj := client.Bucket(bucket).Object(objectPath)
	err := obj.Delete(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}
	log.Printf("Deleted file: %s", objectPath)
	return nil
}

// GenerateSignedURL returns a pre-signed public URL (optional).
func GenerateSignedURL(
	ctx context.Context,
	objectPath string,
	expiresIn time.Duration,
) (string, error) {
	url, err := storage.SignedURL(
		config.AppConfig.GCSBucket,
		objectPath,
		&storage.SignedURLOptions{
			Method:         "GET",
			Expires:        time.Now().Add(expiresIn),
			GoogleAccessID: "YOUR_SERVICE_ACCOUNT_EMAIL", // Set via env or viper
			PrivateKey:     []byte("YOUR_PRIVATE_KEY"),   // Load from secret manager if needed
		},
	)
	if err != nil {
		return "", fmt.Errorf("failed to generate signed URL: %w", err)
	}
	return url, nil
}
