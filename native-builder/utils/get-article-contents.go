package utils

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
	"time"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"

	"go.abhg.dev/goldmark/frontmatter"
	"go.abhg.dev/goldmark/wikilink"
)

type minifiedResolver struct {
}

func (minifiedResolver) ResolveWikilink(n *wikilink.Node) ([]byte, error) {
	modified := "/" + TextNormalizer(string(n.Target)) + ".html"

	return []byte(modified), nil
}

var md = goldmark.New(
	goldmark.WithExtensions(extension.GFM, extension.Linkify,
		&wikilink.Extender{
			Resolver: minifiedResolver{},
		},
		&frontmatter.Extender{}),
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

	var htmlBuffer bytes.Buffer
	ctx := parser.NewContext()

	if err := md.Convert(content, &htmlBuffer, parser.WithContext(ctx)); err != nil {
		panic(err)
	}

	htmlContent := htmlBuffer.String()
	frontmatterContent := frontmatter.Get(ctx)

	var fm types.ArticleFrontmatter

	if frontmatterContent != nil {
		frontmatterContent.Decode(&fm)
	}

	outlinks := getOutlinksFromHTML(htmlContent)

	var parsedTime time.Time

	if fm.Date != "" {
		parsedTime, err = time.Parse(time.RFC3339Nano, fm.Date)

		if err != nil {
			fmt.Println("⛔ Failed to parse date")
		}
	} else {
		fmt.Println("⛔ Date not found in frontmatter")
	}

	file := types.ArticleFile{
		Id:      TextNormalizer(fileName[3 : len(fileName)-3]),
		Title:   strings.Trim(fm.Title, " "),
		Content: htmlContent,
		Outlink: outlinks,
		Lastmod: parsedTime,
	}

	return file
}

var internalLinkRegex = regexp.MustCompile(`<a href="([^":]+)"`)

func getOutlinksFromHTML(htmlContent string) []string {
	matches := internalLinkRegex.FindAllStringSubmatch(htmlContent, -1)
	outlinks := make([]string, 0)

	for _, match := range matches {
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

		if len(normalizedLink) == 0 {
			continue
		}

		if slices.Contains(outlinks, normalizedLink) {
			continue
		}

		outlinks = append(outlinks, normalizedLink)
	}

	// Resize slice to remove empty elements

	return outlinks
}
