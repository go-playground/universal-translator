package sv_SE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type sv_SE struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'sv_SE' locale
func New() locales.Translator {
	return &sv_SE{
		locale:  "sv_SE",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *sv_SE) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'sv_SE'
func (t *sv_SE) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'sv_SE'
func (t *sv_SE) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
