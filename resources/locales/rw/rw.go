package rw

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type rw struct {
	locale string
}

// New returns a new instance of translator for the 'rw' locale
func New() locales.Translator {
	return &rw{
		locale: "rw",
	}
}

// Locale returns the current translators string locale
func (l *rw) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *rw) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
