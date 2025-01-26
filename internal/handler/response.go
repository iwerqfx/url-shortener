package handler

import (
	"github.com/go-chi/render"
	"net/http"
)

func (h *Handler) newResponse(w http.ResponseWriter, r *http.Request, statusCode int, err error) {
	if err != nil {
		if statusCode >= http.StatusInternalServerError {
			h.log.Error(err.Error())
		}

		render.Status(r, statusCode)
		w.Write([]byte(err.Error()))
		return
	}

	render.Status(r, statusCode)
	w.Write([]byte(http.StatusText(statusCode)))
}
