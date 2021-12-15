package apihandler

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"mongou/internal/app/apihandler/generated/specmodels"
	"mongou/internal/app/apihandler/generated/specops"
)

func (h *Handler) handleGetTasks(params specops.GetTasksParams) middleware.Responder {
	tasks, err := h.service.GetAll(params.HTTPRequest.Context())
	if err != nil {
		return specops.NewGetTasksBadRequest().WithPayload(&specmodels.GenericError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	apiTasks := make([]*specmodels.GetTask, 0, len(tasks))

	for i := range tasks {
		apiTasks = append(apiTasks, &specmodels.GetTask{
			ID:   tasks[i].ID,
			Name: tasks[i].Name,
		})
	}

	return specops.NewGetTasksOK().WithPayload(apiTasks)
}
