package rest_errors

import (
	"encoding/json"
	"net/http"
)

type RestErr struct {
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Error   string        `json:"error"`
	Causes  []interface{} `json:"causes"`
}

func NewRestErrorFromBytes(bytes []byte) (*RestErr, error) {
	var apiErr RestErr
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, err
	}
	return &apiErr, nil
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewInternalServerError(message string, err error) *RestErr {
	result := &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
	if err != nil {
		result.Causes = append(result.Causes, err.Error())
	}
	return result
}

func NewUnauthorizedError(message string) *RestErr {
	return &RestErr{
		Message: "Unauthorized",
		Status:  http.StatusUnauthorized,
		Error:   "unauthorized",
	}
}

func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusForbidden,
		Error:   "forbidden",
	}
}
