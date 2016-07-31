package bg_BG

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type bg_BG struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'bg_BG' locale
func New() locales.Translator {
	return &bg_BG{
		locale:  "bg_BG",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *bg_BG) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'bg_BG'
func (t *bg_BG) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'bg_BG'
func (t *bg_BG) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
