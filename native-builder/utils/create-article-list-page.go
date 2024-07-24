package utils

import "garden-builder/types"

func CreateArticleListPage(content string) types.LayoutBuilderInput {
	return types.LayoutBuilderInput{
		Title:        "Garden",
		Description:  "Digital Garden",
		GithubLink:   "https://github.com/rycont/digital-garden",
		MailToString: "mailto:rycont@outlook.kr",
		Content:      content,
	}
}
