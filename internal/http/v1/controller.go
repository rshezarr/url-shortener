package v1

import (
	"github.com/go-chi/chi/v5"
	"url-short/internal/service"
)

type Controller struct {
	svc     *service.Service
	handler *chi.Mux
}

func NewController(svc *service.Service) *Controller {
	return &Controller{
		svc:     svc,
		handler: chi.NewRouter(),
	}
}

func (c *Controller) InitRoutes() *chi.Mux {
	return c.handler
}
