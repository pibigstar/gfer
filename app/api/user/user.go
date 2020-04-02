package user

import "github.com/gogf/gf/net/ghttp"

// 控制器
type C struct{}

// 用户列表
func (c *C) Index(r *ghttp.Request) {
	r.Response.WriteTplDefault()
}
