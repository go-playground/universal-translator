package mua

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type mua struct {
	locale string
}

// New returns a new instance of translator for the 'mua' locale
func New() locales.Translator {
	return &mua{
		locale: "mua",
	}
}

// Locale returns the current translators string locale
func (l *mua) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *mua) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
