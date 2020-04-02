package middleware

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/text/gstr"
)

// 该中间件用于根据URL.Path自动设置mainTpl模板变量
func View(r *ghttp.Request) {
	array := gstr.SplitAndTrim(r.URL.Path, "/")
	if len(array) >= 3 {
		r.Assigns(g.Map{
			"mainTpl": gfile.Join(array[1], array[2]) + ".html",
		})
	} else if len(array) >= 2 {
		r.Assigns(g.Map{
			"mainTpl": gfile.Join("index", array[1]) + ".html",
		})
	}
	r.SetView(g.View("admin"))
	r.Middleware.Next()
}
