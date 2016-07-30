package mgh

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type mgh struct {
	locale string
}

// New returns a new instance of translator for the 'mgh' locale
func New() locales.Translator {
	return &mgh{
		locale: "mgh",
	}
}

// Locale returns the current translators string locale
func (l *mgh) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *mgh) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
