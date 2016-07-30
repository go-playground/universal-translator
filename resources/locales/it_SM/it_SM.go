package it_SM

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type it_SM struct {
	locale string
}

// New returns a new instance of translator for the 'it_SM' locale
func New() locales.Translator {
	return &it_SM{
		locale: "it_SM",
	}
}

// Locale returns the current translators string locale
func (l *it_SM) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *it_SM) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
