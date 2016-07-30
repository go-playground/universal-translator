package bm

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type bm struct {
	locale string
}

// New returns a new instance of translator for the 'bm' locale
func New() locales.Translator {
	return &bm{
		locale: "bm",
	}
}

// Locale returns the current translators string locale
func (l *bm) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *bm) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
