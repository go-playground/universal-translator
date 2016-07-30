package zh_Hans

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type zh_Hans struct {
	locale string
}

// New returns a new instance of translator for the 'zh_Hans' locale
func New() locales.Translator {
	return &zh_Hans{
		locale: "zh_Hans",
	}
}

// Locale returns the current translators string locale
func (l *zh_Hans) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *zh_Hans) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
