package cron

import (
	"fmt"
	"go_md_blog/cache"
	"go_md_blog/constant"
	"io/ioutil"
	"path/filepath"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

// MmNode mind map node
type MmNode struct {
	T string    `json:"t"` // 节点类型
	D int32     `json:"d"` // 节点深度
	V string    `json:"v"` // 节点值
	C []*MmNode `json:"c"` // 子节点列表
}

//syncArticleMm ...
func syncArticleMm() {
	html_dir := constant.ArticlesHtmlPath
	content, err := genContentHtmlMd(html_dir)
	if err != nil {
		fmt.Println(err.Error())
	}

	cache.ArticleMm = content
}

// genContentHtmlMd ...
func genContentHtmlMd(dir string) (string, error) {
	root, err := genMmNodeDfs(dir, 0, true)
	if err != nil {
		return "", err
	}

	js_content, err := jsoniter.MarshalToString(root)
	if err != nil {
		return "", nil
	}

	content := `<script>
    ((e, t) => {
        const { Markmap: r } = e();
        window.mm = r.create("svg#mindmap", null, t)
    })(
        () => window.markmap, ${js_content}
    );
</script>`
	content = strings.ReplaceAll(content, "${js_content}", js_content)

	return content, nil
}

// genMmNodeDfs ...
func genMmNodeDfs(path string, d int32, is_dir bool) (*MmNode, error) {
	node := &MmNode{}
	node.D = d
	switch d {
	case 0:
		node.T = "root"
	default:
		node.T = "heading"
		node.V = genMmValue(path, is_dir)
	}

	if !is_dir {
		return node, nil
	}

	items, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		sub_node, err := genMmNodeDfs(fmt.Sprintf("%s/%s", path, item.Name()), d+1, item.IsDir())
		if err != nil {
			return nil, err
		}

		node.C = append(node.C, sub_node)
	}

	return node, nil
}

// genMmValue ...
func genMmValue(path string, is_dir bool) string {
	var content string
	if is_dir {
		content = `${name}`
	} else {
		content = `<a href="${url}">${name}</a>`
		url := strings.TrimPrefix(path, constant.ArticlesHtmlPath)
		url = fmt.Sprintf("http://tigerhuli.com/article/%s", url)
		content = strings.ReplaceAll(content, "${url}", url)
	}

	name := filepath.Base(path)
	name = strings.Split(name, ".")[0]
	content = strings.ReplaceAll(content, "${name}", name)

	return content
}
