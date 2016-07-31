package qu

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type qu struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'qu' locale
func New() locales.Translator {
	return &qu{
		locale:  "qu",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *qu) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'qu'
func (t *qu) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'qu'
func (t *qu) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
