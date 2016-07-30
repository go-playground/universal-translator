package rw_RW

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type rw_RW struct {
	locale string
}

// New returns a new instance of translator for the 'rw_RW' locale
func New() locales.Translator {
	return &rw_RW{
		locale: "rw_RW",
	}
}

// Locale returns the current translators string locale
func (l *rw_RW) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *rw_RW) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
