package kok_IN

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type kok_IN struct {
	locale string
}

// New returns a new instance of translator for the 'kok_IN' locale
func New() locales.Translator {
	return &kok_IN{
		locale: "kok_IN",
	}
}

// Locale returns the current translators string locale
func (l *kok_IN) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *kok_IN) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
