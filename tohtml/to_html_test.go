package tohtml

import (
	"fmt"
	"testing"
)

func TestToHtml(t *testing.T) {
	md_path := `/data/repos/go_md_blog/input/hello world.md`
	html_path := `/data/repos/go_md_blog/output/hello world.html`
	ToHTML(md_path, html_path)
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
