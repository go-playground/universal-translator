package zh_Hans_MO

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type zh_Hans_MO struct {
	locale string
}

// New returns a new instance of translator for the 'zh_Hans_MO' locale
func New() locales.Translator {
	return &zh_Hans_MO{
		locale: "zh_Hans_MO",
	}
}

// Locale returns the current translators string locale
func (l *zh_Hans_MO) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *zh_Hans_MO) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
