package ln_AO

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ln_AO struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ln_AO' locale
func New() locales.Translator {
	return &ln_AO{
		locale:  "ln_AO",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ln_AO) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ln_AO'
func (t *ln_AO) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ln_AO'
func (t *ln_AO) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n >= 0 && n <= 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
