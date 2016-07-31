package mk_MK

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type mk_MK struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'mk_MK' locale
func New() locales.Translator {
	return &mk_MK{
		locale:  "mk_MK",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *mk_MK) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'mk_MK'
func (t *mk_MK) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'mk_MK'
func (t *mk_MK) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	v := locales.V(num)

	f, err := locales.F(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if (v == 0 && i%10 == 1) || (f%10 == 1) {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
