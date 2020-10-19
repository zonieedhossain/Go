package main

import (
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
	fmt.Fprintf(w, "upload\n")
	//parshe every single part of request type Multipart/FromFile
	r.ParseMultipartForm(10 << 20)

	//take the file from input
	image, handler, err := r.FormFile("image")
	if err != nil {
		fmt.Println("something Wrong Please check input")
		fmt.Println(err)
		return
	}
	defer image.Close()
	fmt.Printf("Upload Image: %+v\n Image Size: %+v\n MIME Header: %+v\n", handler.Filename, handler.Size, handler.Header)

	err = GandU(image, "image")
	if err != nil {
		log.Fatal(err)
		return
	}
	// create a temp derectory and file
	// tempFile, err := ioutil.TempFile("temp-image", "upload-*png")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer tempFile.Close()
	//
	// imagebytes, err := ioutil.ReadAll(image)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// tempFile.Write(imagebytes)

	//if success to upload then return
	fmt.Fprintf(w, "successfully uploaded")
}

func GandU(file multipart.File, image string) error {
	url, err := GenSignedurl(fmt.Sprintf("demo/%s", image))
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
func Routes() {
	http.HandleFunc("/upload", Uploadfile)
	http.ListenAndServe(":8081", nil)
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

	Routes()
}
