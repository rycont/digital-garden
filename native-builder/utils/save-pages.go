package utils

import (
	"bytes"
	"fmt"
	"garden-builder/types"
	"os"
	"path/filepath"
	"text/template"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
)

var layoutTemplate, layoutTemplateError = template.ParseFiles("views/layout")
var minifier = loadHTMLMinifier()

func SavePages(pages map[string]types.LayoutBuilderInput, directory string) {
	if layoutTemplateError != nil {
		panic(layoutTemplateError)
	}

	err := os.RemoveAll(directory)

	if err != nil {
		panic(err)
	}

	for path, content := range pages {
		savePage(path, content, directory)
	}
}

func savePage(path string, content types.LayoutBuilderInput, directory string) {
	filePath := directory + "/" + path + ".html"

	err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)

	if err != nil {
		panic(err)
	}

	file, err := os.Create(filePath)

	if err != nil {
		panic(err)
	}

	useMinify := os.Getenv("GARDEN_NO_MINIFY") != "true"

	if useMinify {
		rendered := new(bytes.Buffer)
		err = layoutTemplate.Execute(rendered, content)
		minifier.Minify("text/html", file, rendered)
	} else {
		err = layoutTemplate.Execute(file, content)
	}

	if err != nil {
		panic(err)
	}
	err = file.Close()

	if err != nil {
		panic(err)
	}

	fmt.Println("✅ Saved page to", filePath)
}

func loadHTMLMinifier() *minify.M {
	minifier := minify.New()
	minifier.AddFunc("text/html", html.Minify)
	minifier.AddFunc("text/css", css.Minify)

	return minifier
}
