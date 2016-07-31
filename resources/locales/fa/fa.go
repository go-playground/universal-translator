package fa

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type fa struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'fa' locale
func New() locales.Translator {
	return &fa{
		locale:  "fa",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *fa) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'fa'
func (t *fa) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'fa'
func (t *fa) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
