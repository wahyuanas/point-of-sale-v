package response

type DataResponse[T any] struct {
	Data T `json:"data"`
}
