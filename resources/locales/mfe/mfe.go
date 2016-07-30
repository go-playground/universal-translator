package mfe

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type mfe struct {
	locale string
}

// New returns a new instance of translator for the 'mfe' locale
func New() locales.Translator {
	return &mfe{
		locale: "mfe",
	}
}

// Locale returns the current translators string locale
func (l *mfe) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *mfe) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
