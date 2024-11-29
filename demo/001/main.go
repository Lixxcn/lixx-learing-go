package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/yuin/goldmark"
)

// 递归查找 markdowns 目录下的所有 .md 文件
func findMarkdownFiles(rootDir string) ([]string, error) {
	var markdownFiles []string

	// 使用 Walk 递归遍历目录
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 如果是 Markdown 文件，添加到列表
		if !info.IsDir() && filepath.Ext(path) == ".md" {
			markdownFiles = append(markdownFiles, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return markdownFiles, nil
}

// 渲染 Markdown 文件为 HTML
func renderMarkdownToHTML(filePath string) (string, error) {
	// 读取 Markdown 文件内容
	mdContent, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// 创建一个 bytes.Buffer 用于存储渲染后的 HTML
	var buf bytes.Buffer

	// 使用 goldmark 渲染 Markdown 为 HTML
	md := goldmark.New()
	err = md.Convert(mdContent, &buf)
	if err != nil {
		return "", err
	}

	// 返回渲染后的 HTML
	return buf.String(), nil
}

// 处理显示目录和文件列表
func listMarkdownFiles(c *gin.Context, dir string) {
	files, err := findMarkdownFiles(dir)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error reading directory"})
		return
	}

	// 构建文件列表的 HTML
	var fileLinks string
	for _, file := range files {
		// 提取文件名并移除路径部分
		fileName := filepath.Base(file)
		fileLinks += fmt.Sprintf("<li><a href='/render/%s'>%s</a></li>", file, fileName)
	}

	// 返回目录页面并显示文件链接
	c.HTML(http.StatusOK, "index.html", gin.H{
		"files": fileLinks,
	})
}

func main() {
	// 初始化 Gin 引擎
	r := gin.Default()

	// 加载 HTML 模板
	r.LoadHTMLFiles("templates/index.html")

	// 处理目录页面显示
	r.GET("/", func(c *gin.Context) {
		listMarkdownFiles(c, "./markdowns")
	})

	// 处理 Markdown 渲染请求
	r.GET("/render/*file", func(c *gin.Context) {
		// 获取文件路径（移除 /render 前缀）
		filePath := c.Param("file")

		// 检查文件是否存在
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			c.JSON(http.StatusNotFound, gin.H{"message": "File not found"})
			return
		}

		// 渲染 Markdown 文件为 HTML
		html, err := renderMarkdownToHTML(filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error rendering markdown"})
			return
		}

		// 返回渲染后的 HTML
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
	})

	// 启动 Web 服务
	r.Run(":8080")
}
