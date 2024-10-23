package models

// JsonResponse defines the structure of the JSON response.
type JsonResponse struct {
    Error   bool        `json:"error"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"` // Use `interface{}` for any type of data
}
