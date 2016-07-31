package sq_MK

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type sq_MK struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'sq_MK' locale
func New() locales.Translator {
	return &sq_MK{
		locale:  "sq_MK",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *sq_MK) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'sq_MK'
func (t *sq_MK) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'sq_MK'
func (t *sq_MK) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
