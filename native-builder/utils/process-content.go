package utils

import (
	"bytes"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/chai2010/webp"
	_ "github.com/strukturag/libheif/go/heif" // Import for side-effects (decoder registration)
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
	"go.abhg.dev/goldmark/frontmatter"
	"go.abhg.dev/goldmark/wikilink"
)

var markdownParser goldmark.Markdown

func init() {
	markdownParser = goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			extension.Linkify,
			&wikilink.Extender{
				Resolver: &combinedResolver{},
			},
			&frontmatter.Extender{},
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
			parser.WithASTTransformers(
				util.Prioritized(&imageConverterTransformer{}, 100),
			),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
}

// isImageEmbed checks if a wikilink node should be treated as an image embed.
// It checks the parent node kind and also the file extension of the target.
func isImageEmbed(n *wikilink.Node) bool {
	if n.Parent() != nil && n.Parent().Kind() == ast.KindImage {
		return true
	}
	// Fallback for custom image types like .heic that goldmark might not recognize as images.
	target := strings.ToLower(string(n.Target))
	imageExtensions := []string{".heic", ".jpg", ".jpeg", ".png", ".gif", ".webp"}
	for _, ext := range imageExtensions {
		if strings.HasSuffix(target, ext) {
			// This is an embed link `![[...]]` if the node has no children (no link text).
			return n.FirstChild() == nil
		}
	}
	return false
}

type combinedResolver struct{}

func (r *combinedResolver) ResolveWikilink(n *wikilink.Node) ([]byte, error) {
	if isImageEmbed(n) {
		webpPath, err := convertImageToWebP(string(n.Target))
		if err != nil {
			log.Printf("Error converting wikilink image %s: %v", string(n.Target), err)
			return n.Target, nil
		}
		return []byte(webpPath), nil
	}
	// It's a regular page link
	pagePath := "/" + TextNormalizer(string(n.Target)) + ".html"
	return []byte(pagePath), nil
}

type imageConverterTransformer struct{}

func (t *imageConverterTransformer) Priority() int {
	return 100
}

func (t *imageConverterTransformer) Transform(node *ast.Document, reader text.Reader, ctx parser.Context) {
	_ = ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}
		if img, ok := n.(*ast.Image); ok {
			src := string(img.Destination)
			if !strings.HasPrefix(src, "http://") && !strings.HasPrefix(src, "https://") {
				webpPath, err := convertImageToWebP(src)
				if err != nil {
					log.Printf("Error converting standard image %s: %v", src, err)
				} else {
					img.Destination = []byte(webpPath)
				}
			}
		}
		return ast.WalkContinue, nil
	})
}

func ProcessMarkdown(source []byte) (string, parser.Context, error) {
	var buf bytes.Buffer
	ctx := parser.NewContext()
	if err := markdownParser.Convert(source, &buf, parser.WithContext(ctx)); err != nil {
		return "", nil, err
	}
	return buf.String(), ctx, nil
}

func convertImageToWebP(srcPath string) (string, error) {
	decodedPath, err := url.PathUnescape(srcPath)
	if err != nil {
		decodedPath = srcPath
	}

	cleanSrcPath := strings.TrimPrefix(decodedPath, "./")
	var fullSrcPath string

	if strings.HasPrefix(cleanSrcPath, "../") {
		fullSrcPath = cleanSrcPath
	} else if !strings.Contains(cleanSrcPath, "/") {
		fullSrcPath = filepath.Join("..", "images", cleanSrcPath)
	} else {
		fullSrcPath = filepath.Join("..", cleanSrcPath)
	}

	fullSrcPath = filepath.Clean(fullSrcPath)

	ext := filepath.Ext(fullSrcPath)
	baseName := strings.TrimSuffix(filepath.Base(fullSrcPath), ext)
	webpBaseName := baseName + ".webp"

	distDirPath := "./dist/images"
	distPath := filepath.Join(distDirPath, webpBaseName)
	webPath := filepath.ToSlash(filepath.Join("/images", webpBaseName))

	if err := os.MkdirAll(distDirPath, 0755); err != nil {
		return "", fmt.Errorf("could not create dist/images directory: %w", err)
	}

	srcStat, err := os.Stat(fullSrcPath)
	if err != nil {
		return "", fmt.Errorf("could not stat source file %s: %w", fullSrcPath, err)
	}

	distStat, err := os.Stat(distPath)
	if err == nil && srcStat.ModTime().Before(distStat.ModTime()) {
		return webPath, nil
	}

	srcFile, err := os.Open(fullSrcPath)
	if err != nil {
		return "", fmt.Errorf("could not open source image %s: %w", fullSrcPath, err)
	}
	defer srcFile.Close()

	img, _, err := image.Decode(srcFile)
	if err != nil {
		return "", fmt.Errorf("could not decode image %s: %w", fullSrcPath, err)
	}

	destFile, err := os.Create(distPath)
	if err != nil {
		return "", fmt.Errorf("could not create destination file %s: %w", distPath, err)
	}
	defer destFile.Close()

	if err := webp.Encode(destFile, img, &webp.Options{Lossless: false, Quality: 80}); err != nil {
		return "", fmt.Errorf("could not encode to webp %s: %w", fullSrcPath, err)
	}

	log.Printf("Successfully converted %s to %s", fullSrcPath, distPath)
	return webPath, nil
}
