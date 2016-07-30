package asa_TZ

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type asa_TZ struct {
	locale string
}

// New returns a new instance of translator for the 'asa_TZ' locale
func New() locales.Translator {
	return &asa_TZ{
		locale: "asa_TZ",
	}
}

// Locale returns the current translators string locale
func (l *asa_TZ) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *asa_TZ) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
