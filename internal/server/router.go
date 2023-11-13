package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jeremycruzz/msds301-wk8/internal/controller"
)

func NewRouter(controller *controller.SummarizerController) *chi.Mux {
	r := chi.NewRouter()
	v1ApiRouter := V1ApiRouter(controller)

	r.Get("/summarize/{topic}", controller.GetSummary)
	r.Post("/summarize/{topic}/update", controller.UpdateSummary)

	r.Mount("/api/v1", v1ApiRouter)

	r.Handle("/assets/*", http.StripPrefix("/assets", http.FileServer(http.Dir("assets"))))

	r.Route("/summarize", func(r chi.Router) {
		r.Get("/", controller.ServeTemplate)
		r.Post("/", controller.ServeTemplate)
	})
	return r
}

func V1ApiRouter(controller *controller.SummarizerController) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/summarize/{topic}", controller.GetSummary)
	r.Post("/summarize/{topic}/update", controller.UpdateSummary)

	return r
}
