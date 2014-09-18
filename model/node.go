package model

import "github.com/dalmirdasilva/gorecord/model"

type Node struct {
  model.Entry
  ParentId int64
  Name string
  Average float64
  FeedbackCount int64
  RequestCount int64
  RefusedRequestCount int64
}



