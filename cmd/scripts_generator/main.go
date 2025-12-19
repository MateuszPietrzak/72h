package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

func mdToHTML(md []byte) []byte {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	renderer := ProjectRenderer()

	return markdown.Render(doc, renderer)
}

// Source - https://stackoverflow.com/a
// Posted by stantonJones
// Retrieved 2025-12-19, License - CC BY-SA 4.0
func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

func renderScript(path string) {
	srcpath := filepath.Join("lesson_scripts", path, "scenariusz.md")
	dstpath := filepath.Join("static/scripts", path, "script.html")

	dat, err := os.ReadFile(srcpath)

	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	html := mdToHTML(dat)
	f, err := create(dstpath)

	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}

	_, err = f.Write(html) //pages.Script(string(html)).Render(context.Background(), f)
	if err != nil {
		log.Fatalf("Failed to write to the output file: %v", err)
	}
}

func main() {
	renderScript("Moduł 1 - Zespół w akcji/1. Budowanie zespołu")
	renderScript("Moduł 1 - Zespół w akcji/2. Komunikacja w zespole")
	renderScript("Moduł 1 - Zespół w akcji/3. Motywacja zespołu")
	renderScript("Moduł 1 - Zespół w akcji/4. Planowanie działania w zespole")
	renderScript("Moduł 1 - Zespół w akcji/5. Wybór lidera")
}
