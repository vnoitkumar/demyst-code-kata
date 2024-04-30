package configurations

type Config struct {
	TodoURL       string `json:"todo_url" validate:"required"`
	TodoListSize  int    `json:"todo_list_size" validate:"required"`
	TodoChunkSize int    `json:"todo_chunk_size" validate:"required"`
}
