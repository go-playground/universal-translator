package lkt_US

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type lkt_US struct {
	locale string
}

// New returns a new instance of translator for the 'lkt_US' locale
func New() locales.Translator {
	return &lkt_US{
		locale: "lkt_US",
	}
}

// Locale returns the current translators string locale
func (l *lkt_US) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *lkt_US) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
