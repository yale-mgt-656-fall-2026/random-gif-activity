package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

var gifChoices = map[string][]string{
	"cat": {
		"https://s3.amazonaws.com/files.656.mba/mgt656/fall-2026/random-gifs/adorbs/AttackingTheCatBuritto.gif",
		"https://s3.amazonaws.com/files.656.mba/mgt656/fall-2026/random-gifs/adorbs/Cat%20Freakout.gif",
	},
	"puppy": {
		"https://s3.amazonaws.com/files.656.mba/mgt656/fall-2026/random-gifs/adorbs/CapedScooterDog.gif",
		"https://s3.amazonaws.com/files.656.mba/mgt656/fall-2026/random-gifs/adorbs/DerpyDog.gif",
	},
}

func pickGIFs(urls []string, count int) []string {
	if len(urls) == 0 {
		return nil
	}
	if count < 1 {
		count = 1
	}
	if count > len(urls) {
		count = len(urls)
	}

	shuffled := append([]string(nil), urls...)
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	return shuffled[:count]
}

func renderPage(query string) string {
	animal := strings.ToLower(strings.TrimSpace(query))
	choices, ok := gifChoices[animal]
	message := ""
	if !ok {
		message = `<p>We don't have that animal yet, so here's a random GIF.</p>`
		for _, animalGIFs := range gifChoices {
			choices = append(choices, animalGIFs...)
		}
	}
	selected := pickGIFs(choices, 1)
	var imgTags strings.Builder
	for _, gifURL := range selected {
		imgTags.WriteString(`<img src="`)
		imgTags.WriteString(gifURL)
		imgTags.WriteString(`" alt="Cute animal GIF" style="max-width: 280px; margin: 12px; border-radius: 12px;">`)
	}

	return `<!doctype html>
<html>
<head>
	<meta charset="utf-8">
	<title>SOM Trivia GIF Club</title>
	<style>
		body {
			margin: 0;
			background: linear-gradient(135deg, #f9f7d4, #f3d49d);
			font-family: Arial, sans-serif;
			color: #1b1b1b;
		}
		.page {
			max-width: 960px;
			margin: 0 auto;
			padding: 32px 20px 48px;
			text-align: center;
		}
		h1 {
			font-size: 2.5rem;
			margin-bottom: 0.25rem;
		}
		h2 {
			margin-top: 0.5rem;
		}
		form {
			margin: 20px 0 30px;
		}
		.gifs {
			display: flex;
			flex-wrap: wrap;
			justify-content: center;
			gap: 12px;
			margin: 24px 0;
		}
		.logo {
			display: block;
			width: min(100%, 780px);
			height: auto;
			margin: 28px auto 0;
		}
	</style>
</head>
<body>
	<div class="page">
		<h1>Hello, world!</h1>
		<h2>Excited to introduce the SOM Trivia Quizzing Club!</h2>
		<h3>Join the slack channel #club-som-trivia-quizzing!</h3>
		<form method="get" action="/">
  <label for="q">Animal:</label>
  <input id="q" name="q" type="text">
  <button type="submit">Search</button>
</form>
		` + message + `
		<div class="gifs">
			` + imgTags.String() + `
		</div>

		<img class="logo" src="https://raw.githubusercontent.com/Abercrombie35/spin-the-blade/main/public/images/som_trivia_cat_logo.png" alt="SOM Trivia Cat logo">
	</div>
</body>
</html>`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		query := r.URL.Query().Get("q")
		page := renderPage(query)
		if _, err := w.Write([]byte(page)); err != nil {
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
