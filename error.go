package choruspro

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type APIError struct {
	CodeRetour int32  `json:"codeRetour"`
	Libelle    string `json:"libelle"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("choruspro: error (code %v) : %v", e.CodeRetour, e.Libelle)
}

func parseError(res *http.Response) error {
	apiErr := APIError{}

	err := json.NewDecoder(res.Body).Decode(&apiErr)
	if err != nil {
		return err
	}

	return apiErr
}


type ValidationError struct {
	Errors []ValidationErrorItem `json:"errors"`
}

type ValidationErrorItem struct {
	Code string `json:"code"`
	Message string `json:"message"`
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("choruspro: validation error : %v", e.Errors)
}