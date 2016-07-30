package mgh_MZ

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type mgh_MZ struct {
	locale string
}

// New returns a new instance of translator for the 'mgh_MZ' locale
func New() locales.Translator {
	return &mgh_MZ{
		locale: "mgh_MZ",
	}
}

// Locale returns the current translators string locale
func (l *mgh_MZ) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *mgh_MZ) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
