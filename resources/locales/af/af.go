package af

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type af struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'af' locale
func New() locales.Translator {
	return &af{
		locale:  "af",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *af) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'af'
func (t *af) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'af'
func (t *af) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
