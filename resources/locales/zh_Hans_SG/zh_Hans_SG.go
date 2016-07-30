package zh_Hans_SG

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type zh_Hans_SG struct {
	locale string
}

// New returns a new instance of translator for the 'zh_Hans_SG' locale
func New() locales.Translator {
	return &zh_Hans_SG{
		locale: "zh_Hans_SG",
	}
}

// Locale returns the current translators string locale
func (l *zh_Hans_SG) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *zh_Hans_SG) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
