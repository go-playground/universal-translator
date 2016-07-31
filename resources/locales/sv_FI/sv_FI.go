package sv_FI

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type sv_FI struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'sv_FI' locale
func New() locales.Translator {
	return &sv_FI{
		locale:  "sv_FI",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *sv_FI) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'sv_FI'
func (t *sv_FI) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'sv_FI'
func (t *sv_FI) CardinalPluralRule(num string) (locales.PluralRule, error) {

	v := locales.V(num)

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if i == 1 && v == 0 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
