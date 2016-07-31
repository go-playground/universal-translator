package ak_GH

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ak_GH struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ak_GH' locale
func New() locales.Translator {
	return &ak_GH{
		locale:  "ak_GH",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ak_GH) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ak_GH'
func (t *ak_GH) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ak_GH'
func (t *ak_GH) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n >= 0 && n <= 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
