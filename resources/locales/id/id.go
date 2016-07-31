package id

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type id struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'id' locale
func New() locales.Translator {
	return &id{
		locale:  "id",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *id) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'id'
func (t *id) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'id'
func (t *id) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
