package fi_FI

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type fi_FI struct {
	locale string
}

// New returns a new instance of translator for the 'fi_FI' locale
func New() locales.Translator {
	return &fi_FI{
		locale: "fi_FI",
	}
}

// Locale returns the current translators string locale
func (l *fi_FI) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *fi_FI) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
