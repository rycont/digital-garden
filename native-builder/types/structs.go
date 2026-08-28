package types

import "time"

type ArticleFile struct {
	Id      string
	Title   string
	Content string
	Outlink []string
	Lastmod time.Time
}

type GraphNode struct {
	Id      string
	Power   int
	Outlink []string
	Inlink  []string
}

type ArticlePage struct {
	Id          string
	Title       string
	Content     string
	Description string
	Outlink     []string
	Inlink      []string
	Score       float64
	Lastmod     time.Time
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
	Lastmod      time.Time
}

type SitemapItem struct {
	Id      string
	Lastmod string
}
