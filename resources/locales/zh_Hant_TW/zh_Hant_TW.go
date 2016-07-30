package zh_Hant_TW

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type zh_Hant_TW struct {
	locale string
}

// New returns a new instance of translator for the 'zh_Hant_TW' locale
func New() locales.Translator {
	return &zh_Hant_TW{
		locale: "zh_Hant_TW",
	}
}

// Locale returns the current translators string locale
func (l *zh_Hant_TW) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *zh_Hant_TW) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
