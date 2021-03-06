package goraph

import (
	"sort"
)

type Edge struct {
	toNodeID int
	weight   float64
}

func NewEdge(toNodeID int, weight float64) Edge {
	return Edge{
		toNodeID: toNodeID,
		weight:   weight,
	}
}

type Node struct {
	nodeID   int
	adjNodes map[int]Edge
}

func NewNode(nodeID int) *Node {
	g := new(Node)
	g.nodeID = nodeID
	g.adjNodes = make(map[int]Edge)
	return g
}

func (g *Node) AddEdge(edge Edge) {
	g.adjNodes[edge.toNodeID] = edge
}

type GraphType int

const (
	GraphTypeDirect GraphType = iota
	GraphTypeUndirect
)

type Graph struct {
	nodes     map[int]*Node
	graphType GraphType
}

func NewGraph(graphType GraphType) *Graph {
	g := new(Graph)
	g.nodes = make(map[int]*Node)
	g.graphType = graphType
	return g
}

func (g *Graph) addEdgeInternally(fromNodeID int, toNodeID int, weight float64) {
	newEdge := NewEdge(toNodeID, weight)

	fromNode, ok := g.nodes[fromNodeID]
	if !ok {
		newNode := NewNode(fromNodeID)
		newNode.AddEdge(newEdge)
		g.nodes[fromNodeID] = newNode
	} else {
		fromNode.AddEdge(newEdge)
	}
}

func (g *Graph) AddEdge(fromNodeID int, toNodeID int, weight float64) {
	g.addEdgeInternally(fromNodeID, toNodeID, weight)

	if g.graphType == GraphTypeUndirect {
		g.addEdgeInternally(toNodeID, fromNodeID, weight)
	}
}

type HeapEdge struct {
	fromNodeID int
	toNodeID   int
	weight     float64
}

func NewHeapEdge(fromNodeID, toNodeID int, weight float64) HeapEdge {
	e := HeapEdge{
		fromNodeID: fromNodeID,
		toNodeID:   toNodeID,
		weight:     weight,
	}
	return e
}

type HeapEdgeArr []HeapEdge

func NewHeapEdgeArr() HeapEdgeArr {
	h := make(HeapEdgeArr, 0)
	return h
}

func (h *HeapEdgeArr) Queue(node HeapEdge) {
	*h = append(*h, node)
	sort.Sort(*h)
}

func (h *HeapEdgeArr) DeQueue() HeapEdge {
	result := (*h)[0]
	*h = (*h)[1:h.Len()]
	return result
}

func (h *HeapEdgeArr) IsEmpty() bool {
	return h.Len() == 0
}

// For Sort interface
func (h HeapEdgeArr) Len() int {
	return len(h)
}

// For Sort interface
func (h HeapEdgeArr) Less(i, j int) bool {
	return h[i].weight < h[j].weight
}

// For Sort interface
func (h HeapEdgeArr) Swap(i, j int) {
	tmp := h[i]
	h[i] = h[j]
	h[j] = tmp
}
