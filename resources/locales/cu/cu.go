package cu

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type cu struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'cu' locale
func New() locales.Translator {
	return &cu{
		locale:  "cu",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *cu) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'cu'
func (t *cu) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'cu'
func (t *cu) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
