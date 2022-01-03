package utils

type Vertex struct {
	Id    string
	Edges []string
}

type DirGraph struct {
	Vertices map[string]Vertex
}

func CreateDirGraph() DirGraph {
	return DirGraph{Vertices: make(map[string]Vertex)}
}

func AddEdgeToGraph(g DirGraph, nodeA string, nodeB string) {
	v, valid := g.Vertices[nodeA]
	if !valid {
		v = Vertex{Id: nodeA, Edges: make([]string, 0)}
	}
	v.Edges = append(v.Edges, nodeB)
	g.Vertices[nodeA] = v
}
