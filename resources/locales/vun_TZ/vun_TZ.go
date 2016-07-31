package vun_TZ

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type vun_TZ struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'vun_TZ' locale
func New() locales.Translator {
	return &vun_TZ{
		locale:  "vun_TZ",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *vun_TZ) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'vun_TZ'
func (t *vun_TZ) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'vun_TZ'
func (t *vun_TZ) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
