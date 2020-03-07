package main

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)

type User struct {
	Id       int
	Username string
}

type Point struct {
	lat float64
	lon float64
}

func main() {

	data := [][]string{
		[]string{"1", "Adam"},
		[]string{"2", "Eve"},
	}

	pointdata := [][]string{
		[]string{"23.3", "33.3"},
		[]string{"23.3", "33.3"},
	}

	// convert data string slice to struct
	// such as
	//user1 := &User{"1","Adam"}
	//user2 := &User{"2","Eve"}

	users := []*User{}
	points := []*Point{}

	for _, v := range data {
		//fmt.Println("data: ", v[0], v[1])

		// convert v[0] to type integer
		id, err := strconv.Atoi(v[0])
		if err != nil {
			log.Fatal(err, v[0])
		}
		user := &User{id, v[1]}
		users = append(users, user)
	}
	for _, v := range pointdata {
		//fmt.Println("data: ", v[0], v[1])

		// convert v[0] to type integer
		lat, err := strconv.ParseFloat(v[0], 64)
		if err != nil {
			log.Fatal(err, v[0])
		}
		lon, err := strconv.ParseFloat(v[1], 64)
		if err != nil {
			log.Fatal(err, v[0])
		}
		point := &Point{lat, lon}
		points = append(points, point)
	}
	// the reflect way
	for i := 0; i < 2; i++ {
		u := reflect.ValueOf(users[i])
		id := reflect.Indirect(u).FieldByName("Id")
		name := reflect.Indirect(u).FieldByName("Username")

		fmt.Println(id, name)
		fmt.Println("==========")
	}
	for i := 0; i < 2; i++ {
		u := reflect.ValueOf(points[i])
		lat := reflect.Indirect(u).FieldByName("lat")
		lon := reflect.Indirect(u).FieldByName("lon")

		fmt.Println(lat, lon)
	}
}
