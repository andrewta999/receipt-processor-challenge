// models/response.go
package models

type ProcessReceiptResponse struct {
	ID string `json:"id"`
}

type GetPointsResponse struct {
	Points int `json:"points"`
}
