package cron

import (
	"bytes"
	"fmt"
	"go_md_blog/cache"
	"go_md_blog/constant"
	"go_md_blog/model"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"gopkg.in/yaml.v2"
)

// syncImpressions 同步Impressions到html
func syncImpressions() {
	input_dir := constant.ImpressionsPath
	md_paths, err := getMdPaths(input_dir)
	if err != nil {
		fmt.Println(err.Error())
	}

	var impressions []model.Impression
	for _, md_path := range md_paths {
		html_path := strings.TrimPrefix(md_path, input_dir)
		html_path = strings.Split(html_path, ".")[0]
		html_path = fmt.Sprintf("%s%s.html", constant.ImpressionsHtmlPath, html_path)

		md_content, err := os.ReadFile(md_path)
		if err != nil {
			log.Println(err.Error())
			return
		}

		impression, md_content, err := getImpressionFromMd(md_path, html_path, md_content)
		if err != nil {
			log.Printf("getArticleFromMdPath failed: %s", err.Error())
			panic(err)
		}

		impressions = append(impressions, impression)

		toHTML(md_path, html_path, md_content)
	}

	sort.Slice(impressions, func(i, j int) bool { return impressions[i].UpdateTime > impressions[j].UpdateTime })
	cache.Impressions = impressions
}

// getArticleFromMd 从md文件中提起Impression信息
func getImpressionFromMd(md_path, html_path string, md_content []byte) (model.Impression, []byte, error) {
	r := regexp.MustCompile(`(^---)([\s\S]*?)(---)`)
	submatches := r.FindSubmatch([]byte(md_content))
	if len(submatches) == 0 {
		return model.Impression{}, nil, fmt.Errorf("no article info find")
	}

	md_content = bytes.Replace(md_content, submatches[0], nil, 1)

	yaml_content := submatches[2]
	yaml_content = bytes.TrimSpace(yaml_content)
	var impression model.Impression
	err := yaml.Unmarshal(yaml_content, &impression)
	if err != nil {
		return model.Impression{}, nil, fmt.Errorf("unmarshal article %s info failed: %s", md_path, err)
	}

	if len(impression.Name) == 0 {
		name := filepath.Base(md_path)
		name = strings.Split(name, ".")[0]
		impression.Name = name
	}

	if len(impression.Url) == 0 {
		url := strings.TrimPrefix(html_path, constant.ImpressionsHtmlPath)
		url = fmt.Sprintf("http://tigerhuli.com/impression/%s", url)
		impression.Url = url
	}

	var emoji string
	switch impression.ContentType {
	case "movie":
		emoji = "&#127909;"
	case "video":
		emoji = "&#128250;"
	case "book":
		emoji = "&#128214;"
	default:
		emoji = "&#127760;"
	}

	impression.Stars = ""
	for i := 1; i <= int(impression.Grade) && i <= 5; i++ {
		impression.Stars += "&#11088;"
	}

	impression.Name = fmt.Sprintf("%s %s", emoji, impression.Name)
	log.Println(impression.Tags)
	impression.ColorTags = getColorTags(impression.Tags)
	log.Println(impression.ColorTags)

	return impression, md_content, nil
}

// getColorTags 生成有色彩的tag列表
func getColorTags(tags []string) string {
	var color_tags_builder strings.Builder
	for _, tag := range tags {
		color_tag := getColorTag(tag)
		color_tags_builder.WriteString(color_tag)
	}

	return color_tags_builder.String()
}

// getColorTag 获取有色彩的tag
func getColorTag(tag string) string {
	bg_color := getTagBGColor(tag)
	color_tag := `<p class="gallery_comment" id="tag" style="background-color:${color};">${tag}</p>`
	color_tag = strings.ReplaceAll(color_tag, "${color}", bg_color)
	color_tag = strings.ReplaceAll(color_tag, "${tag}", tag)
	return color_tag
}

var colorList []string
var colorMap map[string]int

func init() {
	colorList = []string{"#f2cbcb", "#f2cbcb", "#f2e2cb", "#f2e2cb", "#eff2cb", "#d6f2cb", "#d6f2cb",
		"#cbf2d5", "#cbf2d5", "#cbdbf2", "#cccbf2", "#decbf2", "#decbf2", "#f2cbea", "#f2cbea"}
	colorMap = make(map[string]int)
}

// getTagBGColor 获取随机背景延时
func getTagBGColor(tag string) string {
	index, ok := colorMap[tag]
	if ok {
		return colorList[index]
	}

	index = rand.Int() % len(colorList)
	colorMap[tag] = index

	return colorList[index]
}
