package fr_MG

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type fr_MG struct {
	locale string
}

// New returns a new instance of translator for the 'fr_MG' locale
func New() locales.Translator {
	return &fr_MG{
		locale: "fr_MG",
	}
}

// Locale returns the current translators string locale
func (l *fr_MG) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *fr_MG) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if i == 0 || i == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
