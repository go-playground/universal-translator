package rwk_TZ

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type rwk_TZ struct {
	locale string
}

// New returns a new instance of translator for the 'rwk_TZ' locale
func New() locales.Translator {
	return &rwk_TZ{
		locale: "rwk_TZ",
	}
}

// Locale returns the current translators string locale
func (l *rwk_TZ) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *rwk_TZ) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
