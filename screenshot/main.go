package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	var url string
	var file string

	// Prepare a form that you will submit to that URL.
	var buff bytes.Buffer
	w := multipart.NewWriter(&buff)

	// Add your image file
	f, err := os.Open("https://maps.googleapis.com/maps/api/staticmap?size=2480x720&key=AIzaSyCeLaDxdBF1ipP-rADJZhitd15MOZoC-1I&zoom=13&maptype=roadmap&center=23.770474,90.362925&markers=color:green%7C23.770474,90.362925")
	if err != nil {
		fmt.Println("file open error")
		return
	}

	fw, err := w.CreateFormFile("image", file)
	if err != nil {
		fmt.Println("cann't create file")

		return
	}
	if _, err = io.Copy(fw, f); err != nil {
		return
	}
	// Add the other fields
	if fw, err = w.CreateFormField("key"); err != nil {
		return
	}
	if _, err = fw.Write([]byte("KEY")); err != nil {
		return
	}
	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	w.Close()

	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", url, &buff)
	if err != nil {
		return
	}

	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())

	// Submit the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}

	// Check the response
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
	}
	return
}
