package ca_IT

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ca_IT struct {
	locale string
}

// New returns a new instance of translator for the 'ca_IT' locale
func New() locales.Translator {
	return &ca_IT{
		locale: "ca_IT",
	}
}

// Locale returns the current translators string locale
func (l *ca_IT) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ca_IT) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
