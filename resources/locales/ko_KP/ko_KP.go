package ko_KP

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ko_KP struct {
	locale string
}

// New returns a new instance of translator for the 'ko_KP' locale
func New() locales.Translator {
	return &ko_KP{
		locale: "ko_KP",
	}
}

// Locale returns the current translators string locale
func (l *ko_KP) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ko_KP) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
