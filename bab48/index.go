package main

import "fmt"
import "net/http"
import "html/template"

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "apa kabar yanun?")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		var data = map[string]string {
			"Name" : "Yanun",
			"Message" : "Ganteng banget",
		}

		var t, err = template.ParseFiles("template.html")

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, data)
	})

	http.HandleFunc("/index", index)

	fmt.Println("starting web server at http://localhost:1234")
	http.ListenAndServe(":1234", nil)
}

