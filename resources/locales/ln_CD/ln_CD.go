package ln_CD

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ln_CD struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ln_CD' locale
func New() locales.Translator {
	return &ln_CD{
		locale:  "ln_CD",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ln_CD) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ln_CD'
func (t *ln_CD) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ln_CD'
func (t *ln_CD) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n >= 0 && n <= 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
