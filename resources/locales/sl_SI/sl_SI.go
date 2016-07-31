package sl_SI

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type sl_SI struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'sl_SI' locale
func New() locales.Translator {
	return &sl_SI{
		locale:  "sl_SI",
		plurals: []locales.PluralRule{2, 3, 4, 6},
	}
}

// Locale returns the current translators string locale
func (t *sl_SI) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'sl_SI'
func (t *sl_SI) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'sl_SI'
func (t *sl_SI) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	v := locales.V(num)

	if v == 0 && i%100 == 1 {
		return locales.PluralRuleOne, nil
	} else if v == 0 && i%100 == 2 {
		return locales.PluralRuleTwo, nil
	} else if (v == 0 && i%100 >= 3 && i%100 <= 4) || (v != 0) {
		return locales.PluralRuleFew, nil
	}

	return locales.PluralRuleOther, nil
}
