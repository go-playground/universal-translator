package hsb_DE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type hsb_DE struct {
	locale string
}

// New returns a new instance of translator for the 'hsb_DE' locale
func New() locales.Translator {
	return &hsb_DE{
		locale: "hsb_DE",
	}
}

// Locale returns the current translators string locale
func (l *hsb_DE) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *hsb_DE) CardinalPluralRule(num string) (locales.PluralRule, error) {

	v := locales.V(num)

	f, err := locales.F(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if (v == 0 && i%100 == 1) || (f%100 == 1) {
		return locales.PluralRuleOne, nil
	} else if (v == 0 && i%100 == 2) || (f%100 == 2) {
		return locales.PluralRuleTwo, nil
	} else if (v == 0 && i%100 >= 3 && i%100 <= 4) || (f%100 >= 3 && f%100 <= 4) {
		return locales.PluralRuleFew, nil
	}

	return locales.PluralRuleOther, nil
}
