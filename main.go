package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/csrf"

	"github.com/cheeyeo/lenslocked/controllers"
	"github.com/cheeyeo/lenslocked/models"
	"github.com/cheeyeo/lenslocked/templates"
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

	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))
	r.Get("/faq", controllers.FAQ(tpl))

	// Setup database connection
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Setup model services
	userService := models.UserService{
		DB: db,
	}
	usersC := controllers.Users{
		UserService: &userService,
	}
	usersC.Templates.New = views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	usersC.Templates.Signin = views.Must(views.ParseFS(templates.FS, "signin.gohtml", "tailwind.gohtml"))
	r.Get("/signup", usersC.New)
	r.Get("/signin", usersC.Signin)
	r.Post("/signin", usersC.ProcessSignIn)
	r.Post("/signup", usersC.Create)
	r.Get("/users/me", usersC.CurrentUser)

	// Add middleware but only to /gallery route
	tpl = views.Must(views.ParseFS(templates.FS, "newpage.gohtml", "tailwind.gohtml"))
	r.Get("/newpage", controllers.StaticHandler(tpl))

	r.Group(func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Get("/gallery/{galleryID}", galleryHandler)
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("starting server on port 3000...")

	// Add CSRF Middleware
	csrfKey := "gFvi45R4fy5xNBlnEeZtQbfAVCYEIAUX"
	csrfMW := csrf.Protect(
		[]byte(csrfKey),
		// TODO: Fix before deployment
		csrf.Secure(false),
	)
	http.ListenAndServe(":3000", csrfMW(r))
}
