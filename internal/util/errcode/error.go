package errcode

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/samber/lo"
	"runtime/debug"
	"strings"
)

type Code int

const (
	CodeUnknown Code = iota
	CodeInvalidArgument
	CodeNotFound
	CodeAlreadyExists
	CodeInternal
	CodeUnimplemented
	CodeDeadlineExceeded
	CodeContextCanceled
	CodeDataLoss
	CodeFailedPrecondition
	CodeTooManyRequests
	CodeOutOfRange
	CodePermissionDenied
	CodeResourceExhausted
	CodeUnauthenticated
)

func (c Code) String() string {
	switch c {
	case CodeUnknown:
		return "Unknown"
	case CodeInvalidArgument:
		return "InvalidArgument"
	case CodeNotFound:
		return "NotFound"
	case CodeAlreadyExists:
		return "AlreadyExists"
	case CodeInternal:
		return "Internal"
	case CodeUnimplemented:
		return "Unimplemented"
	case CodeDeadlineExceeded:
		return "DeadlineExceeded"
	case CodeDataLoss:
		return "DataLoss"
	case CodeFailedPrecondition:
		return "FailedPrecondition"
	case CodeTooManyRequests:
		return "TooManyRequests"
	case CodeOutOfRange:
		return "OutOfRange"
	case CodePermissionDenied:
		return "PermissionDenied"
	case CodeResourceExhausted:
		return "ResourceExhausted"
	case CodeUnauthenticated:
		return "Unauthenticated"
	case CodeContextCanceled:
		return "ContextCanceled"
	}

	return fmt.Sprintf("Unknown: %d", c)
}

type Error struct {
	Code   Code
	origin error
	stack  string
}

type Errors []Error

func (e Errors) Error() string {
	msgs := lo.Map(e, func(_ Error, i int) string {
		return e[i].Error()
	})
	return strings.Join(msgs, ", ")
}

func (e *Error) Error() string {
	if e.origin == nil {
		return fmt.Sprintf("%s: StackTrace:\n%s", e.Code.String(), e.stack)
	}
	return fmt.Sprintf("%s: %s\nStackTrace:\n%s", e.Code.String(), e.origin.Error(), e.stack)
}

func New(err error) error {
	if err == nil {
		return nil
	}
	// if err is already Error type, nothing.
	var e *Error
	if errors.As(err, &e) {
		return err
	}

	newerr := &Error{
		Code:   CodeInternal,
		origin: err,
		stack:  string(debug.Stack()),
	}

	// check validation error
	var verr validator.ValidationErrors
	if errors.As(err, &verr) {
		newerr.Code = CodeInvalidArgument
		return newerr
	}

	// context canceled
	if errors.Is(err, context.Canceled) {
		newerr.Code = CodeContextCanceled
		return newerr
	}

	// deadline exceeded
	if errors.Is(err, context.DeadlineExceeded) {
		newerr.Code = CodeDeadlineExceeded
		return newerr
	}
	return newerr
}

func IsUnknown(err error) bool {
	var e *Error
	return errors.As(err, &e) && e.Code == CodeUnknown
}

func IsInvalidArgument(err error) bool {
	var e *Error
	return errors.As(err, &e) && e.Code == CodeInvalidArgument
}

func IsNotFound(err error) bool {
	var e *Error
	return errors.As(err, &e) && e.Code == CodeNotFound
}

func IsAlreadyExists(err error) bool {
	var e *Error
	return errors.As(err, &e) && e.Code == CodeAlreadyExists
}

func IsInternal(err error) bool {
	var e *Error
	return errors.As(err, &e) && e.Code == CodeInternal
}

func IsDeadlineExceeded(err error) bool {
	var e *Error
	return errors.As(err, &e) && e.Code == CodeDeadlineExceeded
}

func IsContextCanceled(err error) bool {
	var e *Error
	return errors.As(err, &e) && e.Code == CodeContextCanceled
}

func IsServerError(err error) bool {
	return IsUnknown(err) || IsInternal(err) || IsDeadlineExceeded(err)
}

func NewInvalidArgument(format string, a ...interface{}) error {
	stack := debug.Stack()
	return &Error{
		Code:   CodeInvalidArgument,
		origin: fmt.Errorf(format, a...),
		stack:  string(stack),
	}
}

func NewNotFound(format string, a ...interface{}) error {
	stack := debug.Stack()
	return &Error{
		Code:   CodeNotFound,
		origin: fmt.Errorf(format, a...),
		stack:  string(stack),
	}
}

func NewInternal(format string, a ...interface{}) error {
	stack := debug.Stack()
	return &Error{
		Code:   CodeInternal,
		origin: fmt.Errorf(format, a...),
		stack:  string(stack),
	}
}

func NewAlreadyExists(format string, a ...interface{}) error {
	stack := debug.Stack()
	return &Error{
		Code:   CodeAlreadyExists,
		origin: fmt.Errorf(format, a...),
		stack:  string(stack),
	}
}
