package main

import (
	"fmt"
	"math/rand"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

// Так как ID для категорий и локаций назначаются не плотно,
// то нужно в точности проэмулировать процесс их создания
type Node struct {
  id int64
  parent *Node
  children []*Node
}

func (n *Node) addChild(id int64) {
  newNode := Node {
    id: id,
    parent: n,
    children: nil,
  }
  n.children = append(n.children, &newNode)
}

func (n *Node) traverse() []int64 {
  var res []int64
  for _, child := range n.children {
    res = append(res, child.traverse()...)
  }
  res = append(res, n.id)
  return res
}

func generateTreeIds(h int, nodesOnLevel int) []int64 {
  root := Node {
    id: 1,
    parent: nil,
    children: nil,
  }

  var id int64 = 2
  var generateTree func(node *Node, hCur int);
  generateTree = func(node *Node, hCur int) {
    if(hCur == 0) {
      return
    }
    for i := 0; i < nodesOnLevel; i++ {
      child := Node {
        id: id,
        parent: node,
        children: nil,
      }
      id++
      node.children = append(node.children, &child)
      generateTree(&child, hCur - 1)
    }
  }

  generateTree(&root, h)
  return root.traverse()
}

var categories []int64 = generateTreeIds(4, 10)
var locations []int64 = generateTreeIds(5, 7)
var userIds = []int64{ 2100,
  2200,
  2300,
  2400,
  2500,
  2600,
  2700,
  2800,
  2900,
  3000,
  3100,
  3200,
  3300,
  3400,
  3500,
  3600,
  3700,
  3800,
  3900,
  4000,
  4100,
  4200,
}

func NewCustomTargeter() vegeta.Targeter {
    return func(tgt *vegeta.Target) error {
        if tgt == nil {
            return vegeta.ErrNilTarget
        }
        categoryIndex := rand.Intn(len(categories))
        locationIndex := rand.Intn(len(locations))
        userIndex := rand.Intn(len(userIds))

        category := categories[categoryIndex]
        location := locations[locationIndex]
        userId := userIds[userIndex]

        tgt.Method = "GET"
        tgt.URL = fmt.Sprintf("http://localhost:3000/api/v1/price?location_id=%d&category_id=%d&user_id=%d", location, category, userId)
        return nil
    }
}

func main() {
  rps := 2000
  rate := vegeta.Rate { Freq: rps, Per: time.Second }
  duration := time.Second * 60
  targeter := NewCustomTargeter()
  attacker := vegeta.NewAttacker()

  fmt.Printf("starting load test with RPS = %d\n", rps)
  var metrics vegeta.Metrics
  for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
    metrics.Add(res)
  }
  metrics.Close()

  fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
  fmt.Printf("Mean latency: %s\n", metrics.Latencies.Mean)
  fmt.Printf("50th percentile: %s\n", metrics.Latencies.P50)
  fmt.Printf("Total requests: %d\n", metrics.Requests)
}
