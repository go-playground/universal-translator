package cu_RU

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type cu_RU struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'cu_RU' locale
func New() locales.Translator {
	return &cu_RU{
		locale:  "cu_RU",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *cu_RU) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'cu_RU'
func (t *cu_RU) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'cu_RU'
func (t *cu_RU) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
