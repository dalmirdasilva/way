package lib

import (
  "github.com/dalmirdasilva/way/model"
  "github.com/dalmirdasilva/gorecord/persistence"
  "strings"
)

func GetNodeChainFromUrl(url string) []model.Node {
  urlParts := strings.Split(strings.Trim(url, "/"), "/")
  partsLen := len(urlParts)
  nodes := make([]model.Node, partsLen, partsLen)
  parent := model.Node{}
  for i, part := range urlParts {
    if i > 0 {
      parent = nodes[i - 1]
    }
    nodes[i] = FindOrCreateFromNameAndParent(part, parent)
  }
  return nodes
}

func FindOrCreateFromNameAndParent(name string, parent model.Node) model.Node {
  db := persistence.GetDatabase().DbMap()
  node := model.Node{}
  db.SelectOne(&node, "select * from nodes where Name=? and ParentId=?", name, parent.Id)
  if node.Id == 0 {
    node.ParentId = parent.Id
    node.Name = name
    db.Insert(&node)
  }
  return node
}

func UpdateFeedback(node model.Node, feedback float64) {
  db := persistence.GetDatabase().DbMap()
  avg := node.Average
  fc := node.FeedbackCount
  newAvg := (avg * float64(fc) + feedback) / float64(fc + 1)
  node.Average = newAvg
  node.FeedbackCount = fc + 1
  db.Update(&node)
}

func Update(node model.Node, feedback float64) {
  db := persistence.GetDatabase().DbMap()
  avg := node.Average
  fc := node.FeedbackCount
  newAvg := (avg * float64(fc) + feedback) / float64(fc + 1)
  node.Average = newAvg
  node.FeedbackCount = fc + 1
  db.Update(&node)
}

