package es_VE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type es_VE struct {
	locale string
}

// New returns a new instance of translator for the 'es_VE' locale
func New() locales.Translator {
	return &es_VE{
		locale: "es_VE",
	}
}

// Locale returns the current translators string locale
func (l *es_VE) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *es_VE) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
