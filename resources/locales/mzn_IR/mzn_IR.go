package mzn_IR

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type mzn_IR struct {
	locale string
}

// New returns a new instance of translator for the 'mzn_IR' locale
func New() locales.Translator {
	return &mzn_IR{
		locale: "mzn_IR",
	}
}

// Locale returns the current translators string locale
func (l *mzn_IR) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *mzn_IR) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
