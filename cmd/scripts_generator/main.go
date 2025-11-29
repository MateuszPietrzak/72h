package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/MateuszPietrzak/72h/templates/pages"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

var mds = `# header

Sample text.

[link](http://duckduckgo.com)
`

func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	// htmlFlags := html.CommonFlags | html.HrefTargetBlank
	// opts := html.RendererOptions{Flags: htmlFlags}
	// renderer := html.NewRenderer(opts)

	renderer2 := ProjectRenderer()

	return markdown.Render(doc, renderer2)
}

func main() {
	md := []byte(mds)
	html := mdToHTML(md)

	fmt.Printf("%s\n", html)

	f, err := os.Create("static/scripts/hello.html")
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}

	err = pages.Script(string(html)).Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}

}
