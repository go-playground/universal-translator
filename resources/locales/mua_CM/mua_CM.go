package mua_CM

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type mua_CM struct {
	locale string
}

// New returns a new instance of translator for the 'mua_CM' locale
func New() locales.Translator {
	return &mua_CM{
		locale: "mua_CM",
	}
}

// Locale returns the current translators string locale
func (l *mua_CM) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *mua_CM) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
