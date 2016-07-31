package vai_Latn

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type vai_Latn struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'vai_Latn' locale
func New() locales.Translator {
	return &vai_Latn{
		locale:  "vai_Latn",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *vai_Latn) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'vai_Latn'
func (t *vai_Latn) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'vai_Latn'
func (t *vai_Latn) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
