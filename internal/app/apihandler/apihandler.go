package apihandler

import (
	"net/http"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"

	"mongou/internal/app/apihandler/generated"
	"mongou/internal/app/apihandler/generated/specops"
	"mongou/internal/app/service"
)

//go:generate swagger -q generate server -f ../../../api/mongou.yaml -m generated/specmodels -s generated -a specops  -C ../../../goswagger/config.yaml  --template-dir ../../../goswagger/template
type Handler struct {
	api     *specops.TaskAPIAPI
	service *service.UseCase
}

func (h *Handler) Serve(middleware middleware.Builder) http.Handler {
	return h.api.Serve(middleware)
}

func New(service *service.UseCase) (*Handler, error) {
	swaggerSpec, err := loads.Embedded(generated.SwaggerJSON, generated.FlatSwaggerJSON)
	if err != nil {
		return nil, errors.Wrap(err, "loads.Embedded")
	}

	handler := &Handler{
		api:     specops.NewTaskAPIAPI(swaggerSpec),
		service: service,
	}

	handler.api.GetTaskHandler = specops.GetTaskHandlerFunc(handler.handleGetTask)
	handler.api.GetTasksHandler = specops.GetTasksHandlerFunc(handler.handleGetTasks)

	return handler, nil
}
