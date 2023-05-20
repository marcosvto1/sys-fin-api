package errorable

const (
	NOT_FOUND_REGISTER           = "not_found_register"
	INTERNAL_ERROR               = "internal_error"
	INVALID_PASSWORD             = "invalid_password"
	INVALID_VALUE_FIELD          = "invalid_value_field"
	INVALID_BODY_REQUEST         = "invalid_body_request"
	FAILED_TO_CREATE_WALLET      = "failed_to_create_wallet"
	FAILED_TO_CREATE_TRANSACTION = "failed_to_create_transaction"
	FAILED_TO_CREATE_CATEGORY    = "failed_to_create_category"
	FAILED_TO_UPDATE_WALLET      = "failed_to_update_wallet"
	FAILED_TO_UPDATE_TRANSACTION = "failed_to_update_transaction"
	FAILED_TO_UPDATE_CATEGORY    = "failed_to_update_category"
)

type CtxError struct {
	Context string `json:"context"`
	Message string `json:"message"`
}

func New(context string) *CtxError {
	return &CtxError{
		Context: context,
		Message: getMessageErrorFromContext(context),
	}
}

func NewHttpError(httpStatus int, errors []CtxError) map[string]interface{} {
	var result map[string]interface{}
	result = map[string]interface{}{
		"status": httpStatus,
		"errors": errors,
	}
	return result
}

func HttpResponse(httpStatus int, errors []error) map[string]interface{} {
	var result map[string]interface{}

	listErr := []map[string]interface{}{}
	for _, err := range errors {
		listErr = append(listErr, map[string]interface{}{
			"context": err.Error(),
			"message": getMessageErrorFromContext(err.Error()),
		})
	}

	result = map[string]interface{}{
		"status": httpStatus,
		"errors": listErr,
	}
	return result
}

func Get(context string) string {
	return getMessageErrorFromContext(context)
}

func getMessageErrorFromContext(context string) string {
	switch context {
	case NOT_FOUND_REGISTER:
		return "Register not found"
	case INTERNAL_ERROR:
		return "Internal error"
	case INVALID_PASSWORD:
		return "Invalid password"
	case INVALID_VALUE_FIELD:
		return "Invalid value field"
	case INVALID_BODY_REQUEST:
		return "invalid body request"
	case FAILED_TO_CREATE_TRANSACTION:
		return "Failed to create transaction"
	case FAILED_TO_CREATE_CATEGORY:
		return "Failed to create category"
	case FAILED_TO_CREATE_WALLET:
		return "Failed to create wallet"
	default:
		return "Not mapped error"
	}
}
