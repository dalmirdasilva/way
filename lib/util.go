package lib

import (
  "github.com/dalmirdasilva/way/model"
  "github.com/dalmirdasilva/gorecord/persistence"
  "strings"
)

const LOCAL_FEEDBACK_AVG_THRESHOLD = 0.5
const CUMULATIVE_FEEDBACK_AVG_THRESHOLD = 0.5

func IsNodeAcceptable(node model.Node) bool {
  if node.ChildrenCount > 1 && (node.LocalFeedbackAverage < LOCAL_FEEDBACK_AVG_THRESHOLD || node.CumulativeFeedbackAverage < CUMULATIVE_FEEDBACK_AVG_THRESHOLD){
    return false
  }
  return true
}

func TryGenerateAcceptableNodeChain(url string) (bool, []model.Node) {
  parts := decomposeUrl(url)
  return generateConditionalNodeChain(parts, func(node model.Node) bool {
    return IsNodeAcceptable(node)
  })
}

func GenerateNodeChain(url string) []model.Node {
  parts := decomposeUrl(url)
  _, nodes := generateConditionalNodeChain(parts, func(_ model.Node) bool {
    return true
  })
  return nodes
}

func FindOrCreateNode(name string, parent model.Node) model.Node {
  db := persistence.GetDatabase().DbMap()
  node := model.Node{}
  db.SelectOne(&node, "select * from nodes where Name=? and ParentId=?", name, parent.Id)
  if node.Id == 0 {
    node.Name = name
    insertAndAssociateNodeToParent(&node, parent)
  }
  return node
}

func UpdateFeedback(node model.Node, feedback float64) {
  db := persistence.GetDatabase().DbMap()
  avg := node.LocalFeedbackAverage
  fc := node.FeedbackCount
  newAvg := computeNewAverage(avg, float64(fc), feedback)
  node.LocalFeedbackAverage = newAvg
  node.FeedbackCount = fc + 1
  db.Update(&node)
}

func propagateFeedback(nodes []model.Node) {
  db := persistence.GetDatabase().DbMap()
  if len(nodes) <= 2 {
    return
  }
  for i := len(nodes) - 2; i >= 0; i-- {

  }
  for _, node := range nodes {
    node.RequestCount++
    db.Update(&node)
  }
}

func updateRequestForNodeChain(nodes []model.Node) {
  db := persistence.GetDatabase().DbMap()
  n := len(nodes)
  for _, node := range nodes {
    if n--; n > 0 {
      node.ThroughRequestCount++
    } else {
      node.RequestCount++
    }
    db.Update(&node)
  }
}

func insertAndAssociateNodeToParent(node *model.Node, parent model.Node) {
  db := persistence.GetDatabase().DbMap()
  node.ParentId = parent.Id
  parent.ChildrenCount++
  db.Update(&parent)
  db.Insert(&node)
}

func computeNewAverage(currentAverage, feedbackCount, newFeedback float64) float64 {
  return (currentAverage * feedbackCount + newFeedback) / feedbackCount + 1
}

func generateConditionalNodeChain(names []string, assert func(model.Node) bool) (bool, []model.Node) {
  nodes := make([]model.Node, 0, 0)
  parent := model.Node{}
  for _, name := range names {
    node := FindOrCreateNode(name, parent)
    if !assert(node) {
      return false, nodes
    }
    nodes = append(nodes, node)
    parent = node
  }
  return true, nodes
}

func decomposeUrl(url string) []string {
  return strings.Split(strings.Trim(url, "/"), "/")
}
