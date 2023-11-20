package choruspro

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// APIError is the error returned by the API
type APIError struct {
	CodeRetour int32  `json:"codeRetour"`
	Libelle    string `json:"libelle"`
}

// Error returns the error message
func (e APIError) Error() string {
	return fmt.Sprintf("choruspro: error (code %v) : %v", e.CodeRetour, e.Libelle)
}

// parseError parses the error returned by the API
func parseError(res *http.Response) error {
	apiErr := APIError{}

	err := json.NewDecoder(res.Body).Decode(&apiErr)
	if err != nil {
		return err
	}

	return apiErr
}
