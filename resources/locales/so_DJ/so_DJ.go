package so_DJ

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type so_DJ struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'so_DJ' locale
func New() locales.Translator {
	return &so_DJ{
		locale:  "so_DJ",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *so_DJ) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'so_DJ'
func (t *so_DJ) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'so_DJ'
func (t *so_DJ) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
