package hsb

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type hsb struct {
	locale string
}

// New returns a new instance of translator for the 'hsb' locale
func New() locales.Translator {
	return &hsb{
		locale: "hsb",
	}
}

// Locale returns the current translators string locale
func (l *hsb) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *hsb) CardinalPluralRule(num string) (locales.PluralRule, error) {

	f, err := locales.F(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	v := locales.V(num)

	if (v == 0 && i%100 == 1) || (f%100 == 1) {
		return locales.PluralRuleOne, nil
	} else if (v == 0 && i%100 == 2) || (f%100 == 2) {
		return locales.PluralRuleTwo, nil
	} else if (v == 0 && i%100 >= 3 && i%100 <= 4) || (f%100 >= 3 && f%100 <= 4) {
		return locales.PluralRuleFew, nil
	}

	return locales.PluralRuleOther, nil
}
