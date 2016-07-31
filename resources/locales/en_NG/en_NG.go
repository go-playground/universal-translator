package en_NG

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type en_NG struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'en_NG' locale
func New() locales.Translator {
	return &en_NG{
		locale:  "en_NG",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *en_NG) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'en_NG'
func (t *en_NG) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'en_NG'
func (t *en_NG) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	v := locales.V(num)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
