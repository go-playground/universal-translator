package nl_CW

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type nl_CW struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'nl_CW' locale
func New() locales.Translator {
	return &nl_CW{
		locale:  "nl_CW",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *nl_CW) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'nl_CW'
func (t *nl_CW) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'nl_CW'
func (t *nl_CW) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
