package tzm

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type tzm struct {
	locale string
}

// New returns a new instance of translator for the 'tzm' locale
func New() locales.Translator {
	return &tzm{
		locale: "tzm",
	}
}

// Locale returns the current translators string locale
func (l *tzm) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *tzm) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if (n >= 0 && n <= 1) || (n >= 11 && n <= 99) {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
