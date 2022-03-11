package logic

import (
	"strings"

	"k8scommerce/services/api/client/internal/types"
)

func httpResponse(statusCode int64, message ...string) types.ResponseStatus {
	return types.ResponseStatus{
		StatusCode:    statusCode,
		StatusMessage: strings.Join(message, " "),
	}
}
