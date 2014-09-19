package lib

import "fmt"

func WayExists(url string) bool {
  exists, _ := TryGenerateAcceptableNodeChain(url)
  if !exists {
    go func(url string) {
      GenerateNodeChain(url)
    }(url)
  }
  return exists
}

func AddNewFeedback(url string, weight float64) bool {
  nodes := GenerateNodeChain(url)
  UpdateFeedback(nodes[len(nodes) - 1])
  PropagateFeedback(nodes)
  return true
}
