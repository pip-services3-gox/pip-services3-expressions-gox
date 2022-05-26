package errors

import (
	cerrors "github.com/pip-services3-gox/pip-services3-commons-gox/errors"
	"strconv"
)

// Exception that can be thrown by Mustache Template.
func NewMustacheError(correlationId, code, message string, line, column int) *cerrors.ApplicationError {
	if line != 0 || column != 0 {
		message = message + " at line " + strconv.Itoa(line) + " and column " + strconv.Itoa(column)
	}
	return &cerrors.ApplicationError{
		Category:      cerrors.BadRequest,
		CorrelationId: correlationId,
		Code:          code,
		Message:       message,
		Status:        400,
	}
}
