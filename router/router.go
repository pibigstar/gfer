package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gfer/app/api/blog"
	"github.com/gogf/gfer/app/api/index"
	"github.com/gogf/gfer/app/api/news"
	"github.com/gogf/gfer/app/api/question"
	"github.com/gogf/gfer/app/api/user"
	"github.com/gogf/gfer/app/service/middleware"
)

func init() {
	g.Server().Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.View)
		group.ALL("/index", new(index.C))
		group.ALL("/news", new(news.C))
		group.ALL("/blog", new(blog.C))
		group.ALL("/user", new(user.C))
		group.ALL("/question", new(question.C))
	})
}
