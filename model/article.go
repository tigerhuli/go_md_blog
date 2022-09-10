package model

// Article 文章类型结构
type Article struct {
	Title      string `yaml:"title"`       // 标题
	Author     string `yaml:"author"`      // 作者
	CreateTime string `yaml:"create_time"` // 创建时间
	UpdateTime string `yaml:"update_time"` // 更新时间
	Cover      string `yaml:"cover"`       // 文章封面
	Url        string `yaml:"url"`         // 文章链接
}
