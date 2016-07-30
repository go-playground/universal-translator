package zh

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type zh struct {
	locale string
}

// New returns a new instance of translator for the 'zh' locale
func New() locales.Translator {
	return &zh{
		locale: "zh",
	}
}

// Locale returns the current translators string locale
func (l *zh) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *zh) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
