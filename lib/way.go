package lib

import "fmt"

func WayExists(url string) bool {
  return true
}

func AddNewFeedback(url string, weight float64) bool {
  nodes := ParseUrlIntoNodes(url)
  fmt.Println(nodes)
  return true
}
