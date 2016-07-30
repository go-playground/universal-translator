package mfe_MU

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type mfe_MU struct {
	locale string
}

// New returns a new instance of translator for the 'mfe_MU' locale
func New() locales.Translator {
	return &mfe_MU{
		locale: "mfe_MU",
	}
}

// Locale returns the current translators string locale
func (l *mfe_MU) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *mfe_MU) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
