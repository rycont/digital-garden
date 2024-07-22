package utils

import (
	"fmt"
	"garden-builder/types"
	"os"
	"path/filepath"
	"text/template"
)

var layoutTemplate, layoutTemplateError = template.ParseFiles("views/layout")

func SavePages(pages map[string]types.LayoutBuilderInput, directory string) {
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

	err = layoutTemplate.Execute(file, content)

	if err != nil {
		panic(err)
	}

	err = file.Close()

	if err != nil {
		panic(err)
	}

	fmt.Println("✅ Saved page to", filePath)
}
