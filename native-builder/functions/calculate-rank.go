package functions

import (
	"garden-builder/types"
	"math"
	"slices"
)

func CalculateRank(graph map[string]types.GraphNode) map[string]float64 {
	scores := make(map[string]float64)

	for id, node := range graph {
		inlinkPower := len(node.Inlink) + 1
		outlinkPower := len(node.Outlink) + 10

		inlinkPower = inlinkPower * inlinkPower

		score := float64(inlinkPower) / math.Log(float64(outlinkPower))
		scores[id] = score
	}

	for i := 0; i < 5; i++ {
		nextScores := make(map[string]float64)

		for id, node := range graph {
			inlinkScore := 0.0
			outlinkScore := 0.0

			for _, inlink := range node.Inlink {
				if slices.Contains(node.Outlink, inlink) {
					inlinkScore += scores[inlink] / 3
				} else {
					inlinkScore += scores[inlink]
				}
			}

			for _, outlink := range node.Outlink {
				outlinkScore += scores[outlink]
			}

			nextScores[id] = scores[id] + math.Log(inlinkScore+outlinkScore)
		}

		scores = nextScores
	}

	return scores
}
