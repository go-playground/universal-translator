package gu

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type gu struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'gu' locale
func New() locales.Translator {
	return &gu{
		locale:  "gu",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *gu) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'gu'
func (t *gu) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'gu'
func (t *gu) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
