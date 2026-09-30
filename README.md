## Changes made by Hrishi
Along with instructions for the 5 stories, I also added a custom image by adding it to a public repo. The h1, h2 and h3 paint a cohesive picture and lead to the image. The cat GIFs make it easy to relate to the logo. I will be setting the default option for GIFs to be cat only.

# Random GIF activity

Build a tiny Go web app that shows a GIF. You'll work in pairs and take turns
making changes. The point is to see how a **server** handles a browser's HTTP
request, reads information from that request, and sends back HTML.

## Get started

1. Find a partner! Meet somebody new. Make a [human] friend.
2. Fork this repo.
3. Open your fork in a GitHub Codespace (**Code → Codespaces → Create
   codespace on main**) or clone it to a computer with Go installed.
4. In the terminal, run `go run .`.
5. Open `http://localhost:8080`. In a Codespace, use the **Ports** tab to open
   port 8080 in your browser. You should see `Hello, world!`.
6. Open `main.go`. The function passed to `http.HandleFunc` handles requests
   for `/`. The `w` value is the response sent to the browser; `r` is the
   incoming request. Change the greeting, save the file, stop the server with
   Ctrl-C, and run `go run .` again to see your change.
7. File issues for the user stories below, then work through them together.
   You can use an AI coding assistant or write the code yourselves. Include
   a way to check each issue's result in its acceptance criteria.

The [GIF list](gifs.txt) links to 146 animal GIFs from the
[`adorbs` collection](https://github.com/snipe/animated-gifs/tree/master/adorbs).
The images are hosted in course storage, so you can use them without an API key.
Open a URL from the list in your browser to see what it shows.

## User stories

Work through these in order. We'll pause after each one to compare approaches.

1. As a visitor, I see **Hello, world!** at `/`. The starter already does this.
2. As a visitor, I see an HTML page with a heading and a colored background.
   Before writing the HTML response, set its content type with
   `w.Header().Set("Content-Type", "text/html; charset=utf-8")`.
3. As a visitor, I see a GIF on the page. Copy a URL from `gifs.txt` into an
   HTML image tag, such as `<img src="GIF_URL" alt="A cute animal">`.
4. As a visitor, I get a different GIF when I reload the page. Put at least
   two GIF URLs in a Go slice and choose one in the request handler.
   Browse `gifs.txt` and pick your favorites. Or choose them all!
5. As a visitor, I can choose to see one, two, or three GIFs at once. Add an
   HTML form with `method="get"` and `action="/"` and a field named `count`.
   The server reads `r.URL.Query().Get("count")` and sends back that many
   random GIFs. Use one GIF when the value is missing or invalid.

The server handles the form and chooses URLs; the browser then asks course
storage for those GIFs. No API key or image proxy is needed.

## Check your work

- Refreshing `/` changes the GIF at least some of the time.
- Choosing three produces a URL containing `?count=3` and shows three GIFs.
- An invalid count still shows one GIF and does not crash the server.
- The page still works after stopping and restarting `go run .`.

If you have time, add another way to choose GIFs or make the page look
nicer. This is an in-class exercise; you do not need to deploy or submit it.

## GIF credits

The GIFs are from [Snipe's `adorbs` collection](https://github.com/snipe/animated-gifs/tree/d5ff840d028c2438497e7a7709d6bb9d5f7c6d68/adorbs)
at commit `d5ff840d028c2438497e7a7709d6bb9d5f7c6d68`.
The [source README](https://github.com/snipe/animated-gifs/blob/master/README.md)
credits the respective image copyright holders; it does not provide a license.
Course copies expire after December 30, 2026.
