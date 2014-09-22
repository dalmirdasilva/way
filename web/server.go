package web

import (
  "fmt"
  "net/http"
  "github.com/dalmirdasilva/way/lib"
  "strconv"
)

func RunServer() {
  http.HandleFunc("/way", wayHandler)
  http.HandleFunc("/feedback", feedbackHandler)
  http.ListenAndServe(":8080", nil)
}

func wayHandler(w http.ResponseWriter, r *http.Request) {
  url := r.URL.Path[1:]
  url = r.FormValue("url")
  fmt.Println(url)
  return
  exists := lib.WayExists(url)
  h := w.Header()
  h.Set("Content-Type", "application/json; charset=utf-8")
  fmt.Fprintf(w, "{\"exists\": \"%t\"}", exists)
}

func feedbackHandler(w http.ResponseWriter, r *http.Request) {
  url := paramFrom(r, "url")[0]
  weight := paramFrom(r, "weight")[0]
  f64, _ := strconv.ParseFloat(weight, 64)
  ok := lib.AddNewFeedback(url, f64)
  fmt.Fprintf(w, "Feedback %v added to %v (%v). Thanks.", weight, url, ok)
}

func paramFrom(r *http.Request, name string) []string {
  query := r.URL.Query()
  return query[name]
}
