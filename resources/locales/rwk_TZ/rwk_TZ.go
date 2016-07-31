package rwk_TZ

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type rwk_TZ struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'rwk_TZ' locale
func New() locales.Translator {
	return &rwk_TZ{
		locale:  "rwk_TZ",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *rwk_TZ) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'rwk_TZ'
func (t *rwk_TZ) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'rwk_TZ'
func (t *rwk_TZ) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
