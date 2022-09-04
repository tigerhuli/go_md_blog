# develop log

## 将首页渲染为mind map

- status: doing
- update time: 2022-09-04 09:29:31

这一块最开始想简单了, 以为可以加一下css样式来实现这个功能. 后面搜索了一段时间后发现没有很好的解决方案.

在搜索的时候发看到一个问题[markmap: Insert the mindmap to the HTML](https://stackoverflow.com/questions/65517545/markmap-insert-the-mindmap-to-the-html), 这里面想要实现的功能和我的很像, 但是最终需要用到js, 看了一下需要渲染额内容主要是这样的:

```json
{
        "t": "root", "d": 0, "v": "", "c": [
            {
                "t": "heading", "d": 1, "v": "教學", "c": [
                    { "t": "heading", "d": 2, "v": "<a href=\"https://markmap.js.org/repl/\">Hugo</a>" },
                    {
                        "t": "heading", "d": 2, "v": "JS", "c": [
                            { "t": "heading", "d": 3, "v": "jquery" },
                            { "t": "heading", "d": 3, "v": "bootstrap" }
                        ]
                    }
                ]
            },
            { "t": "heading", "d": 1, "v": "科技", "c": [] }
        ]
}
```

就是一个json文本, 可以使用自定义结构体来构造.

## 补充首页生成

- status: done
- update time: 2022-09-04 00:46:13

任务详情:

- 根据生成的html路径生成首页导航.

## 补充favicon.ico

- status: done
- update time: 2022-09-04 00:00:12

在开发的过程中, 发现chrome会发送一个`GET /favicon.ico"`的请求, 通过搜索后才知道这个请求用来[设置地址栏的图标的](https://blog.csdn.net/allway2/article/details/109115253).

经过试验发现, favicon并非每次都会拉取, chrome会对其进行缓存. favicon的设置参考[How to Add a Favicon to your Site](https://www.w3.org/2005/10/howto-favicon)

## markdown转html

- status: done
- update time: 2022-09-04 00:00:05

需要处理的任务有:

1. 递归处理目标文件夹中的md文件, 并转换为html.
2. markdown中的本地图片地址替换为服务器图片地址.
3. 添加定时任务, 每天凌晨同步一次markdown到html.

有一个问题需要注意, 就是markdown中的图片需要处理相对路径的问题.

## 拉取文章

- status: done
- update time: 2022-09-03 16:59:17

拉取文章的接口, 返回html页面.

## 在页面中展示markdown

- status: done
- update time: 2022-09-03 11:37:46

两种转方式:

1. 在html中嵌入markdown tag. 一个比较好的参考是: [How to Render Markdown on a Web Page With md-block](https://www.makeuseof.com/md-block-render-markdown-web-page/).
2. 将markdown转换为html然后再嵌入到html中. Go中的装换工具有: blackfriday, goldmark等.

最后选择使用第二种方式, 因为第一种方式浏览器打开速度有点慢. 在选择第二种方式后, 最终使用goldmark作为转换工具, 主要是因为goldmark有插件能够支持解析mermaid.

使用gin的模板嵌入html时, 有一个非常棘手的问题, 就是嵌入的内容被加上了双引号. 后来搜索资料得知这是一种防注入的策略, 参考[Gin之渲染HTML模板](https://blog.csdn.net/weixin_52690231/article/details/125021658)中的内容解决了问题.

## 拉取图片

- status: done
- update time: 2022-09-03 11:35:33

使用Gin返回图片信息, 路径为: `/image/:name`.
