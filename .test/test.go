package main

import (
	"fmt"
	"github.com/gogf/gf-alpha/app/model/article"
)

func main() {
	r, err := article.Model.Where(1).All()
	fmt.Println(err)
	fmt.Println(r)
}
