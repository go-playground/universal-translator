package luy

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type luy struct {
	locale string
}

// New returns a new instance of translator for the 'luy' locale
func New() locales.Translator {
	return &luy{
		locale: "luy",
	}
}

// Locale returns the current translators string locale
func (l *luy) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *luy) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
