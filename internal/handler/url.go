package handler

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/iwerqfx/url-shortener/internal/model"
	"github.com/iwerqfx/url-shortener/internal/service"
	"net/http"
)

type URLHandler struct {
	*Handler
	service       service.URLService
	serverAddress string
}

func NewURLHandler(handler *Handler, service service.URLService, serverAddress string) *URLHandler {
	return &URLHandler{
		Handler:       handler,
		service:       service,
		serverAddress: serverAddress,
	}
}

type urlCreateRequest struct {
	URL string `json:"url"`
}

type urlCreateResponse struct {
	ShortURL string `json:"short_url"`
}

func (h *URLHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req urlCreateRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		h.newResponse(w, r, http.StatusBadRequest, model.ErrBadRequest)
		return
	}

	alias, err := h.service.Create(req.URL)
	if err != nil {
		h.newResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, urlCreateResponse{
		ShortURL: h.generateShortURL(alias),
	})
}

func (h *URLHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	alias := chi.URLParam(r, "alias")
	if alias == "" {
		h.newResponse(w, r, http.StatusBadRequest, model.ErrBadRequest)
		return
	}

	url, err := h.service.GetByAlias(alias)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, model.ErrURLNotFound) {
			statusCode = http.StatusNotFound
		}

		h.newResponse(w, r, statusCode, err)
		return
	}

	http.Redirect(w, r, url.URL, http.StatusFound)
}

func (h *URLHandler) generateShortURL(alias string) string {
	return fmt.Sprintf("http://%s/%s", h.serverAddress, alias)
}
