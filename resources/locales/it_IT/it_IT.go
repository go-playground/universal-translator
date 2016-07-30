package it_IT

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type it_IT struct {
	locale string
}

// New returns a new instance of translator for the 'it_IT' locale
func New() locales.Translator {
	return &it_IT{
		locale: "it_IT",
	}
}

// Locale returns the current translators string locale
func (l *it_IT) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *it_IT) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
