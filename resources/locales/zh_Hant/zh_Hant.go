package zh_Hant

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type zh_Hant struct {
	locale string
}

// New returns a new instance of translator for the 'zh_Hant' locale
func New() locales.Translator {
	return &zh_Hant{
		locale: "zh_Hant",
	}
}

// Locale returns the current translators string locale
func (l *zh_Hant) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *zh_Hant) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
