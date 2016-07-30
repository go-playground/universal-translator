package luo

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type luo struct {
	locale string
}

// New returns a new instance of translator for the 'luo' locale
func New() locales.Translator {
	return &luo{
		locale: "luo",
	}
}

// Locale returns the current translators string locale
func (l *luo) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *luo) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
