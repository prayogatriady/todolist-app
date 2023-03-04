package web

type TaskRequest struct {
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type TaskResponse struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
	// DueDate     time.Time `json:"due_date"`
	Completed bool  `json:"completed"`
	ListID    int64 `json:"list_id"`
	UserID    int64 `json:"user_id"`
}
