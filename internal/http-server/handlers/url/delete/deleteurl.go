package deleteurl

import (
	"log/slog"
	"net/http"
	"shorturl/internal/lib/api"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type URLDeleter interface {
	DeleteURL(alias string) error
}

type Response struct {
	api.Response
}

func New(log *slog.Logger, urlDeleter URLDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.delete.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		alias := chi.URLParam(r, "alias")
		if alias == "" {
			log.Info("alias is empty")
			return
		}

		err := urlDeleter.DeleteURL(alias)
		if err != nil {
			log.Error(err.Error())
			render.JSON(w, r, api.Error("internal error"))
			return
		}

		log.Info("url deleted", slog.String("alias", alias))

		responseOK(w, r)
	}
}

func responseOK(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, Response{
		Response: api.OK(),
	})
}
