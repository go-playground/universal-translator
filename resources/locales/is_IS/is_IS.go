package is_IS

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type is_IS struct {
	locale string
}

// New returns a new instance of translator for the 'is_IS' locale
func New() locales.Translator {
	return &is_IS{
		locale: "is_IS",
	}
}

// Locale returns the current translators string locale
func (l *is_IS) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *is_IS) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	t, err := locales.T(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if (t == 0 && i%10 == 1 && i%100 != 11) || (t != 0) {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
