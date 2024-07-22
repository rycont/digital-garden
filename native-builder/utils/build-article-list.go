package utils

import (
	"bytes"
	"garden-builder/types"
	"html/template"
)

var articleListTemplate, articleListErr = template.ParseFiles("views/article-list")

func BuildArticleList(articles []types.ArticlePage) string {
	if articleListErr != nil {
		panic(articleListErr)
	}

	renderedArticleList := new(bytes.Buffer)
	articleListErr = articleListTemplate.Execute(renderedArticleList, articles)

	if articleListErr != nil {
		panic(articleListErr)
	}

	return renderedArticleList.String()
}
