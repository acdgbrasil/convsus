package conectararos

import (
	"github.com/acdgbrasil/convsus"
)

// this package module error
const PACKAGE_ERROR_MODULE = convsus.MODULE_CONECTARAROS

// error ids
const (
	ERROR_VALIDATION convsus.ERROR_ID = iota + 1
)

var ErrValidation = func(extra string) *convsus.Error {
	return convsus.NewError(PACKAGE_ERROR_MODULE, ERROR_VALIDATION, "error trying to validate field. %s", extra)
}
