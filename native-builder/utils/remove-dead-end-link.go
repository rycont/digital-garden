package utils

import (
	"garden-builder/types"
)

func RemoveDeadEndLink(idArticlePageMap map[string]types.ArticlePage) map[string]types.ArticlePage {
	newMap := make(map[string]types.ArticlePage)

	for id := range idArticlePageMap {
		newArticle := idArticlePageMap[id]
		newArticle.Content = removeDeadEndLinkPage(id, idArticlePageMap)

		newMap[id] = newArticle
	}

	return newMap
}

func removeDeadEndLinkPage(id string, idArticlePageMap map[string]types.ArticlePage) string {
	if len(idArticlePageMap[id].Outlinks) == 0 {
		return idArticlePageMap[id].Content
	}

	article := idArticlePageMap[id]

	for _, outlinks := range article.Outlinks {
		for _, outlink := range outlinks {
			targetArticle, exists := idArticlePageMap[outlink.Link]

			if !exists {
				continue
			}

			isDeadEnd := len(targetArticle.Inlink) == 1 && targetArticle.Content == ""
			if !isDeadEnd {
				continue
			}

			article.Content = article.Content[:outlink.Range[0]] + "<z" + article.Content[(outlink.Range[0]+2):]
			article.Content = article.Content[:outlink.Range[2]] + "</z" + article.Content[(outlink.Range[2]+3):]
		}
	}

	return article.Content
}
