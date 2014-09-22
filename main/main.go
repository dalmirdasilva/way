package main

import (
  "github.com/dalmirdasilva/gorecord/persistence"
  "github.com/dalmirdasilva/way/model"
  "github.com/dalmirdasilva/way/web"
)

func main() {
  persistence.Initialize("mysql", map[string]string {
    "database": "way",
  })
  persistence.RegisterTables(map[string]interface{} {
    "nodes": model.Node{},
  })

  /*
  node1 := model.Node{Name: "Xota1"}
  log.Println(db.Insert(&node1))
  node2 := model.Node{Name: "Xota2", ParentId: node1.Id}
  log.Println(db.Insert(&node2))
*/
  //db.SelectOne(&node, "select * from nodes where Name=? and Parent=?", node.Id, parent)
  //nodes := lib.GetNodesChainFromUrl("www.google.com/a/b/c/d/")
  //fmt.Println(nodes[len(nodes) - 1].Name)
  //lib.UpdateFeedback(nodes[len(nodes) - 1], -9.0)

  web.RunServer()
}
