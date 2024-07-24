package utils

import (
	"garden-builder/types"
	"net/url"
)

const email = "rycont@outlook.kr"
const linkPrefix = "https://garden.postica.app/"
const sourcePrefix = "https://github.com/rycont/digital-garden/blob/main/"

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
		GithubLink:   sourcePrefix + articlePage.Id + ".md",
		MailToString: createMailtoString(articlePage.Title, articlePage.Id),
	}
}

func createMailtoString(
	artileTitle string,
	articleId string,
) string {
	subject := "Reply: " + artileTitle
	body := "---\n원본 글: " + linkPrefix + articleId

	subject = url.QueryEscape(subject)
	body = url.QueryEscape(body)

	return "mailto:" + email + "?subject=" + subject + "&body=" + body
}
