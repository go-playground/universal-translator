package ja_JP

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ja_JP struct {
	locale string
}

// New returns a new instance of translator for the 'ja_JP' locale
func New() locales.Translator {
	return &ja_JP{
		locale: "ja_JP",
	}
}

// Locale returns the current translators string locale
func (l *ja_JP) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ja_JP) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
