package utils

import (
	"bytes"
	"cmp"
	"garden-builder/types"
	"slices"
	"text/template"
)

var articlePageTemplate, articlePageError = template.ParseFiles("views/article-page")
var emptyPageTemplate, emptyPageError = template.ParseFiles("views/empty-page")

func BuildArticlePages(articles map[string]types.ArticlePage) map[string]string {
	if articlePageError != nil {
		panic(articlePageError)
	}

	articleIdHtmlContentMap := make(map[string]string)

	for id := range articles {
		articleIdHtmlContentMap[id] = buildArticlePage(id, articles)
	}

	return articleIdHtmlContentMap
}

func buildArticlePage(id string, articlesMap map[string]types.ArticlePage) string {
	article := articlesMap[id]

	if len(article.Inlink) > 1 {
		slices.SortFunc(article.Inlink, func(i, j string) int {
			return cmp.Compare(articlesMap[j].Score, articlesMap[i].Score)
		})
	}

	inlinks := make([]types.ArticlePageBuilderInputInlink, len(article.Inlink))

	for i, inlinkId := range article.Inlink {
		inlinks[i] = types.ArticlePageBuilderInputInlink{
			Id:    inlinkId,
			Title: inlinkId,
			Score: articlesMap[inlinkId].Score,
		}
	}

	content := article.Content

	if len(content) == 0 {
		content = buildEmptyContentPage(article.Title)
	}

	input := types.ArticlePageBuilderInput{
		Content:    content,
		HasInlinks: len(inlinks) > 0,
		Inlinks:    inlinks,
	}

	renderedArticlePage := new(bytes.Buffer)

	articlePageError := articlePageTemplate.Execute(renderedArticlePage, input)

	if articlePageError != nil {
		panic(articlePageError)
	}

	return renderedArticlePage.String()
}

func buildEmptyContentPage(title string) string {
	if emptyPageError != nil {
		panic(emptyPageError)
	}

	renderedArticlePage := new(bytes.Buffer)

	articlePageError := emptyPageTemplate.Execute(renderedArticlePage, title)

	if articlePageError != nil {
		panic(articlePageError)
	}

	return renderedArticlePage.String()
}
