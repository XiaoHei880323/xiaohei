package viewHandler

import (
	"net/http"
	"path/filepath"
	"strings"
)

// ViewFileHandler 提供 view 根目录下的静态 HTML 文件（兼容旧路径）
func ViewFileHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := filepath.Base(r.URL.Path)
		if strings.Contains(name, "..") || name == "." || name == "" {
			http.NotFound(w, r)
			return
		}
		filePath := filepath.Join("view", name+".html")
		http.ServeFile(w, r, filePath)
	}
}

// AdminViewFileHandler 提供 view/admin 目录下的顶层 HTML 页面
// 访问路径: GET /admin/page/:name  ->  view/admin/{name}.html
func AdminViewFileHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := filepath.Base(r.URL.Path)
		if strings.Contains(name, "..") || name == "." || name == "" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		filePath := filepath.Join("view", "admin", name+".html")
		http.ServeFile(w, r, filePath)
	}
}

// AdminModuleViewFileHandler 提供 view/admin/{module} 目录下的模块页面
// 访问路径: GET /admin/page/:module/:name  ->  view/admin/{module}/{name}.html
func AdminModuleViewFileHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// URL 形如 /admin/page/user/list，取后两段
		parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/admin/page/"), "/")
		if len(parts) != 2 {
			http.NotFound(w, r)
			return
		}
		module := filepath.Clean(parts[0])
		name := filepath.Clean(parts[1])
		if strings.Contains(module, "..") || strings.Contains(name, "..") ||
			module == "." || name == "." {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		filePath := filepath.Join("view", "admin", module, name+".html")
		http.ServeFile(w, r, filePath)
	}
}

// AdminStaticFileHandler 提供 view/admin 目录下的静态资源文件（JS/CSS）
// 访问路径: GET /admin/static/:name  ->  view/admin/{name}
func AdminStaticFileHandler() http.HandlerFunc {
	allowed := map[string]bool{
		".js":  true,
		".css": true,
		".png": true,
		".jpg": true,
		".ico": true,
		".svg": true,
	}
	return func(w http.ResponseWriter, r *http.Request) {
		name := filepath.Base(r.URL.Path)
		if strings.Contains(name, "..") || name == "." || name == "" {
			http.NotFound(w, r)
			return
		}
		ext := strings.ToLower(filepath.Ext(name))
		if !allowed[ext] {
			http.NotFound(w, r)
			return
		}
		filePath := filepath.Join("view", "admin", name)
		http.ServeFile(w, r, filePath)
	}
}
