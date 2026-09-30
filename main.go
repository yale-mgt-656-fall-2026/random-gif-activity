package main

import (
	"fmt"
	"html"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

const base = "https://s3.amazonaws.com/files.656.mba/mgt656/fall-2026/random-gifs/adorbs/"

var gifsByWord = map[string][]string{
	"cat": {
		base + "CatPat.gif",
		base + "CatPopsBaloon.gif",
		base + "DeterminedToiletCat.gif",
	},
	"puppy": {
		base + "CapedScooterDog.gif",
		base + "DogCarryon.gif",
		base + "Dog-Sofa.gif",
	},
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		// Read the search word from the URL, e.g. /?q=cat
		q := strings.ToLower(strings.TrimSpace(r.URL.Query().Get("q")))

		gifs, ok := gifsByWord[q]
		message := ""
		if !ok {
			// Empty or unknown word: choose from every GIF.
			gifs = append(gifsByWord["cat"], gifsByWord["puppy"]...)
			if q != "" {
				message = "No GIFs for that word. Here is a random one."
			}
		}
		gif := gifs[rand.Intn(len(gifs))]

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if _, err := fmt.Fprintf(w, `<!doctype html>
<html>
<body style="background: lightyellow; font-family: sans-serif; text-align: center">
  <h1>Random GIF</h1>
  <form method="get" action="/">
    <input name="q" placeholder="cat or puppy" value="%s">
    <button type="submit">Search</button>
  </form>
  <p>%s</p>
  <img src="%s" alt="A cute animal">
</body>
</html>`, html.EscapeString(q), message, gif); err != nil {
			log.Printf("write response: %v", err)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Open http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}