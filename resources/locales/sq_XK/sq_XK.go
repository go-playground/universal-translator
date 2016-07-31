package sq_XK

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type sq_XK struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'sq_XK' locale
func New() locales.Translator {
	return &sq_XK{
		locale:  "sq_XK",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *sq_XK) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'sq_XK'
func (t *sq_XK) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'sq_XK'
func (t *sq_XK) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
