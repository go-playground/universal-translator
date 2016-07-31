package bs_Cyrl_BA

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type bs_Cyrl_BA struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'bs_Cyrl_BA' locale
func New() locales.Translator {
	return &bs_Cyrl_BA{
		locale:  "bs_Cyrl_BA",
		plurals: []locales.PluralRule{2, 4, 6},
	}
}

// Locale returns the current translators string locale
func (t *bs_Cyrl_BA) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'bs_Cyrl_BA'
func (t *bs_Cyrl_BA) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'bs_Cyrl_BA'
func (t *bs_Cyrl_BA) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	v := locales.V(num)

	f, err := locales.F(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if (v == 0 && i%10 == 1 && i%100 != 11) || (f%10 == 1 && f%100 != 11) {
		return locales.PluralRuleOne, nil
	} else if (v == 0 && i%10 >= 2 && i%10 <= 4 && i%100 < 12 && i%100 > 14) || (f%10 >= 2 && f%10 <= 4 && f%100 < 12 && f%100 > 14) {
		return locales.PluralRuleFew, nil
	}

	return locales.PluralRuleOther, nil
}
