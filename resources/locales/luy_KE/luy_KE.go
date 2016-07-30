package luy_KE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type luy_KE struct {
	locale string
}

// New returns a new instance of translator for the 'luy_KE' locale
func New() locales.Translator {
	return &luy_KE{
		locale: "luy_KE",
	}
}

// Locale returns the current translators string locale
func (l *luy_KE) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *luy_KE) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
