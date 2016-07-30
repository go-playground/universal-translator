package luo_KE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type luo_KE struct {
	locale string
}

// New returns a new instance of translator for the 'luo_KE' locale
func New() locales.Translator {
	return &luo_KE{
		locale: "luo_KE",
	}
}

// Locale returns the current translators string locale
func (l *luo_KE) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *luo_KE) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
