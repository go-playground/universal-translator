package es_PR

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type es_PR struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'es_PR' locale
func New() locales.Translator {
	return &es_PR{
		locale:  "es_PR",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *es_PR) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'es_PR'
func (t *es_PR) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'es_PR'
func (t *es_PR) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
