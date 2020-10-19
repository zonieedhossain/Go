package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
s
	"cloud.google.com/go/storage"
	"github.com/prometheus/common/log"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/iam/v1"
)

var (
	// iamService is a client for calling the signBlob API.
	iamService *iam.Service

	// serviceAccountName represents Service Account Name.
	// See more details: https://cloud.google.com/iam/docs/service-accounts
	serviceAccountName string

	// serviceAccountID follows the below format.
	// "projects/%s/serviceAccounts/%s"
	serviceAccountID string

	// uploadableBucket is the destination bucket.
	// All users will upload files directly to this bucket by using generated Signed URL.
	uploadableBucket string
)

func signHandler(w http.ResponseWriter, r *http.Request) {
	// Accepts only POST method.
	// Otherwise, this handler returns 405.
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Only POST is supported", http.StatusMethodNotAllowed)
		return
	}

	ct := r.FormValue("content_type")
	if ct == "" {
		http.Error(w, "content_type must be set", http.StatusBadRequest)
		return
	}

	// Generates an object key for use in new Cloud Storage Object.
	// It's not duplicate with any object keys because of UUID.
	key := uuid.New().String()
	if ext := r.FormValue("ext"); ext != "" {
		key += fmt.Sprintf(".%s", ext)
	}

	// Generates a signed URL for use in the PUT request to GCS.
	// Generated URL should be expired after 15 mins.
	url, err := storage.SignedURL(uploadableBucket, key, &storage.SignedURLOptions{
		GoogleAccessID: serviceAccountName,
		Method:         "PUT",
		Expires:        time.Now().Add(15 * time.Minute),
		ContentType:    ct,
		// To avoid management for private key, use SignBytes instead of PrivateKey.
		// In this example, we are using the `iam.serviceAccounts.signBlob` API for signing bytes.
		// If you hope to avoid API call for signing bytes every time,
		// you can use self hosted private key and pass it in Privatekey.
		SignBytes: func(b []byte) ([]byte, error) {
			resp, err := iamService.Projects.ServiceAccounts.SignBlob(
				serviceAccountID,
				&iam.SignBlobRequest{BytesToSign: base64.StdEncoding.EncodeToString(b)},
			).Context(r.Context()).Do()
			if err != nil {
				return nil, err
			}
			return base64.StdEncoding.DecodeString(resp.Signature)
		},
	})
	if err != nil {
		log.Printf("sign: failed to sign, err = %v\n", err)
		http.Error(w, "failed to sign by internal server error", http.StatusInternalServerError)
		return
	}
	GetMapWithPolygon("/home/torque/go/src/BasicGo/capture-image/map.jpeg", "")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, url)
}

func main() {
	cred, err := google.DefaultClient(context.Background(), iam.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}
	iamService, err = iam.New(cred)
	if err != nil {
		log.Fatal(err)
	}

	uploadableBucket = os.Getenv("UPLOADABLE_BUCKET")
	serviceAccountName = os.Getenv("SERVICE_ACCOUNT")
	serviceAccountID = fmt.Sprintf(
		"projects/%s/serviceAccounts/%s",
		os.Getenv("GOOGLE_CLOUD_PROJECT"),
		serviceAccountName,
	)

	http.HandleFunc("/sign", signHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil))
}
func GetMapWithPolygon(output string, v string) {
	endpoint, _ := url.Parse("https://maps.googleapis.com/maps/api/staticmap?center=mohammadpur&zoom=13&size=600x300&maptype=roadmap&markers=color:green%7C23.770474,90.362925&key=API_KEY")
	queryParams := endpoint.Query()

	endpoint.RawQuery = queryParams.Encode()
	response, err := http.Get(endpoint.String())

	if err != nil {
		fmt.Printf("the http request faild %s\n", err)
	} else {
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("the http request faild %s\n", err)
			return
		}

		// body := new(bytes.Buffer)
		// writer := multipart.NewWriter(body)
		// part, err := writer.CreateFormFile("paramName", "map")
		// if err != nil {
		// 	fmt.Println("----------------------")
		// 	return
		// }
		// part.Write(data)
		// var params map[string]string
		// for key, val := range params {
		// 	_ = writer.WriteField(key, val)
		// }
		// err = writer.Close()
		// if err != nil {
		// 	fmt.Println("----------------------")
		// 	return
		// }

		url, err := signedurl.GenSignedurl(fmt.Sprintf("campaign/%s/driver/%s/%s", cid, did, image))
		if err != nil {
			return
		}
		
		req, err := http.NewRequest(http.MethodPut, url, bytes.Reader(data))
		if err != nil {
			return
		}
		req.Header.Add("Content-Type", "image/jpeg")
		log.Info(req)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Errorf("--------------")
			return
		}
		if int(resp.StatusCode/100) != int(2) {
			fmt.Errorf("something went wrong. expect 2XX, but got %v", resp.StatusCode)
			return
		}

	}
	return
}
