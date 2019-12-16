package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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

	// create a temp derectory and file
	tempFile, err := ioutil.TempFile("temp-image", "upload-*png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tempFile.Close()

	imagebytes, err := ioutil.ReadAll(image)
	if err != nil {
		fmt.Println(err)
		return
	}
	tempFile.Write(imagebytes)

	//if success to upload then return
	fmt.Fprintf(w, "successfully uploaded")
}
func Routes() {
	http.HandleFunc("/upload", Uploadfile)
	http.ListenAndServe(":8081", nil)
}
func main() {
	fmt.Println("starting upload file")
	Routes()
}
