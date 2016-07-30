package fur_IT

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type fur_IT struct {
	locale string
}

// New returns a new instance of translator for the 'fur_IT' locale
func New() locales.Translator {
	return &fur_IT{
		locale: "fur_IT",
	}
}

// Locale returns the current translators string locale
func (l *fur_IT) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *fur_IT) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
