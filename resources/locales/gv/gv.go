package gv

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type gv struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'gv' locale
func New() locales.Translator {
	return &gv{
		locale:  "gv",
		plurals: []locales.PluralRule{2, 3, 4, 5, 6},
	}
}

// Locale returns the current translators string locale
func (t *gv) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'gv'
func (t *gv) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'gv'
func (t *gv) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	v := locales.V(num)

	if v == 0 && i%10 == 1 {
		return locales.PluralRuleOne, nil
	} else if v == 0 && i%10 == 2 {
		return locales.PluralRuleTwo, nil
	} else if v == 0 && (i%100 == 0 || i%100 == 20 || i%100 == 40 || i%100 == 60 || i%100 == 80) {
		return locales.PluralRuleFew, nil
	} else if v != 0 {
		return locales.PluralRuleMany, nil
	}

	return locales.PluralRuleOther, nil
}
