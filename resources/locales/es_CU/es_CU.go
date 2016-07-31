package es_CU

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type es_CU struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'es_CU' locale
func New() locales.Translator {
	return &es_CU{
		locale:  "es_CU",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *es_CU) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'es_CU'
func (t *es_CU) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'es_CU'
func (t *es_CU) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
