package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func main() {
	bucket := "hello-consul"
	err := listFiles(bucket)
	if err != nil {
		fmt.Println(err)
	}
}

type Details struct {
	imageName string
	imageTime time.Time
}

func listFiles(bucket string) error {
	// bucket := "bucket-name"
	sapath := os.Getenv("SERVICE_ACCOUNT_PATH")
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(sapath))
	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	it := client.Bucket(bucket).Objects(ctx, &storage.Query{
		Prefix: "route/driver/42/",
	})

	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("Bucket(%q).Objects: %v", bucket, err)
		}
		fmt.Println(attrs.Name)

	}

	return nil
}

// GenSignedurl will generate signed url for uploading image
func GenSignedurl(name, contentType string) (string, error) {
	buck := os.Getenv("BUCKET_NAME")
	sapath := os.Getenv("SERVICE_ACCOUNT_PATH")
	return generateSignedUrl(sapath, buck, contentType, "PUT", name)
}

func GenerateSignedGetUrl(name string) (string, error) {
	buck := os.Getenv("BUCKET_NAME")
	sapath := os.Getenv("SERVICE_ACCOUNT_PATH")
	return generateSignedUrl(sapath, buck, "", "GET", name)
}

func generateSignedUrl(sapath, buck, contentType, method, name string) (string, error) {
	sakey, err := ioutil.ReadFile(sapath)
	if err != nil {
		return "", err
	}

	cfg, err := google.JWTConfigFromJSON(sakey)
	if err != nil {
		return "", err
	}

	url, err := storage.SignedURL(buck, name, &storage.SignedURLOptions{
		GoogleAccessID: cfg.Email,
		PrivateKey:     cfg.PrivateKey,
		ContentType:    contentType,
		Method:         method,
		Expires:        time.Now().Add(10 * time.Minute),
	})
	if err != nil {
		return "", err
	}
	fmt.Println(url)
	return url, nil
}
