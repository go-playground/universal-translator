package rn_BI

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type rn_BI struct {
	locale string
}

// New returns a new instance of translator for the 'rn_BI' locale
func New() locales.Translator {
	return &rn_BI{
		locale: "rn_BI",
	}
}

// Locale returns the current translators string locale
func (l *rn_BI) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *rn_BI) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
