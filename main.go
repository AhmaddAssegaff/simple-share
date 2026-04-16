package main

import (
	"fmt"
	"html/template"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

type File struct {
	Name string
	Size string
}

func formatSize(size int64) string {
	const (
		KB = 1024
		MB = 1024 * KB
		GB = 1024 * MB
	)

	switch {
	case size >= GB:
		return fmt.Sprintf("%.2f GB", float64(size)/GB)
	case size >= MB:
		return fmt.Sprintf("%.2f MB", float64(size)/MB)
	case size >= KB:
		return fmt.Sprintf("%.2f KB", float64(size)/KB)
	default:
		return fmt.Sprintf("%d B", size)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir("./public")
	if err != nil {
		http.Error(w, "Gagal baca folder", 500)
		return
	}

	var fileList []File
	for _, f := range files {
		info, _ := f.Info()
		fileList = append(fileList, File{
			Name: f.Name(),
			Size: formatSize(info.Size()),
		})
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template error", 500)
		return
	}

	tmpl.Execute(w, fileList)
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "localhost"
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "localhost"
}

func main() {
	http.HandleFunc("/", handler)

	http.HandleFunc("/download/", func(w http.ResponseWriter, r *http.Request) {
		file := r.URL.Path[len("/download/"):]
		http.ServeFile(w, r, "./public/"+file)
	})

	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		err := r.ParseMultipartForm(100 << 20)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		defer file.Close()

		dst, err := os.Create("./public/" + handler.Filename)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer dst.Close()

		_, err = io.Copy(dst, file)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	http.HandleFunc("/delete/", func(w http.ResponseWriter, r *http.Request) {
		file := r.URL.Path[len("/delete/"):]

		decoded, err := url.PathUnescape(file)
		if err != nil {
			http.Error(w, "Invalid filename", 400)
			return
		}

		filename := filepath.Base(decoded)

		err = os.Remove("./public/" + filename)
		if err != nil {
			http.Error(w, "Gagal hapus file", 500)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	ip := getLocalIP()
	fmt.Println("Buka di device lain:", "http://"+ip+":8080")

	http.ListenAndServe("0.0.0.0:8080", nil)
}
