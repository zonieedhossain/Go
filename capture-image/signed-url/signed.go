package signed_url

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"golang.org/x/oauth2/google"
)

// GenSignedurl will generate signed url for uploading image
func GenSignedurl(name string) (string, error) {
	buck := os.Getenv("BUCKET_NAME")
	sapath := os.Getenv("SERVICE_ACCOUNT_PATH")
	return generateSignedUrl(sapath, buck, "PUT", name)
}

func GenerateSignedGetUrl(name string) (string, error) {
	buck := os.Getenv("BUCKET_NAME")
	sapath := os.Getenv("SERVICE_ACCOUNT_PATH")
	return generateSignedUrl(sapath, buck, "GET", name)
}

func generateSignedUrl(sapath, buck, method, name string) (string, error) {
	sakey, err := ioutil.ReadFile(sapath)
	if err != nil {
		return "", err
	}

	cfg, err := google.JWTConfigFromJSON(sakey)
	if err != nil {
		return "", err
	}
	fmt.Println("-----------------", cfg)
	var contenttype string
	if method != "GET" {
		contenttype = "image/jpeg"
	}
	fmt.Println("-----------------", cfg)
	url, err := storage.SignedURL(buck, name, &storage.SignedURLOptions{
		GoogleAccessID: cfg.Email,
		PrivateKey:     cfg.PrivateKey,
		ContentType:    contenttype,
		Method:         method,
		Expires:        time.Now().Add(10 * time.Minute),
	})
	if err != nil {
		return "", err
	}
	fmt.Println(url)
	return url, nil
}
