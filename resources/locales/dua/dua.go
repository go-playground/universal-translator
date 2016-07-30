package dua

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type dua struct {
	locale string
}

// New returns a new instance of translator for the 'dua' locale
func New() locales.Translator {
	return &dua{
		locale: "dua",
	}
}

// Locale returns the current translators string locale
func (l *dua) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *dua) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
