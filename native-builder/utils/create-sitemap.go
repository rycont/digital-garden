package utils

import (
	"garden-builder/types"
	"os"
	"text/template"
	"time"
)

var sitemapGroupTemplate, sitemapGroupTemplateError = template.ParseFiles("views/sitemap-group")
var sitemapTemplate, sitemapTemplateError = template.ParseFiles("views/sitemap")

func CreateSitemapFile(
	pages []types.ArticlePage,
	path string,
) {
	createSitemapGroupFile(path)
	CreateSitemapEntriesFile(pages, path)
}

func CreateSitemapEntriesFile(
	pages []types.ArticlePage,
	path string,
) {
	entries := createSitemapEntries(pages)

	if sitemapTemplateError != nil {
		panic(sitemapTemplateError)
	}

	file, err := os.Create(path + "/sitemap.xml")

	if err != nil {
		panic(err)
	}

	err = sitemapTemplate.Execute(file, entries)

	if err != nil {
		panic(err)
	}
}

func createSitemapEntries(pages []types.ArticlePage) []types.SitemapItem {
	entries := make([]types.SitemapItem, len(pages))

	for i, page := range pages {
		entries[i] = types.SitemapItem{
			Id:      page.Id,
			Lastmod: page.Lastmod.Format(time.RFC3339Nano),
		}
	}

	return entries
}

func createSitemapGroupFile(path string) {
	if sitemapGroupTemplateError != nil {
		panic(sitemapGroupTemplateError)
	}

	file, err := os.Create(path + "/sitemap-index.xml")

	if err != nil {
		panic(err)
	}

	currentTime := time.Now().Format(time.RFC3339Nano)
	err = sitemapGroupTemplate.Execute(file, currentTime)

	if err != nil {
		panic(err)
	}

	err = file.Close()

	if err != nil {
		panic(err)
	}
}
