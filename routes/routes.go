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
		r.Route("/rms", func(rms chi.Router) {
			rms.Use(middleware.AuthMiddleware)
			rms.Route("/user", func(user chi.Router) {
				user.Use(middleware.UserMiddleware)
				user.Get("/", handler.AllRestaurant)
				user.Post("/newaddress", handler.AddAddress)
				user.Post("/", handler.AllDish)

				//rms.Route("/{id}", func(restaurant chi.Router) {
				//})

			})
			rms.Route("/admin", func(admin chi.Router) {
				admin.Use(middleware.AdminMiddleware)
				admin.Get("/subadmin", handler.Subadmin)
				admin.Get("/allusers", handler.UsersByAdmin)
				admin.Get("/restaurant", handler.AllRestaurant)

				admin.Route("/add", func(add chi.Router) {
					add.Post("/dish", handler.AddDish)
					add.Post("/subadmin", handler.AddSubAdmin)
					add.Post("/restaurant", handler.AddRestaurant)
					add.Post("/user", handler.AddUser)

				})

			})
		})
	})
	return &Server{router}
}

func (svc *Server) Run(port string) error {
	return http.ListenAndServe(port, svc)
}
