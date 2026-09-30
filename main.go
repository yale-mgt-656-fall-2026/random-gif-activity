package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		html := `<!DOCTYPE html>
<html>
<head>
  <title>My Go Server</title>
  <style>
    body { background-color: lightblue; }
  </style>
</head>
<body>
  <h1>Hello from Go!</h1>
  <img src="https://s3.amazonaws.com/files.656.mba/mgt656/fall-2026/random-gifs/adorbs/AttackingTheCatBuritto.gif" alt="A cute animal">
</body>
</html>`

		if _, err := 
		
		w.Write([]byte(html)); err != nil {
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
