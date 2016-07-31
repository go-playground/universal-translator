package af_ZA

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type af_ZA struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'af_ZA' locale
func New() locales.Translator {
	return &af_ZA{
		locale:  "af_ZA",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *af_ZA) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'af_ZA'
func (t *af_ZA) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'af_ZA'
func (t *af_ZA) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
