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
  return generateConditionalNodeChain(url, func(node model.Node) bool {
    return IsNodeAcceptable(node)
  })
}

func GenerateNodeChain(url string) []model.Node {
  _, nodes := generateConditionalNodeChain(url, func(_ model.Node) bool {
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
    InsertAndAssociateNodeToParent(&node, parent)
  }
  return node
}

func InsertAndAssociateNodeToParent(node *model.Node, parent model.Node) {
  db := persistence.GetDatabase().DbMap()
  node.ParentId = parent.Id
  parent.ChildrenCount++
  db.Update(&parent)
  db.Insert(&node)
}

func UpdateFeedback(node model.Node, feedback float64) {
  db := persistence.GetDatabase().DbMap()
  avg := node.LocalAverage
  fc := node.FeedbackCount
  newAvg := computeNewAverage(avg, float64(fc), feedback)
  node.LocalAverage = newAvg
  node.FeedbackCount = fc + 1
  db.Update(&node)
}

func UpdateRequestForNodeChain(nodes []model.Node) {
  db := persistence.GetDatabase().DbMap()
  for _, node := range nodes {
    node.RequestCount++
    db.Update(&node)
  }
}

// Propagates the feedback up
func PropagateFeedback(nodes []model.Node) {
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

func computeNewAverage(avg, feedbackCount, newFeedback float64) float64 {
  return (avg * feedbackCount + newFeedback) / feedbackCount + 1
}

func generateConditionalNodeChain(url string, assert func(model.Node) bool) (bool, []model.Node) {
  urlParts := strings.Split(strings.Trim(url, "/"), "/")
  partsLen := len(urlParts)
  nodes := make([]model.Node, partsLen, partsLen)
  parent := model.Node{}
  for i, name := range urlParts {
    if i > 0 {
      parent = nodes[i - 1]
    }
    nodes[i] = FindOrCreateNode(name, parent)
    if !assert(nodes[i]) {
      return false, nodes[0:i]
    }
  }
  return true, nodes
}


