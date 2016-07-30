package fr_MQ

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type fr_MQ struct {
	locale string
}

// New returns a new instance of translator for the 'fr_MQ' locale
func New() locales.Translator {
	return &fr_MQ{
		locale: "fr_MQ",
	}
}

// Locale returns the current translators string locale
func (l *fr_MQ) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *fr_MQ) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if i == 0 || i == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
