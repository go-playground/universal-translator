package az_Cyrl_AZ

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type az_Cyrl_AZ struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'az_Cyrl_AZ' locale
func New() locales.Translator {
	return &az_Cyrl_AZ{
		locale:  "az_Cyrl_AZ",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *az_Cyrl_AZ) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'az_Cyrl_AZ'
func (t *az_Cyrl_AZ) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'az_Cyrl_AZ'
func (t *az_Cyrl_AZ) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
