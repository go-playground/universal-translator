package zh_Hans_HK

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type zh_Hans_HK struct {
	locale string
}

// New returns a new instance of translator for the 'zh_Hans_HK' locale
func New() locales.Translator {
	return &zh_Hans_HK{
		locale: "zh_Hans_HK",
	}
}

// Locale returns the current translators string locale
func (l *zh_Hans_HK) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *zh_Hans_HK) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
