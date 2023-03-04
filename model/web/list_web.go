package web

type ListRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ListResponse struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      int64  `json:"user_id"`
}
