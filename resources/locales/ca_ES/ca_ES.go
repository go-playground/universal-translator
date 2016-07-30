package ca_ES

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ca_ES struct {
	locale string
}

// New returns a new instance of translator for the 'ca_ES' locale
func New() locales.Translator {
	return &ca_ES{
		locale: "ca_ES",
	}
}

// Locale returns the current translators string locale
func (l *ca_ES) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ca_ES) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
