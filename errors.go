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
var _ error = new(ErrMissingPluralTranslation)
var _ error = new(ErrExistingTranslator)

// ErrExistingTranslator is the error representing a conflicting translator
type ErrExistingTranslator struct {
	locale string
}

// Error returns ErrExistingTranslator's internal error text
func (e *ErrExistingTranslator) Error() string {
	return fmt.Sprintf("error: conflicting translator for locale '%s'", e.locale)
}

// ErrConflictingTranslation is the error representing a conflicting translation
type ErrConflictingTranslation struct {
	key  interface{}
	rule locales.PluralRule
	text string
}

// Error returns ErrConflictingTranslation's internal error text
func (e *ErrConflictingTranslation) Error() string {

	if _, ok := e.key.(string); !ok {
		return fmt.Sprintf("error: conflicting key '%#v' rule '%s' with text '%s', value being ignored", e.key, e.rule, e.text)
	}

	return fmt.Sprintf("error: conflicting key '%s' rule '%s' with text '%s', value being ignored", e.key, e.rule, e.text)
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

// ErrMissingPluralTranslation is the error signifying a missing translation given
// the locales plural rules.
type ErrMissingPluralTranslation struct {
	key             interface{}
	rule            locales.PluralRule
	translationType string
}

// Error returns ErrMissingPluralTranslation's internal error text
func (e *ErrMissingPluralTranslation) Error() string {

	if _, ok := e.key.(string); !ok {
		return fmt.Sprintf("error: missing '%s' plural rule '%s' for translation with key '%#v'", e.translationType, e.rule, e.key)
	}

	return fmt.Sprintf("error: missing '%s' plural rule '%s' for translation with key '%s'", e.translationType, e.rule, e.key)
}

// ErrMissingBracket is the error representing a missing bracket in a translation
// eg. This is a {0 <-- missing ending '}'
type ErrMissingBracket struct {
}

// Error returns ErrMissingBracket error message
func (e *ErrMissingBracket) Error() string {
	return fmt.Sprint("error: missing bracket, '{' or '}', in translation")
}

// ErrBadParamSyntax is the error representing a bad parameter definition in a translation
// eg. This is a {must-be-int}
type ErrBadParamSyntax struct {
	param string
}

// Error returns ErrBadParamSyntax error message
func (e *ErrBadParamSyntax) Error() string {
	return fmt.Sprintf("error: bad parameter syntax, missing parameter '%s'", e.param)
}

// import/export errors

// ErrMissingLocale is the error representing an expected locale that could
// not be found aka locale not registered with the UniversalTranslator Instance
type ErrMissingLocale struct {
	locale string
}

// Error returns ErrMissingLocale's internal error text
func (e *ErrMissingLocale) Error() string {
	return fmt.Sprintf("error: locale '%s' not registered.", e.locale)
}

// ErrBadPluralDefinition is the error representing an incorrect plural definition
// usually found within translations defined within files during the import process.
type ErrBadPluralDefinition struct {
	tl translation
}

// Error returns ErrBadPluralDefinition's internal error text
func (e *ErrBadPluralDefinition) Error() string {
	return fmt.Sprintf("error: bad plural definition '%#v'", e.tl)
}
