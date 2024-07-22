package types

type ArticleFile struct {
	Id      string
	Title   string
	Content string
	Outlink []string
}

type GraphNode struct {
	Id      string
	Power   int
	Outlink []string
	Inlink  []string
}

type ArticlePage struct {
	Id      string
	Title   string
	Content string
	Outlink []string
	Inlink  []string
	Score   float64
}

type ArticleFrontmatter struct {
	Title string `yaml:"title"`
	Date  string `yaml:"date"`
}

type ArticlePageBuilderInput struct {
	Content    string
	HasInlinks bool
	Inlinks    []ArticlePageBuilderInputInlink
}

type ArticlePageBuilderInputInlink struct {
	Id    string
	Title string
	Score float64
}

type LayoutBuilderInput struct {
	Content      string
	Description  string
	Title        string
	MailToString string
	GithubLink   string
}
