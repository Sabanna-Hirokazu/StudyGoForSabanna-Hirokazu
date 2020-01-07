package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("html/*")
	// r.LoadHTMLFiles("html/top.html", "html/read.html") //特定のファイルを読み込みに使うことができる。，を使って複数同時に読み込みも可能
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "top.html", gin.H{})
	})
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "top.html", gin.H{})
	})
	r.GET("/tmpl", func(c *gin.Context) {
		c.HTML(http.StatusOK, "read.html", gin.H{
			"title": "Sample tmpl",
		})
	})
	r.GET("/tmpl/hoge", func(c *gin.Context) {
		c.HTML(http.StatusOK, "read.html", gin.H{
			"title":       "Sample tmpl",
			"description": "Hoge",
		})
	})
	r.GET("/tmpl/fuga", func(c *gin.Context) {
		c.HTML(http.StatusOK, "read.html", gin.H{
			"title":       "Sample tmpl",
			"description": "Fuga",
		})
	})
	r.Run()
}
