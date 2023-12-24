package models

type (
	TaskRequest struct {
		Title       string `json:"title" validate:"required"`
		Description string `json:"description" validate:"required"`
	}
	UpdateStatusRequest struct {
		Status string `json:"status" validate:"require"`
	}

	TaskResponse struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      string `json:"status"`
	}
)
