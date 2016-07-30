package pa

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type pa struct {
	locale string
}

// New returns a new instance of translator for the 'pa' locale
func New() locales.Translator {
	return &pa{
		locale: "pa",
	}
}

// Locale returns the current translators string locale
func (l *pa) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *pa) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n >= 0 && n <= 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
