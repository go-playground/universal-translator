package cs

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type cs struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'cs' locale
func New() locales.Translator {
	return &cs{
		locale:  "cs",
		plurals: []locales.PluralRule{2, 4, 5, 6},
	}
}

// Locale returns the current translators string locale
func (t *cs) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'cs'
func (t *cs) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'cs'
func (t *cs) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	v := locales.V(num)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne, nil
	} else if i >= 2 && i <= 4 && v == 0 {
		return locales.PluralRuleFew, nil
	} else if v != 0 {
		return locales.PluralRuleMany, nil
	}

	return locales.PluralRuleOther, nil
}
