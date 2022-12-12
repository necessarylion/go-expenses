package request

type ListingRequest struct {
	Limit interface{} `json:"limit" validate:"required"`
	Page  interface{} `json:"page" validate:"required"`
}
