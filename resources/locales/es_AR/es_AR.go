package es_AR

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type es_AR struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'es_AR' locale
func New() locales.Translator {
	return &es_AR{
		locale:  "es_AR",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *es_AR) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'es_AR'
func (t *es_AR) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'es_AR'
func (t *es_AR) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
