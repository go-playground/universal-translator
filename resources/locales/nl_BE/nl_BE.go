package nl_BE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type nl_BE struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'nl_BE' locale
func New() locales.Translator {
	return &nl_BE{
		locale:  "nl_BE",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *nl_BE) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'nl_BE'
func (t *nl_BE) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'nl_BE'
func (t *nl_BE) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
