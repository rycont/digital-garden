package main

import (
	"cmp"
	"garden-builder/types"
	"garden-builder/utils"
	"math"
	"slices"
	"strings"
	"time"

	"github.com/microcosm-cc/bluemonday"
	cp "github.com/otiai10/copy"
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
	p := bluemonday.StripTagsPolicy().AddSpaceWhenStrippingTag(false)

	for i, id := range ids {
		file := files[id]

		title := file.Title

		if len(title) == 0 {
			title = id
		}

		if file.Lastmod.IsZero() {
			file.Lastmod = time.Now()
		}

		score := math.Floor(scoreById[id]*100) / 100

		description := p.Sanitize(file.Content)
		description = strings.ReplaceAll(description, "\n", " ")
		description = strings.ReplaceAll(description, "  ", " ")

		runeDescription := []rune(description)

		if len(runeDescription) > 150 {
			runeDescription = runeDescription[:150]
		}

		description = string(runeDescription)

		sortedArticles[i] = types.ArticlePage{
			Id:          id,
			Title:       title,
			Content:     file.Content,
			Description: description,
			Outlink:     file.Outlink,
			Inlink:      graph[id].Inlink,
			Score:       score,
			Lastmod:     file.Lastmod,
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

	utils.SavePages(pages, "./dist")
	utils.CreateSitemapFile(sortedArticles, "./dist")
	cp.Copy("./public", "./dist")
	cp.Copy("../images", "./dist/images")
}
