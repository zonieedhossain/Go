package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func GetMapWithPolygon(output string, geometry string) {
	endpoint, _ := url.Parse("https://maps.googleapis.com/maps/api/staticmap?center=Dhaka&zoom=13&size=2480x720&key=")
	queryParams := endpoint.Query()
	queryParams.Set("ppi", "320")
	queryParams.Set("w", "1280")
	queryParams.Set("h", "720")
	queryParams.Set("z", "1")
	queryParams.Set("a0", geometry)
	endpoint.RawQuery = queryParams.Encode()
	response, err := http.Get(endpoint.String())
	if err != nil {
		fmt.Printf("the http request faild %s\n", err)
	} else {
		f, err := os.Create(output)
		if err != nil {
			fmt.Printf("the http request faild %s\n", err)

			return
		}
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("the http request faild %s\n", err)

			return
		}
		f.Write(data)
		f.Close()
	}
}
func main() {
	fmt.Println("Starting coding.....")
	GetMapWithPolygon("map.jpeg", "23.879519, 90.396992,23.878037, 90.397432,23.878067, 90.398263")
}
