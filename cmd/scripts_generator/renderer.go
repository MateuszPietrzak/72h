package main

import (
	"fmt"
	"io"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
)

func renderParagraph(w io.Writer, p *ast.Paragraph, entering bool) {
	if entering {
		io.WriteString(w, `<p class="text-lg ml-8 mb-4">`)
	} else {
		io.WriteString(w, `</p>`)
	}
}

func renderHeader(w io.Writer, h *ast.Heading, entering bool) {
	if entering {
		str := ""
		switch h.Level {
		case 1:
			str = `<h1 class="text-5xl font-bold mb-8">`
		case 2:
			str = `<h2 class="text-3xl font-bold mb-8">`
		case 3:
			str = `<h3 class="text-2xl font-bold mb-4">`
		default:
			str = fmt.Sprintf(`<h%d>`, h.Level)
		}
		io.WriteString(w, str)
	} else {
		str := fmt.Sprintf(`</h%d>`, h.Level)
		io.WriteString(w, str)
	}
}

func renderHook(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
	if para, ok := node.(*ast.Paragraph); ok {
		renderParagraph(w, para, entering)
		return ast.GoToNext, true
	}

	if header, ok := node.(*ast.Heading); ok {
		renderHeader(w, header, entering)
		return ast.GoToNext, true
	}
	return ast.GoToNext, false
}

func ProjectRenderer() *html.Renderer {
	opts := html.RendererOptions{
		Flags:          html.CommonFlags,
		RenderNodeHook: renderHook,
	}
	return html.NewRenderer(opts)
}
