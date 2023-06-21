package controllers

import (
	"bufio"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if CheckError404(w, r) {
		return
	}

	RenderTemplate(w, "home", nil)
}

// permet de revifier une erreur de type 404 ( page non trouvé )
func CheckError404(w http.ResponseWriter, r *http.Request) bool {
	req := r.URL.Path
	if req != "/" && req != "/artists" && req != "/artistDetails" && req != "/relation" && req != "/location" && req != "/date" {
		w.WriteHeader(http.StatusNotFound)
		tmpl := template.Must(template.ParseFiles("client/pages/404.html"))
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return true
		}
		tmpl.Execute(w, struct{ Success bool }{true})
		return true
	}
	return false
}

// declancher une erreur 404
func Status404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	tmpl := template.Must(template.ParseFiles("client/pages/404.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	tmpl.Execute(w, struct{ Success bool }{true})
}

func Status500(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	tmpl := template.Must(template.ParseFiles("client/pages/500.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	tmpl.Execute(w, struct{ Success bool }{true})
}

// declancher une erreur 404
func Status400(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	tmpl := template.Must(template.ParseFiles("client/pages/400.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	tmpl.Execute(w, struct{ Success bool }{true})
}

// allows to load the file .env
func LoadEnv(path string) error {
	file, err := os.Open(path)
	if err != nil {
		log.Println("🚨 " + err.Error())
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			log.Println("🚨 Your env file must be set")
		}
		key := parts[0]
		value := parts[1]
		os.Setenv(key, value)
	}
	return scanner.Err()
}
