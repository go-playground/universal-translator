package mer

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type mer struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'mer' locale
func New() locales.Translator {
	return &mer{
		locale:  "mer",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *mer) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'mer'
func (t *mer) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'mer'
func (t *mer) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
