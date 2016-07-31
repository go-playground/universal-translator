package en_SH

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type en_SH struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'en_SH' locale
func New() locales.Translator {
	return &en_SH{
		locale:  "en_SH",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *en_SH) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'en_SH'
func (t *en_SH) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'en_SH'
func (t *en_SH) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
