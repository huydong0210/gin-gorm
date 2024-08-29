package request

type TodoItemUpdateRequest struct {
	Name  string `json:"name"`
	State string `json:"state"`
}

type TodoItemCreateRequest struct {
	Name  string `json:"name"`
	State string `json:"state"`
}
