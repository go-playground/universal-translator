package se_NO

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type se_NO struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'se_NO' locale
func New() locales.Translator {
	return &se_NO{
		locale:  "se_NO",
		plurals: []locales.PluralRule{2, 3, 6},
	}
}

// Locale returns the current translators string locale
func (t *se_NO) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'se_NO'
func (t *se_NO) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'se_NO'
func (t *se_NO) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
