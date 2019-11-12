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
		dist := distance(new_lat, new_lon, prev_lat, prev_lon)
		fmt.Println(dist)
		prev_lat = new_lat
		prev_lon = new_lon
	}
}
func distance(lat1 float64, lon1 float64, lat2 float64, lon2 float64) float64 {
	const PI float64 = 3.141592653589793
	if (lat2 == 0.0 && lon2 == 0.0) || (lat1 == lat2 && lon1 == lon2) {
		dist := 0.0
		return dist
	} else {
		rlat1 := float64(PI * lat1 / 180)
		rlat2 := float64(PI * lat2 / 180)

		t := float64(lon1 - lon2)
		rt := float64(PI * t / 180)

		dist := math.Sin(rlat1)*math.Sin(rlat2) + math.Cos(rlat1)*math.Cos(rlat2)*math.Cos(rt)

		if dist > 1 {
			dist = 1
		}

		dist = math.Acos(dist)
		dist = dist * 180 / PI
		dist = dist * 60 * 1.1515
		dist = dist * 1.609344
		return dist
	}

}
