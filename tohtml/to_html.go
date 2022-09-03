package tohtml

import (
	"fmt"
	"io/ioutil"

	"github.com/russross/blackfriday"
)

func ToHTML() {
	name := "/data/repos/go_md_blog/input/hello world.md"
	mk_content, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err.Error())
	}
	html_content := blackfriday.MarkdownCommon(mk_content)
	ioutil.WriteFile("/data/repos/go_md_blog/output/hello world.html", html_content, 666)
}
