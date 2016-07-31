package bn_BD

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type bn_BD struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'bn_BD' locale
func New() locales.Translator {
	return &bn_BD{
		locale:  "bn_BD",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *bn_BD) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'bn_BD'
func (t *bn_BD) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'bn_BD'
func (t *bn_BD) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
