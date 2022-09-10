package model

// Impression 感想集合
type Impression struct {
	Name        string `yaml:"name"`         // 标题
	ContentType string `yaml:"content_type"` // 内容类型
	ContentUrl  string `yaml:"content_url"`  // 内容指向URL
	Grade       string `yaml:"grade"`        // 打分
	Comment     string `yaml:"comment"`      // 评语
	CreateTime  string `yaml:"create_time"`  // 创建时间
	UpdateTime  string `yaml:"update_time"`  // 更新时间
	Cover       string `yaml:"cover"`        // impression封面
	Url         string `yaml:"url"`          // impression链接
}
