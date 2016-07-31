package tr_TR

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type tr_TR struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'tr_TR' locale
func New() locales.Translator {
	return &tr_TR{
		locale:  "tr_TR",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *tr_TR) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'tr_TR'
func (t *tr_TR) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'tr_TR'
func (t *tr_TR) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
