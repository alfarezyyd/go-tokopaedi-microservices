package exception

import (
	"fmt"

	"google.golang.org/grpc/codes"
)

type ApplicationError struct {
	HttpStatusCode       int                    `json:"status_code"`
	ConventionStatusCode string                 `json:"convention_status_code"`
	Message              string                 `json:"message"`
	Details              map[string]interface{} `json:"details"`
	GrpcCode             codes.Code             `json:"grpc_code"`
}

func (applicationError *ApplicationError) Error() string {
	return fmt.Sprintf("Error %d-%s: %s", applicationError.HttpStatusCode, applicationError.ConventionStatusCode, applicationError.Message)
}

func NewApplicationError(statusCode int, message string) *ApplicationError {
	return &ApplicationError{
		HttpStatusCode: statusCode,
		Message:        message,
	}
}

func NewApplicationErrorWithDetails(statusCode int, message string, details map[string]interface{}) *ApplicationError {
	return &ApplicationError{
		HttpStatusCode: statusCode,
		Message:        message,
		Details:        details,
	}
}

func NewApplicationErrorSpecific(statusCode int, conventionStatusCode string, message string, details map[string]interface{}) *ApplicationError {
	return &ApplicationError{
		HttpStatusCode:       statusCode,
		ConventionStatusCode: conventionStatusCode,
		Message:              message,
		Details:              details,
	}
}

func ThrowApplicationError(applicationError *ApplicationError) {
	panic(applicationError)
}

func mapHttpToGrpcCode(httpCode int) codes.Code {
	switch httpCode {
	case 400:
		return codes.InvalidArgument
	case 401:
		return codes.Unauthenticated
	case 403:
		return codes.PermissionDenied
	case 404:
		return codes.NotFound
	case 409:
		return codes.AlreadyExists
	case 429:
		return codes.ResourceExhausted
	case 504:
		return codes.DeadlineExceeded
	default:
		return codes.Internal
	}
}
