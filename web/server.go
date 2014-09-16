package web

import (
  "fmt"
  "net/http"
  "github.com/dalmirdasilva/way/lib"
)

func paramFrom(r *http.Request, name string) interface{} {
  query := r.URL.Query()
  return query[name]
}

func wayHandler(w http.ResponseWriter, r *http.Request) {
  url := r.URL.Path[1:]
  exists := lib.WayExists(url)
  fmt.Fprintf(w, "Is the a way? %t!", exists)
}

func feedbackHandler(w http.ResponseWriter, r *http.Request) {
  url := paramFrom(r, "url")
  weight := paramFrom(r, "weight")
  fmt.Fprintf(w, "Feedback %v added to %v. Thanks.", weight, url)
}

func InitWayServer() {
  http.HandleFunc("/way", wayHandler)
  http.HandleFunc("/feedback", feedbackHandler)
  http.ListenAndServe(":8080", nil)
}
