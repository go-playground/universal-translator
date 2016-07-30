package nl_BQ

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type nl_BQ struct {
	locale string
}

// New returns a new instance of translator for the 'nl_BQ' locale
func New() locales.Translator {
	return &nl_BQ{
		locale: "nl_BQ",
	}
}

// Locale returns the current translators string locale
func (l *nl_BQ) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *nl_BQ) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
