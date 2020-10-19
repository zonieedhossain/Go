package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"time"

	"cloud.google.com/go/storage"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
)

var (
	serviceAccountName string
	serviceAccountID   string
	uploadableBucket   string
	cfg                *jwt.Config
)

func Uploadfile(w http.ResponseWriter, r *http.Request) {
	file, err := op.Open("/home/xtreme/go/src/test/upload/test.txt")
	if err != nil {
		fmt.Println("Cann't open the file")
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Cann't read the file")
	}
	fi, err := file.Stat()
	if err != nil {
		fmt.Println("Cann't stat the file")
	}
	file.Close()
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, fi.Name())
	if err != nil {
		fmt.Println("Cann't create the part")
	}
	part.Write(fileContents)

	err = writer.Close()
	if err != nil {
		fmt.Println("writter close")
	}

	err = GandU(image, "image")
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Fprintf(w, "successfully uploaded")
}

func GandU(file multipart.File, image string) error {
	url, err := GenSignedurl(fmt.Sprintf("ping/%s", image))
	if err != nil {
		return err
	}
	err = uploadIntoBucket(file, url)
	if err != nil {
		return err
	}
	return nil
}

func uploadIntoBucket(file multipart.File, url string) error {
	req, err := http.NewRequest(http.MethodPut, url, file)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "image/jpeg")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	fmt.Println(resp)
	return nil
}

func GenSignedurl(name string) (string, error) {
	jsonkey, err := ioutil.ReadFile("json_bucket_url")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", jsonkey)
	cfg, err = google.JWTConfigFromJSON(jsonkey, "https://oauth2.googleapis.com/token")
	if err != nil {
		panic(err)
	}
	uploadableBucket = "hello-consul"
	serviceAccountName = cfg.Email
	serviceAccountID = fmt.Sprintf(
		"projects/%s/serviceAccounts/%s",
		"project",
		serviceAccountName,
	)
	url, err := storage.SignedURL(uploadableBucket, name, &storage.SignedURLOptions{
		GoogleAccessID: serviceAccountName,
		Method:         "PUT",
		Expires:        time.Now().Add(15 * time.Minute),
		ContentType:    "image/jpeg",
		PrivateKey:     cfg.PrivateKey,
	})
	if err != nil {

		return "", err
	}
	return url, err
}
func main() {
	fmt.Println("starting upload file")
	Uploadfile()
}
