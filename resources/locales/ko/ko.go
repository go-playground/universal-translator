package ko

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ko struct {
	locale string
}

// New returns a new instance of translator for the 'ko' locale
func New() locales.Translator {
	return &ko{
		locale: "ko",
	}
}

// Locale returns the current translators string locale
func (l *ko) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ko) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
