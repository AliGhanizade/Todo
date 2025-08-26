package todo

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"id_completed"`
	CreateTime  string `json:"created_time"`
}
