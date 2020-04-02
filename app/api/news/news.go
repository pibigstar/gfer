package news

import "github.com/gogf/gf/net/ghttp"

// 控制器
type C struct{}

// 动态首页
func (c *C) Index(r *ghttp.Request) {
	r.Response.WriteTplDefault()
}
