package dz_BT

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type dz_BT struct {
	locale string
}

// New returns a new instance of translator for the 'dz_BT' locale
func New() locales.Translator {
	return &dz_BT{
		locale: "dz_BT",
	}
}

// Locale returns the current translators string locale
func (l *dz_BT) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *dz_BT) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
