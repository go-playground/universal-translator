package qu_EC

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type qu_EC struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'qu_EC' locale
func New() locales.Translator {
	return &qu_EC{
		locale:  "qu_EC",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *qu_EC) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'qu_EC'
func (t *qu_EC) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'qu_EC'
func (t *qu_EC) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
