package ig

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ig struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ig' locale
func New() locales.Translator {
	return &ig{
		locale:  "ig",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *ig) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ig'
func (t *ig) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ig'
func (t *ig) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
