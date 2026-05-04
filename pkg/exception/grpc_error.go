package exception

import (
	"go-tokopaedi-microservices/protobuf/genproto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func New(code codes.Code, errCode genproto.ErrorCode, message string, violations ...*genproto.FieldViolation) error {
	grpcStatus := status.New(code, message)

	detail := &genproto.ErrorDetail{
		Code:       errCode,
		Message:    message,
		Violations: violations,
	}

	statusWithDetails, err := grpcStatus.WithDetails(detail)
	if err != nil {
		return grpcStatus.Err()
	}

	return statusWithDetails.Err()
}

func ToGrpcError(appErr *ApplicationError) error {
	if appErr == nil {
		return nil
	}

	st := status.New(appErr.GrpcCode, appErr.Message)

	detail := &genproto.ErrorDetail{
		Message: appErr.Message,
		Code:    mapConventionToProto(appErr.ConventionStatusCode),
	}

	stWithDetails, err := st.WithDetails(detail)
	if err != nil {
		return st.Err()
	}

	return stWithDetails.Err()
}

func mapConventionToProto(code string) genproto.ErrorCode {
	switch code {
	case StatusValidationError:
		return genproto.ErrorCode_VALIDATION_ERROR
	case StatusNotFoundError:
		return genproto.ErrorCode_NOT_FOUND
	case StatusDuplicateError:
		return genproto.ErrorCode_CONFLICT
	case StatusAuthError:
		return genproto.ErrorCode_UNAUTHORIZED
	case StatusDatabaseError:
		return genproto.ErrorCode_INTERNAL_ERROR
	default:
		return genproto.ErrorCode_UNKNOWN
	}
}
