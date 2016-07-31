package zgh_MA

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type zgh_MA struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'zgh_MA' locale
func New() locales.Translator {
	return &zgh_MA{
		locale:  "zgh_MA",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *zgh_MA) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'zgh_MA'
func (t *zgh_MA) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'zgh_MA'
func (t *zgh_MA) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
