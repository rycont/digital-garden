package main

import (
	"cmp"
	"garden-builder/types"
	"garden-builder/utils"
	"math"
	"slices"
)

func main() {
	files := utils.GetArticleFiles("..")

	graph := utils.CreateGraph(files)
	scoreById := utils.CalculateScore(graph)

	ids := make([]string, len(scoreById))

	i := 0

	for id := range scoreById {
		ids[i] = id
		i++
	}

	slices.SortFunc(ids, func(i, j string) int {
		return cmp.Compare(scoreById[j], scoreById[i])
	})

	sortedArticles := make([]types.ArticlePage, len(graph))

	for i, id := range ids {
		file := files[id]

		title := file.Title

		if len(title) == 0 {
			title = id
		}

		score := math.Floor(scoreById[id]*100) / 100

		sortedArticles[i] = types.ArticlePage{
			Id:      id,
			Title:   title,
			Content: file.Content,
			Outlink: file.Outlink,
			Inlink:  graph[id].Inlink,
			Score:   score,
		}
	}

	articleListPageContent := utils.BuildArticleList(sortedArticles)

	idArticlePageMap := make(map[string]types.ArticlePage)

	for _, article := range sortedArticles {
		idArticlePageMap[article.Id] = article
	}

	articleIdHtmlContentMap := utils.BuildArticlePages(idArticlePageMap)

	pages := make(map[string]types.LayoutBuilderInput)

	listPage := utils.CreateArticleListPage(articleListPageContent)
	articlePages := utils.CreateArticlePages(articleIdHtmlContentMap, idArticlePageMap)

	pages["index"] = listPage

	for id, articlePage := range articlePages {
		pages[id] = articlePage
	}

	utils.SavePages(pages, "../dist")
}
