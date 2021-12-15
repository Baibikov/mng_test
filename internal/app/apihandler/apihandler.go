package apihandler

import (
	"net/http"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"

	"mongou/internal/app/apihandler/generated"
	"mongou/internal/app/apihandler/generated/specmodels"
	"mongou/internal/app/apihandler/generated/specops"
)

//go:generate swagger -q generate server -f ../../../api/mongou.yaml -m generated/specmodels -s generated -a specops  -C ../../../goswagger/config.yaml  --template-dir ../../../goswagger/template
type Handler struct {
	api *specops.TaskAPIAPI
}

func (h *Handler) Serve(middleware middleware.Builder) http.Handler {
	return h.api.Serve(middleware)
}

func New() (*Handler, error) {
	swaggerSpec, err := loads.Embedded(generated.SwaggerJSON, generated.FlatSwaggerJSON)
	if err != nil {
		return nil, errors.Wrap(err, "loads.Embedded")
	}

	api := specops.NewTaskAPIAPI(swaggerSpec)

	api.GetTaskHandler = specops.GetTaskHandlerFunc(func(params specops.GetTaskParams) middleware.Responder {
		return specops.NewGetTaskOK().WithPayload(&specmodels.GetTask{
			ID:   params.ID,
			Name: "name",
		})

	})

	handler := &Handler{
		api: api,
	}
	return handler, nil
}
