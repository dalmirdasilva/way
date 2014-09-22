package lib

func WayExists(url string) bool {
  exists, _ := TryGenerateAcceptableNodeChain(url)
  if !exists {
    ProcessDelayedJob(func(url interface{}) {
      GenerateNodeChain(url.(string))
    }, url)
  }
  return exists
}

func AddNewFeedback(url string, weight float64) bool {
  /*nodes := GenerateNodeChain(url)
  UpdateFeedback(nodes[len(nodes) - 1])
  PropagateFeedback(nodes)*/
  return true
}
