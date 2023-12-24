package utils

// SuccessResp struct is an API response returned on successful API calls
type SuccessResp struct {
	Message string `json:"message" example:"Record added successfully"`
}

// ErrResp struct is an API response returned on failed API calls
type ErrResp struct {
	Error   string `json:"error" example:"record not found"`
	Message string `json:"message" example:"Record not found"`
}
