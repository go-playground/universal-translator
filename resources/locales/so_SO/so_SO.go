package so_SO

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type so_SO struct {
	locale string
}

// New returns a new instance of translator for the 'so_SO' locale
func New() locales.Translator {
	return &so_SO{
		locale: "so_SO",
	}
}

// Locale returns the current translators string locale
func (l *so_SO) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *so_SO) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
