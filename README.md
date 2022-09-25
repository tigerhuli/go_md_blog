# readme

## 项目介绍

这是一个兴趣使然([for the fun of it](https://en.wikipedia.org/wiki/List_of_One-Punch_Man_characters#Saitama))的个人博客系统🤪, 使用[Gin](https://github.com/gin-gonic/gin)作为服务框架, 博客的主要内容以markdown组织, 示例网站: http://tigerhuli.com/.

虽然有参考[HUGO](https://gohugo.io/)的项目结构, 但这是一个仍是一个非常定制化的博客网站, 有很多硬编码逻辑, 不适合直接使用🙅🏻‍♂️.

该框架将尽量向非侵入式的目标靠拢. 非侵入的意思是, 框架和markdown文本是独立的, 不需要向框架进行"post"之类的操作. 使用时只需要指定一个目录, 由框架完成目录生成, 网页生成的工作. 这样的好处是, 可以比较自由地组织markdown文件.

项目结构说明:

```shell
.
├── cache # 缓存, 主要用于统计访问量
├── CHANGELOG.md # 变更日志, 基本没用
├── constant # 常量
├── content -> ../my_blog # 输入内容, 使用软连接连接到真实的内容中
├── content_html # markdown转html结果
│   ├── articles
│   └── impressions
├── cron # 定时任务, 包括markdown转html, 导航页生成等
│   ├── clear_expired_html.go
│   ├── common.go
│   ├── cron.go
│   ├── cron_test.go
│   ├── flush_page_view.go
│   ├── sync_article_mm.go
│   ├── sync_articles.go
│   ├── sync_article_tree.go
│   ├── sync_impressions.go
│   ├── to_html.go
│   └── to_html_test.go
├── DEVELOPLOG.md # 开发日志, 记录开发的功能和遇到的一些问题
├── go.mod
├── go.sum
├── layout # 前端展示的html页面
│   ├── about.html
│   ├── article.html
│   ├── article_navi.html
│   ├── footer.html
│   ├── head_extend.html
│   ├── impression.html
│   ├── impression_navi.html
│   ├── impression_navi_more.html
│   ├── index.html
│   ├── index_more.html
│   ├── navi.html
│   └── no_more.html
├── main.go # 主函数入口
├── model # 结构体定义
│   ├── article.go
│   └── impression.go
├── README.md
├── router # 请求处理逻辑
│   ├── about.go
│   ├── article.go
│   ├── article_navi.go
│   ├── image.go
│   ├── impression.go
│   ├── impression_navi.go
│   ├── impression_navi_more.go
│   ├── index.go
│   ├── index_more.go
│   └── root.go
└── static # 网站静态资源存储
```