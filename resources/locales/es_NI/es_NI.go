package es_NI

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type es_NI struct {
	locale string
}

// New returns a new instance of translator for the 'es_NI' locale
func New() locales.Translator {
	return &es_NI{
		locale: "es_NI",
	}
}

// Locale returns the current translators string locale
func (l *es_NI) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *es_NI) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
