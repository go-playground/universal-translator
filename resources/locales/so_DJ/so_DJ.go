package so_DJ

import (
	"math"

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

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'so_DJ'
func (t *so_DJ) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
