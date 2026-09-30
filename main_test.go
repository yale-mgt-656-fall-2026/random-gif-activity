package main

import (
	"strings"
	"testing"
)

func TestPickGIFs(t *testing.T) {
	urls := []string{"a", "b", "c", "d"}
	picked := pickGIFs(urls, 3)
	if len(picked) != 3 {
		t.Fatalf("len(picked) = %d; want 3", len(picked))
	}
	for _, u := range picked {
		if u == "" {
			t.Fatal("picked GIF URL is empty")
		}
	}
}

func TestRenderPageAlwaysIncludesLogo(t *testing.T) {
	page := renderPage("cat")
	if !strings.Contains(page, "spin-the-blade/main/public/images/som_trivia_cat_logo.png") {
		t.Fatal("page does not include the SOM trivia logo")
	}
	if !strings.Contains(page, "SOM Trivia Cat logo") {
		t.Fatal("page does not include the logo alt text")
	}
}

func TestRenderPageUnknownAnimalFallsBack(t *testing.T) {
	page := renderPage("banana")
	if !strings.Contains(page, "We don't have that animal yet") {
		t.Fatal("page does not include the fallback message")
	}
	if !strings.Contains(page, `<img src="`) {
		t.Fatal("page does not include a fallback GIF")
	}
}
