package utils

import (
	"fmt"
	"garden-builder/types"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"time"

	"go.abhg.dev/goldmark/frontmatter"
)

// GetArticleFiles finds all markdown files in a directory and processes them.
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

	// Use the new centralized markdown processing function
	htmlContent, ctx, err := ProcessMarkdown(content)
	if err != nil {
		panic(err)
	}

	// Extract frontmatter from the context returned by ProcessMarkdown
	frontmatterContent := frontmatter.Get(ctx)

	var fm types.ArticleFrontmatter
	if frontmatterContent != nil {
		if err := frontmatterContent.Decode(&fm); err != nil {
			fmt.Printf("⛔ Failed to decode frontmatter for %s: %v\n", fileName, err)
		}
	}

	outlinks := getOutlinksFromHTML(htmlContent)

	var parsedTime time.Time
	if fm.Date != "" {
		parsedTime, err = time.Parse(time.RFC3339Nano, fm.Date)
		if err != nil {
			fmt.Printf("⛔ Failed to parse date for %s: %v\n", fileName, err)
			// Fallback to current time if date parsing fails
			parsedTime = time.Now()
		}
	} else {
		// If no date in frontmatter, use file's mod time as a fallback
		fileInfo, statErr := os.Stat(fileName)
		if statErr == nil {
			parsedTime = fileInfo.ModTime()
		} else {
			parsedTime = time.Now()
		}
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
			// Strip extension for internal linking
			link = strings.TrimSuffix(link, linkExtension)
		}

		decodedLink, err := url.QueryUnescape(link)
		if err != nil {
			// Log error but continue
			fmt.Printf("⛔ Failed to unescape link '%s': %v\n", link, err)
			continue
		}

		normalizedLink := TextNormalizer(decodedLink)
		if len(normalizedLink) == 0 {
			continue
		}

		if !slices.Contains(outlinks, normalizedLink) {
			outlinks = append(outlinks, normalizedLink)
		}
	}

	return outlinks
}
