package dua_CM

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type dua_CM struct {
	locale string
}

// New returns a new instance of translator for the 'dua_CM' locale
func New() locales.Translator {
	return &dua_CM{
		locale: "dua_CM",
	}
}

// Locale returns the current translators string locale
func (l *dua_CM) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *dua_CM) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
