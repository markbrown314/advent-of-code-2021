package graph

type Vertex struct {
	Id    string
	Edges map[string]bool
}

type Graph struct {
	Vertices map[string]Vertex
}

func New() Graph {
	return Graph{Vertices: make(map[string]Vertex)}
}

func (g *Graph) insertVertex(id string) Vertex {
	v, valid := g.Vertices[id]
	if !valid {
		v = Vertex{Id: id, Edges: make(map[string]bool)}
	}
	g.Vertices[id] = v
	return v
}

func linkVertices(vA Vertex, vB Vertex, directed bool) {
	vA.Edges[vB.Id] = true
	if !directed {
		vB.Edges[vA.Id] = true
	}
}

func (g *Graph) AddEdge(idA string, idB string, directed bool) {
	vA := g.insertVertex(idA)
	vB := g.insertVertex(idB)
	linkVertices(vA, vB, directed)
}
