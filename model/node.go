package model

import "github.com/dalmirdasilva/gorecord/model"

type Node struct {
  model.Entry

  ParentId int64
  Name string

  RefusedRequestCount int64
  ThroughRefusedRequestCount int64

  RequestCount int64
  FeedbackCount int64
  LocalFeedbackAverage float64

  ChildrenCount int64
  ThroughRequestCount int64
  CumulativeFeedbackAverage float64
}



