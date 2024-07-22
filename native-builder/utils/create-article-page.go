package utils

import "garden-builder/types"

func CreateArticlePages(
	articleIdHtmlContentMap map[string]string,
	idArticlePageMap map[string]types.ArticlePage,
) map[string]types.LayoutBuilderInput {
	pages := make(map[string]types.LayoutBuilderInput)

	for id, articlePage := range idArticlePageMap {
		htmlContent := articleIdHtmlContentMap[id]
		layoutBuilderInput := articlePagesToLayoutBuilderInput(htmlContent, articlePage)

		pages[id] = layoutBuilderInput
	}

	return pages
}

func articlePagesToLayoutBuilderInput(
	htmlContent string,
	articlePage types.ArticlePage,
) types.LayoutBuilderInput {
	return types.LayoutBuilderInput{
		Content:      htmlContent,
		Title:        articlePage.Title,
		Description:  articlePage.Title + "에 관련한 글입니다.",
		GithubLink:   "https://github.com/rycont/digital-garden",
		MailToString: "mailto:rycont@outlook.kr",
	}
}
