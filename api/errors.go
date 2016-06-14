package api

// Errors to include in any JSON response object
type Errors struct {
	List []struct {
		ID      string `json:"id"`
		Message string `json:"message"`
	} `json:"errors"`
}
