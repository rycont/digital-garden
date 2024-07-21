package main

import (
	"cmp"
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

	slices.SortFunc(ids, func(i, j string) int {
		return cmp.Compare(scoreById[j], scoreById[i])
	})

	for _, id := range ids {
		fmt.Println(id, scoreById[id])
	}

	// [TODO]: 점수에 글 길이 반영하기
}
