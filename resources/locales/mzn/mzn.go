package mzn

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type mzn struct {
	locale string
}

// New returns a new instance of translator for the 'mzn' locale
func New() locales.Translator {
	return &mzn{
		locale: "mzn",
	}
}

// Locale returns the current translators string locale
func (l *mzn) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *mzn) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
