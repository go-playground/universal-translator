package ko_KR

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ko_KR struct {
	locale string
}

// New returns a new instance of translator for the 'ko_KR' locale
func New() locales.Translator {
	return &ko_KR{
		locale: "ko_KR",
	}
}

// Locale returns the current translators string locale
func (l *ko_KR) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ko_KR) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
