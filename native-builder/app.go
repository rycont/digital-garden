package main

import (
	"cmp"
	"garden-builder/functions"
	"garden-builder/types"
	"slices"
)

func main() {
	files := functions.GetArticleFiles("..")

	graph := functions.CreateGraph(files)
	scoreById := functions.CalculateScore(graph)

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
		sortedArticles[i] = types.ArticlePage{
			Id:      id,
			Content: file.Content,
			Outlink: file.Outlink,
			Inlink:  graph[id].Inlink,
			Score:   scoreById[id],
		}
	}

	functions.BuildArticleList(sortedArticles)
}
