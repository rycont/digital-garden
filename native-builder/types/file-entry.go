package types

type ArticleFile struct {
	Id      string
	Content string
	Outlink []string
}

type GraphNode struct {
	Id      string
	Power   int
	Outlink []string
	Inlink  []string
}
