package cron

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestToHtml(t *testing.T) {
	md_path := `/data/repos/go_md_blog/input/hello world.md`
	html_path := `/data/repos/go_md_blog/output/hello world.html`
	md_content, err := os.ReadFile(md_path)
	if err != nil {
		log.Println(err.Error())
		return
	}

	toHTML(md_path, html_path, md_content)
}

func TestReplaceHtmlImage(t *testing.T) {
	content := `<h1>hello world</h1>
	<p>hello world</p>
	<p><img src="hello_world.png" alt=""></p>
	<div class="mermaid">flowchart

	a --&gt; b
	</div><script src="https://cdn.jsdelivr.net/npm/mermaid/dist/mermaid.min.js"></script><script>mermaid.initialize({startOnLoad: true});</script>`

	new_content := replaceHtmlImage("./input/test/hello world.md", []byte(content))
	fmt.Println(string(new_content))
}
