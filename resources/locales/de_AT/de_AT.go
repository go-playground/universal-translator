package de_AT

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type de_AT struct {
	locale string
}

// New returns a new instance of translator for the 'de_AT' locale
func New() locales.Translator {
	return &de_AT{
		locale: "de_AT",
	}
}

// Locale returns the current translators string locale
func (l *de_AT) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *de_AT) CardinalPluralRule(num string) (locales.PluralRule, error) {

	v := locales.V(num)

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if i == 1 && v == 0 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
