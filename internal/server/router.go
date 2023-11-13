package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/jeremycruzz/msds301-wk8/internal/controller"
)

func NewRouter(controller *controller.SummarizerController) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/summarize/{topic}", controller.GetSummary)
	r.Post("/summarize/{topic}/update", controller.UpdateSummary)

	return r
}
