package entities

type (
	Task struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      string `json:"status"`
	}
)

const (
	StatusTodo       = "To Do"
	StatusInProgress = "In Progress"
	StatusDone       = "Done"
)
