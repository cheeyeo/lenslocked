package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/cheeyeo/lenslocked/controllers"
	"github.com/cheeyeo/lenslocked/views"
)

func galleryHandler(w http.ResponseWriter, r *http.Request) {
	galleryID := chi.URLParam(r, "galleryID")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<h1>Gallery PAGE</h1>
<p>
GALLERY TO SHOW: %s
</p>
`, galleryID)
}

func main() {
	fmt.Println("Starting server...")

	r := chi.NewRouter()

	tpl := views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.Parse(filepath.Join("templates", "contact.gohtml")))
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.Parse(filepath.Join("templates", "faq.gohtml")))
	r.Get("/faq", controllers.StaticHandler(tpl))
	// Add middleware but only to /gallery route
	r.Group(func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Get("/gallery/{galleryID}", galleryHandler)
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("starting server on port 3000...")
	http.ListenAndServe(":3000", r)
}
