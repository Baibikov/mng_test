package apihandler

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"mongou/internal/app/apihandler/generated/specmodels"
	"mongou/internal/app/apihandler/generated/specops"
)

func (h *Handler) handleGetTask(params specops.GetTaskParams) middleware.Responder {
	task, err := h.service.Get(params.HTTPRequest.Context(), params.ID)
	if err != nil {
		return specops.NewGetTaskBadRequest().WithPayload(&specmodels.GenericError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return specops.NewGetTaskOK().WithPayload(&specmodels.GetTask{
		ID:   task.ID,
		Name: task.Name,
	})
}
