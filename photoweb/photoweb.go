package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
)

const UPLOAD_DIR = "./uploads"

var templates = make(map[string]*template.Template)

func init() {
	for _, tmpl := range []string{"upload", "list"} {
		t := template.Must(template.ParseFiles(tmpl + ".html"))
		templates[tmpl] = t
	}
}
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if err := renderHtml(w, "upload", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	if r.Method == "POST" {
		file, header, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		filename := header.Filename
		defer file.Close()
		targetFile, err := os.Create(UPLOAD_DIR + "/" + filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer targetFile.Close()
		if _, err := io.Copy(targetFile, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	if exists := isExist(imagePath); !exists {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	entries, err := os.ReadDir(UPLOAD_DIR)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	locals := make(map[string]interface{})
	images := []string{}
	for _, entry := range entries {
		fileInfo, err := entry.Info()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		images = append(images, fileInfo.Name())
	}
	locals["images"] = images
	if err = renderHtml(w, "list", locals); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func renderHtml(w http.ResponseWriter, templates string, locals map[string]interface{}) error {
	t, err := template.ParseFiles(templates + ".html")
	if err != nil {
		return err
	}
	return t.Execute(w, locals)
}

func main() {
	http.HandleFunc("/", listHandler)
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/view", viewHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAdnServe: " + err.Error())
	}
}
