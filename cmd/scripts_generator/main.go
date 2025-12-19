package main

import (
	"encoding/json"
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
	jsonpath := filepath.Join("lesson_scripts", path, "metadata.json")
	dstpath := filepath.Join("static/scripts", path, "script.html")

	var v struct {
		Title string `json:"title"`
	}

	jsondat, err := os.ReadFile(jsonpath)

	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	if err := json.Unmarshal(jsondat, &v); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	prefix := []byte("# " + v.Title + "\n")

	dat, err := os.ReadFile(srcpath)

	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	html := mdToHTML(append(prefix, dat...))
	f, err := create(dstpath)

	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}

	_, err = f.Write(html)
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
