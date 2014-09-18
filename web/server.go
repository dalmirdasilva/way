package web

import (
  "fmt"
  "net/http"
  "github.com/dalmirdasilva/way/lib"
  "strconv"
)

func paramFrom(r *http.Request, name string) []string {
  query := r.URL.Query()
  return query[name]
}

func wayHandler(w http.ResponseWriter, r *http.Request) {
  url := r.URL.Path[1:]
  exists := lib.WayExists(url)
  fmt.Fprintf(w, "Is the a way? %t!", exists)
}

func feedbackHandler(w http.ResponseWriter, r *http.Request) {
  url := paramFrom(r, "url")[0]
  weight := paramFrom(r, "weight")[0]
  f64, _ := strconv.ParseFloat(weight, 64)
  ok := lib.AddNewFeedback(url, f64)
  fmt.Fprintf(w, "Feedback %v added to %v (%v). Thanks.", weight, url, ok)
}

func RunServer() {
  http.HandleFunc("/way", wayHandler)
  http.HandleFunc("/feedback", feedbackHandler)
  http.ListenAndServe(":8080", nil)
}
