package kw_GB

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type kw_GB struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'kw_GB' locale
func New() locales.Translator {
	return &kw_GB{
		locale:  "kw_GB",
		plurals: []locales.PluralRule{2, 3, 6},
	}
}

// Locale returns the current translators string locale
func (t *kw_GB) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'kw_GB'
func (t *kw_GB) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'kw_GB'
func (t *kw_GB) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
