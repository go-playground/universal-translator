package gd

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type gd struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'gd' locale
func New() locales.Translator {
	return &gd{
		locale:  "gd",
		plurals: []locales.PluralRule{2, 3, 4, 6},
	}
}

// Locale returns the current translators string locale
func (t *gd) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'gd'
func (t *gd) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'gd'
func (t *gd) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 || n == 11 {
		return locales.PluralRuleOne, nil
	} else if n == 2 || n == 12 {
		return locales.PluralRuleTwo, nil
	} else if (n >= 3 && n <= 10) || (n >= 13 && n <= 19) {
		return locales.PluralRuleFew, nil
	}

	return locales.PluralRuleOther, nil
}
