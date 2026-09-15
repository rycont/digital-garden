package utils

import (
	"garden-builder/types"
)

func CreateGraph(files map[string]types.ArticleFile) map[string]types.GraphNode {
	inlinkMap := make(map[string][]string)

	for id := range files {
		inlinkMap[id] = make([]string, 0)
	}

	for id, file := range files {
		for outLinkId := range file.Outlinks {
			inlinkMap[outLinkId] = append(inlinkMap[outLinkId], id)
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
			Id:       id,
			Power:    len(idArticleMap[id].Content) / 800,
			Outlinks: idArticleMap[id].Outlinks,
			Inlink:   inlinks,
		}
	}

	return graph
}
