package main

import (
	"encoding/json"
	"net/http"
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
	resp := Message(true, "ok")
	resp["Result"] = a
	Respond(w, resp)
}
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
