package functions

import (
	"garden-builder/types"
)

func CreateGraph(files []types.ArticleFile) map[string]types.GraphNode {
	inlinkMap := make(map[string][]string)

	for id := range files {
		inlinkMap[files[id].Id] = make([]string, 0)
	}

	for _, file := range files {
		for _, outlink := range file.Outlink {
			inlinkMap[outlink] = append(inlinkMap[outlink], file.Id)
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

	// for _, value := range files {
	// 	key := value.Id

	// 	graph[key] = types.GraphNode{
	// 		Id:      value.Id,
	// 		Outlink: value.Outlink,
	// 		Inlink:  inlinkMap[value.Id],
	// 	}
	// }

	for id, inlinks := range inlinkMap {
		graph[id] = types.GraphNode{
			Id:      id,
			Outlink: idArticleMap[id].Outlink,
			Inlink:  inlinks,
		}
	}

	return graph
}
