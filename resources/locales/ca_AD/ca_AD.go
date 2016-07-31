package ca_AD

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ca_AD struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ca_AD' locale
func New() locales.Translator {
	return &ca_AD{
		locale:  "ca_AD",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ca_AD) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ca_AD'
func (t *ca_AD) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ca_AD'
func (t *ca_AD) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
