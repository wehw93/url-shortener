package deleteURL

import (
	"log/slog"
	"net/http"
	resp "url-shortener/internal/lib/api/response"
	"url-shortener/internal/lib/logger/sl"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type URLDeleter interface {
	DeleteUrl(alias string) error
}

func New(log *slog.Logger, urlDeleter URLDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.delet.New"
		log = log.With(
			slog.String("op", op),
			slog.String("requset_id", middleware.GetReqID(r.Context())),
		)
		alias := chi.URLParam(r, "alias")
		if alias == "" {
			log.Info("alias is empty")
			render.JSON(w, r, resp.Error("not found"))
			return
		}
		err := urlDeleter.DeleteUrl(alias)
		if err != nil {
			log.Error("failes to delete alias", sl.Err(err))
			render.JSON(w, r, resp.Error("inernal error"))
			return
		}
		log.Info("alias deleted")
		render.JSON(w, r, "succes deleted")
	}
}
