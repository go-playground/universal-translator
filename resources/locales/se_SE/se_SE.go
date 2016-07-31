package se_SE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type se_SE struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'se_SE' locale
func New() locales.Translator {
	return &se_SE{
		locale:  "se_SE",
		plurals: []locales.PluralRule{2, 3, 6},
	}
}

// Locale returns the current translators string locale
func (t *se_SE) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'se_SE'
func (t *se_SE) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'se_SE'
func (t *se_SE) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	} else if n == 2 {
		return locales.PluralRuleTwo, nil
	}

	return locales.PluralRuleOther, nil
}
