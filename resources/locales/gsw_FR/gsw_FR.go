package gsw_FR

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type gsw_FR struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'gsw_FR' locale
func New() locales.Translator {
	return &gsw_FR{
		locale:  "gsw_FR",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *gsw_FR) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'gsw_FR'
func (t *gsw_FR) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'gsw_FR'
func (t *gsw_FR) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
