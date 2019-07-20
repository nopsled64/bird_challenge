package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web_res/welcome.html")
	})

	http.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web_res/start.html")
	})

	http.HandleFunc("/chaffinch", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web_res/chaffinch.html")
	})

	http.HandleFunc("/bird.mp3", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web_res/bird.mp3")
	})

	http.HandleFunc("/penguin", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web_res/penguin.html")
	})

	http.HandleFunc("/bird.jpg", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web_res/bird.jpg")
	})

	http.HandleFunc("/human.png", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web_res/human.png")
	})

	http.HandleFunc("/puffin", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web_res/puffin.html")
	})

	// http.HandleFunc("/magpie", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "web_res/puffin.html")
	// })

	// fs := http.FileServer(http.Dir("static/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)

	//appengine.Main()

}