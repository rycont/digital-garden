package main

import (
	"fmt"
	"garden-builder/functions"
	"slices"
)

func main() {
	files, err := functions.GetArticleFiles("..")

	if err != nil {
		fmt.Println("Failed to get list of files")
		panic(err)
	}

	graph := functions.CreateGraph(files)
	scoreById := functions.CalculateRank(graph)

	ids := make([]string, len(scoreById))

	i := 0
	for id := range scoreById {
		ids[i] = id
		i++
	}

	slices.Sort(ids)

	for _, id := range ids {
		fmt.Println(id)
	}
}
