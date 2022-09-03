# develop log

## 补充favicon.ico

- status: doing
- update time: 2022-09-03 17:00:04

在开发的过程中, 发现chrome会发送一个`GET /favicon.ico"`的请求, 通过搜索后才知道这个请求用来[设置地址栏的图标的](https://blog.csdn.net/allway2/article/details/109115253).

## markdown转html

- status: todo
- update time: 2022-09-03 15:19:11

主要有两个任务需要处理:

1. 递归处理目标文件夹中的md文件, 并转换为html
2. markdown中的本地图片地址替换为服务器图片地址.

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
