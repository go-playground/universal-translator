package es_GQ

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type es_GQ struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'es_GQ' locale
func New() locales.Translator {
	return &es_GQ{
		locale:  "es_GQ",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *es_GQ) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'es_GQ'
func (t *es_GQ) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'es_GQ'
func (t *es_GQ) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
