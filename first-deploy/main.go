package main

import (
	"net/http"
	"test/BasicGo/deploy/api"
)

func main() {
	Route()
}
func Route() {
	http.HandleFunc("/deploy", FirstDeploy)
	http.ListenAndServe(":8080", nil)
}

func FirstDeploy(w http.ResponseWriter, r *http.Request) {
	a := "This is my first deploy"
	resp := api.Message(true, "ok")
	resp["Result"] = a
	api.Respond(w, resp)
}
