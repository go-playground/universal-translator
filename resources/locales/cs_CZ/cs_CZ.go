package cs_CZ

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type cs_CZ struct {
	locale string
}

// New returns a new instance of translator for the 'cs_CZ' locale
func New() locales.Translator {
	return &cs_CZ{
		locale: "cs_CZ",
	}
}

// Locale returns the current translators string locale
func (l *cs_CZ) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *cs_CZ) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	v := locales.V(num)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne, nil
	} else if i >= 2 && i <= 4 && v == 0 {
		return locales.PluralRuleFew, nil
	} else if v != 0 {
		return locales.PluralRuleMany, nil
	}

	return locales.PluralRuleOther, nil
}
