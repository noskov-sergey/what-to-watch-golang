package http

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.ru/noskov-sergey/what_to_watch_golang/internal/model"
)

type Usecase interface {
	Get(ctx context.Context) (*model.Opinion, error)
	Create(ctx context.Context, opinion model.Opinion) (int, error)
	GetById(ctx context.Context, id int) (*model.Opinion, error)
}

type router struct {
	chi.Router
	usecase Usecase
}

func New(usecase Usecase) *router {
	r := &router{
		Router:  chi.NewRouter(),
		usecase: usecase,
	}

	r.Get("/", r.getRandomHandler)
	r.Get("/add", r.NewOpinionHandler)
	r.Post("/add", r.createOpinionHandler)
	r.Route("/opinions", func(oR chi.Router) {
		oR.Get("/{opinionID}", r.getOpinionHandler)
	})

	r.Get("/404", func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, "templates/errors/404.html")
	})
	r.Get("/500", func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, "templates/errors/500.html")
	})

	r.Handle("/static/*", http.FileServer(http.Dir("templates/")))

	return r
}
