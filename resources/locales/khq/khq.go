package khq

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type khq struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'khq' locale
func New() locales.Translator {
	return &khq{
		locale:  "khq",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *khq) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'khq'
func (t *khq) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'khq'
func (t *khq) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
