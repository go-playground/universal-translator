package bn_IN

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type bn_IN struct {
	locale string
}

// New returns a new instance of translator for the 'bn_IN' locale
func New() locales.Translator {
	return &bn_IN{
		locale: "bn_IN",
	}
}

// Locale returns the current translators string locale
func (l *bn_IN) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *bn_IN) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
	}

	return locales.PluralRuleOther, nil
}
