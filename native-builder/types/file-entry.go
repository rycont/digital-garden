package types

type ArticleFile struct {
	Id      string
	Content string
	Outlink []string
}

type GraphNode struct {
	Id      string
	Outlink []string
	Inlink  []string
}
