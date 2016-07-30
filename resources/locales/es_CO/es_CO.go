package es_CO

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type es_CO struct {
	locale string
}

// New returns a new instance of translator for the 'es_CO' locale
func New() locales.Translator {
	return &es_CO{
		locale: "es_CO",
	}
}

// Locale returns the current translators string locale
func (l *es_CO) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *es_CO) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
