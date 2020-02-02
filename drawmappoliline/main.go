package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Point struct {
	Latitude  float64
	Longitude float64
}
type Polygon struct {
	Points []Point
}

func (point *Point) toString() string {
	return fmt.Sprintf("%f,%f", point.Latitude, point.Longitude)
}
func (polygon *Polygon) toString() string {
	var result string
	var points []string
	for _, point := range polygon.Points {
		points = append(points, point.toString())
	}
	result = strings.Join(points, ",")
	return result
}
func GetMapWithPolygon(output string, geometry string) {
	endpoint, _ := url.Parse("https://image.maps.api.here.com/mia/1.6/region")
	fmt.Println(endpoint)
	queryParams := endpoint.Query()
	fmt.Println(queryParams)
	queryParams.Set("api_key", "Uf8AolmmpEuzzgUDCt3V_Kzu7SetGaL9yLxXZlha-RE")

	queryParams.Set("ppi", "320")
	queryParams.Set("w", "1280")
	queryParams.Set("h", "720")
	queryParams.Set("z", "11")
	queryParams.Set("a0", geometry)
	fmt.Println(queryParams)
	endpoint.RawQuery = queryParams.Encode()
	fmt.Println(queryParams)
	response, err := http.Get(endpoint.String())
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		f, _ := os.Create(output)
		fmt.Println(f)
		data, _ := ioutil.ReadAll(response.Body)
		f.Write(data)
		fmt.Println(data)

		defer f.Close()
	}
}
func main() {
	polygon := Polygon{
		Points: []Point{
			Point{Latitude: 37.7397, Longitude: -121.4252},
			Point{Latitude: 37.7974, Longitude: -121.2161},
			Point{Latitude: 37.6391, Longitude: -120.9969},
			Point{Latitude: 37.7397, Longitude: -121.4252},
		},
	}
	GetMapWithPolygon("map.jpeg", polygon.toString())
}
