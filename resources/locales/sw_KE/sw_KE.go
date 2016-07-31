package sw_KE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type sw_KE struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'sw_KE' locale
func New() locales.Translator {
	return &sw_KE{
		locale:  "sw_KE",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *sw_KE) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'sw_KE'
func (t *sw_KE) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'sw_KE'
func (t *sw_KE) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
