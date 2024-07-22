package functions

import (
	"bytes"
	"fmt"
	"garden-builder/types"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"

	"go.abhg.dev/goldmark/wikilink"
)

var md = goldmark.New(
	goldmark.WithExtensions(extension.GFM, extension.Linkify,
		&wikilink.Extender{}),
	goldmark.WithParserOptions(
		parser.WithAutoHeadingID(),
	),
	goldmark.WithRendererOptions(
		html.WithHardWraps(),
		html.WithXHTML(),
	),
)

func GetArticleFiles(
	dirPath string,
) map[string]types.ArticleFile {
	fileNames, err := filepath.Glob(dirPath + "/*.md")

	if err != nil {
		panic(err)
	}

	files := make(map[string]types.ArticleFile)

	for _, fileName := range fileNames {
		file := createArticleNodeFromFileName(fileName)
		files[file.Id] = file
	}

	return files
}

func createArticleNodeFromFileName(fileName string) types.ArticleFile {
	content, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("⛔ Failed to read file")
		panic(err)
	}

	var buf bytes.Buffer
	if err := md.Convert(content, &buf); err != nil {
		panic(err)
	}

	htmlContent := buf.String()

	outlinks := getOutlinksFromHTML(htmlContent)

	file := types.ArticleFile{
		Id:      TextNormalizer(fileName[3 : len(fileName)-3]),
		Content: htmlContent,
		Outlink: outlinks,
	}

	return file
}

var internalLinkRegex = regexp.MustCompile(`<a href="([^":]+)"`)

func getOutlinksFromHTML(htmlContent string) []string {
	matches := internalLinkRegex.FindAllStringSubmatch(htmlContent, -1)
	outlinks := make([]string, len(matches))

	for i, match := range matches {
		link := match[1]

		linkExtension := filepath.Ext(link)

		if linkExtension == ".html" || linkExtension == ".md" {
			link = strings.Split(link, ".")[0]
		}

		decodedLink, err := url.QueryUnescape(link)

		if err != nil {
			panic(err)
		}

		normalizedLink := TextNormalizer(decodedLink)

		if slices.Contains(outlinks, normalizedLink) {
			continue
		}

		outlinks[i] = normalizedLink
	}

	return outlinks
}
