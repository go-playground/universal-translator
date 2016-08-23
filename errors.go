package ut

import (
	"errors"
	"fmt"

	"github.com/go-playground/locales"
)

var (
	// ErrUnknowTranslation indicates the translation could not be found
	ErrUnknowTranslation = errors.New("Unknown Translation")
)

var _ error = new(ErrConflictingTranslation)
var _ error = new(ErrRangeTranslation)
var _ error = new(ErrOrdinalTranslation)
var _ error = new(ErrCardinalTranslation)
var _ error = new(ErrLocaleNotFound)
var _ error = new(ErrMissingPluralTranslation)

// ErrConflictingTranslation is the error representing a conflicting translation
type ErrConflictingTranslation struct {
	key  interface{}
	rule locales.PluralRule
	text string
}

// Error returns ErrConflictingTranslation's internal error text
func (e *ErrConflictingTranslation) Error() string {
	return fmt.Sprintf("warning: conflicting key '%#v' rule '%d' with text '%s', value being ignored", e.key, e.rule, e.text)
}

// ErrRangeTranslation is the error representing a range translation error
type ErrRangeTranslation struct {
	text string
}

// Error returns ErrRangeTranslation's internal error text
func (e *ErrRangeTranslation) Error() string {
	return e.text
}

// ErrOrdinalTranslation is the error representing an ordinal translation error
type ErrOrdinalTranslation struct {
	text string
}

// Error returns ErrOrdinalTranslation's internal error text
func (e *ErrOrdinalTranslation) Error() string {
	return e.text
}

// ErrCardinalTranslation is the error representing a cardinal translation error
type ErrCardinalTranslation struct {
	text string
}

// Error returns ErrCardinalTranslation's internal error text
func (e *ErrCardinalTranslation) Error() string {
	return e.text
}

// ErrLocaleNotFound is the error signifying a locale which could not be found
type ErrLocaleNotFound struct {
	text string
}

// Error returns ErrLocaleNotFound's internal error text
func (e *ErrLocaleNotFound) Error() string {
	return e.text
}

// ErrMissingPluralTranslation is the error signifying a missing translation given
// the locales plural rules.
type ErrMissingPluralTranslation struct {
	key             interface{}
	rule            locales.PluralRule
	translationType string
}

// Error returns ErrMissingPluralTranslation's internal error text
func (e *ErrMissingPluralTranslation) Error() string {
	return fmt.Sprintf("error: missing %s plural rule '%s' for translation with key '%#v", e.translationType, e.rule, e.key)
}
