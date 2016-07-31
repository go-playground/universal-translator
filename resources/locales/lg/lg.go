package lg

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type lg struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'lg' locale
func New() locales.Translator {
	return &lg{
		locale:  "lg",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *lg) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'lg'
func (t *lg) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'lg'
func (t *lg) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
