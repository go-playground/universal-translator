package ewo_CM

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ewo_CM struct {
	locale string
}

// New returns a new instance of translator for the 'ewo_CM' locale
func New() locales.Translator {
	return &ewo_CM{
		locale: "ewo_CM",
	}
}

// Locale returns the current translators string locale
func (l *ewo_CM) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ewo_CM) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
