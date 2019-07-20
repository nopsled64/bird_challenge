package main

import (
	"net/http"
	"time"

	"google.golang.org/appengine"
)

func addCookie(w http.ResponseWriter, name string, value string) {
	expire := time.Now().AddDate(0, 0, 5)
	cookie := http.Cookie{
		Name:    name,
		Value:   value,
		Expires: expire,
	}
	http.SetCookie(w, &cookie)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		addCookie(w, "Bird", "GoldFinch")
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

	http.HandleFunc("/magpie", func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("Bird")
		if err != nil {
			w.Write([]byte("You're no bird.... Go back to the beginning.\n"))
		} else if c.Value == "Greenfinch" {
			w.Write([]byte("Hey, Greenfinch! Great to see you!\n\nIt's always great to see another Greenfinch around here.\n\nI'll let you into a little secret, there's more to this bird shirt challenge than meets the eye...\n\nThere's been somewhat of a cagey character behind the scenes, capitalising some very interesting alphabetical characters...\n\nAll I know is that there are three four letter words hidden in there, and something about a 'text record'"))
		} else if c.Value == "GoldFinch" {
			w.Write([]byte("Sorry, you're a Goldfinch. I don't let in Goldfinches. Greenfinches only I'm afraid"))
		}
	})

	http.HandleFunc("/bigbird", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web_res/bigbird.mp3")
	})

	http.HandleFunc("/deadlydodo", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web_res/deadlydodo.html")
	})

	// fs := http.FileServer(http.Dir("static/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	//http.ListenAndServe(":8080", nil)

	appengine.Main()

}
