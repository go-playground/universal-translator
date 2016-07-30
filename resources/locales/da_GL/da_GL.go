package da_GL

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type da_GL struct {
	locale string
}

// New returns a new instance of translator for the 'da_GL' locale
func New() locales.Translator {
	return &da_GL{
		locale: "da_GL",
	}
}

// Locale returns the current translators string locale
func (l *da_GL) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *da_GL) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	t, err := locales.T(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if (n == 1) || (t != 0 && (i == 0 || i == 1)) {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
