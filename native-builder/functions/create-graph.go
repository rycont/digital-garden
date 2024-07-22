package functions

import (
	"garden-builder/types"
)

func CreateGraph(files map[string]types.ArticleFile) map[string]types.GraphNode {
	inlinkMap := make(map[string][]string)

	for id := range files {
		inlinkMap[id] = make([]string, 0)
	}

	for id, file := range files {
		for _, outlink := range file.Outlink {
			inlinkMap[outlink] = append(inlinkMap[outlink], id)
		}
	}

	idArticleMap := make(map[string]types.ArticleFile)

	for _, file := range files {
		idArticleMap[file.Id] = file
	}

	graph := make(
		map[string]types.GraphNode,
		len(files),
	)

	for id, inlinks := range inlinkMap {
		graph[id] = types.GraphNode{
			Id:      id,
			Power:   len(idArticleMap[id].Content) / 800,
			Outlink: idArticleMap[id].Outlink,
			Inlink:  inlinks,
		}
	}

	return graph
}
