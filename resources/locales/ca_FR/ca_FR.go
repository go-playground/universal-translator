package ca_FR

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ca_FR struct {
	locale string
}

// New returns a new instance of translator for the 'ca_FR' locale
func New() locales.Translator {
	return &ca_FR{
		locale: "ca_FR",
	}
}

// Locale returns the current translators string locale
func (l *ca_FR) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ca_FR) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
