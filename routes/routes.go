package routes

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"rmsProject/handler"
	"rmsProject/middleware"
)

type Server struct {
	chi.Router
}

func Route() *Server {
	router := chi.NewRouter()
	router.Route("/", func(r chi.Router) {
		r.Route("/public", func(public chi.Router) {
			public.Post("/signup", handler.Signup)
			public.Post("/login", handler.Login)
		})
		r.Route("/user", func(user chi.Router) {
			user.Use(middleware.UserMiddleware)
			user.Get("/", handler.AllRestaurant)
			user.Post("/", handler.AllDish)

		})
		r.Route("/admin", func(admin chi.Router) {
			admin.Use(middleware.AdminMiddleware)
			admin.Post("/addrestaurant", handler.AddRestaurant)
			admin.Post("/addsubadmin", handler.AddSubAdmin)
			admin.Post("/adddish", handler.AddDish)

		})
	})
	return &Server{router}
}

func (svc *Server) Run(port string) error {
	return http.ListenAndServe(port, svc)
}
