package shi_Tfng

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type shi_Tfng struct {
	locale string
}

// New returns a new instance of translator for the 'shi_Tfng' locale
func New() locales.Translator {
	return &shi_Tfng{
		locale: "shi_Tfng",
	}
}

// Locale returns the current translators string locale
func (l *shi_Tfng) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *shi_Tfng) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne, nil
	} else if n >= 2 && n <= 10 {
		return locales.PluralRuleFew, nil
	}

	return locales.PluralRuleOther, nil
}
