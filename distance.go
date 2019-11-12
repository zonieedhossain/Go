package main

import (
	"fmt"
	"math"
)

func main() {
	var new_lat, new_lon float64
	prev_lat := 0.0
	prev_lon := 0.0
	for {
		fmt.Scanln(&new_lat)
		fmt.Scanln(&new_lon)
		if prev_lat == 0.0 && prev_lon == 0.0 {
			dist := 0.0
			fmt.Println(dist)

		} else {
			dist := distance(new_lat, new_lon, prev_lat, prev_lon)
			fmt.Println(dist)
		}
		prev_lat = new_lat
		prev_lon = new_lon
	}
}
func distance(lat1 float64, lon1 float64, lat2 float64, lon2 float64) float64 {
	const PI float64 = 3.141592653589793

	radlat1 := float64(PI * lat1 / 180)
	radlat2 := float64(PI * lat2 / 180)

	theta := float64(lon1 - lon2)
	radtheta := float64(PI * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.1515

	if lat1 == 0.0 && lon1 == 0.0 {
		dist = 0.0
	} else {
		dist = dist * 1.609344
	}

	return dist
}
