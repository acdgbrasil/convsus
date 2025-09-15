package convsus

import (
	"fmt"
	"strings"
)

type ERROR_MODULE int
type ERROR_ID int

type Error struct {
	module  ERROR_MODULE
	id      ERROR_ID
	message string
}

func NewError(module ERROR_MODULE, id ERROR_ID, message string, extra ...any) *Error {
	if module < 1 || id < 1 {
		// Avoiding initiation cycle by generating ErrInvalidError without constructor
		return (&Error{
			module:  MODULE_CONVSUS,
			id:      ERROR_INVALID_ERROR,
			message: "[ERROR-%03d%03d] failed to create a new error.",
		}).Format(int(module), int(id))
	}

	return (&Error{
		module:  module,
		id:      id,
		message: message,
	}).Format(extra...)
}

const FORMATTED_EXTRA_PREFIX = "%!(EXTRA"

// Fill extra inforamtion and explicitly demonstrate something was not added
func (e *Error) Format(extra ...any) *Error {
	nm := fmt.Sprintf(e.message, extra...)
	// Find missing format symbols
	formatSymbolIndexes := make([]int, 0)
	extraDataIndex := -1
	words := strings.Split(nm, " ")
	for i, word := range words {
		// Find format symbols
		if len(word) > 0 && word[0] == '%' {
			if strings.Contains(word, FORMATTED_EXTRA_PREFIX) {
				extraDataIndex = i
				continue
			}
			formatSymbolIndexes = append(formatSymbolIndexes, i)
		}
	}
	// Replace format symbols for default missing tag
	for _, index := range formatSymbolIndexes {
		words[index] = "[ missing ]"
	}
	// Ignore extra data
	if extraDataIndex >= 0 {
		// If the prefix of the word is not the FORMATTED_EXTRA_PREFIX. It means an expected word is fused with it.
		// So we remove it as it will be a suffix before removing the rest
		if words[extraDataIndex][:len(FORMATTED_EXTRA_PREFIX)] != FORMATTED_EXTRA_PREFIX {
			words[extraDataIndex], _ = strings.CutSuffix(words[extraDataIndex], FORMATTED_EXTRA_PREFIX)
		}
		words = words[:extraDataIndex]
	}
	// Update error
	e.message = strings.Join(words, " ")
	return e
}

// implement method to be accepeted by the interface error
func (e *Error) Error() string {
	// return formatted string with module and identifier and message
	return fmt.Sprintf("[ERROR-%03d%03d] %s", int(e.module), int(e.id), e.message)
}

func (e *Error) Equals(other error) bool {
	cOther, ok := other.(*Error)
	return ok && e.module == cOther.module && e.id == cOther.id
}

// Modules
const (
	MODULE_CONVSUS ERROR_MODULE = iota + 1
	MODULE_CONECTARAROS
)

// Ids
const (
	ERROR_INVALID_ERROR ERROR_ID = iota + 1
	ERROR_UNMARSHAL_JSON
	ERROR_PARSING_DATA
)

const PACKAGE_ERROR_MODULE = MODULE_CONVSUS

var (
	ErrUnmarshalJson = func(extra error) *Error {
		return NewError(PACKAGE_ERROR_MODULE, ERROR_UNMARSHAL_JSON, "failed to unmarshal json data. %s", extra.Error())
	}
	ErrParsingData = func(extra error) *Error {
		return NewError(PACKAGE_ERROR_MODULE, ERROR_PARSING_DATA, "failed to parse data. %s", extra.Error())
	}
)
