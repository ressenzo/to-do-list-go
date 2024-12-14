package schemas

type Task struct {
	Id          string
	Description string
}

type TaskResponse struct {
	Id          string `json:"id"`
	Description string `json:"description"`
}
