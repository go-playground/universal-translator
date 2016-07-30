package de_CH

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type de_CH struct {
	locale string
}

// New returns a new instance of translator for the 'de_CH' locale
func New() locales.Translator {
	return &de_CH{
		locale: "de_CH",
	}
}

// Locale returns the current translators string locale
func (l *de_CH) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *de_CH) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	v := locales.V(num)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
