package dto

type Response[T any] struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Data T `json:"data"`
}

func CreateResponsError(message string) Response[string] {
	return  Response[string]{
		Code: "99",
		Message: message,
		Data: "",
	}
}

func  CreateResponsSucces[T any](data T) Response[T] {
	return  Response[T]{
		Code: "99",
		Message: "succses",
		Data: data,
	}
}

func CreateResponsErrorData(message string, data map[string]string) Response[map[string]string] {
	return  Response[map[string]string]{
		Code: "99",
		Message: message,
		Data: data,
	}
}

