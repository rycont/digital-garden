package utils

import (
	"garden-builder/types"
	"net/url"
	"strings"
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
	content := htmlContent
	if !articlePage.Lastmod.IsZero() {
		dateString := "<p>최종 수정: " + articlePage.Lastmod.Format("2006-01-02") + "</p>"
		content = strings.Replace(content, "</h1>", "</h1>"+dateString, 1)
	}

	return types.LayoutBuilderInput{
		Content:      content,
		Title:        articlePage.Title,
		Description:  articlePage.Description,
		GithubLink:   sourcePrefix + articlePage.Id + ".md",
		MailToString: createMailtoString(articlePage.Title, articlePage.Id),
		Lastmod:      articlePage.Lastmod,
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
