package lkt

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type lkt struct {
	locale string
}

// New returns a new instance of translator for the 'lkt' locale
func New() locales.Translator {
	return &lkt{
		locale: "lkt",
	}
}

// Locale returns the current translators string locale
func (l *lkt) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *lkt) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
