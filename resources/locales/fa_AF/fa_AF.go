package fa_AF

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type fa_AF struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'fa_AF' locale
func New() locales.Translator {
	return &fa_AF{
		locale:  "fa_AF",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *fa_AF) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'fa_AF'
func (t *fa_AF) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'fa_AF'
func (t *fa_AF) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
