package lb_LU

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type lb_LU struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'lb_LU' locale
func New() locales.Translator {
	return &lb_LU{
		locale:  "lb_LU",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *lb_LU) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'lb_LU'
func (t *lb_LU) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'lb_LU'
func (t *lb_LU) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
