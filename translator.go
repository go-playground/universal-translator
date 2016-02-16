package ut

import (
	"errors"
	"fmt"
	"log"
	"time"
)

type translation struct {
	text string
}

// map[key]map[plural type other, many, few, single]*translation
type translations map[PluralRule]map[string]*translation
type groups map[string][]*translation

// Translator holds the locale translation instance
type Translator struct {
	locale       *Locale
	ruler        PluralRuler
	translations translations
	groups       groups
}

func newTranslator(loc *Locale) *Translator {

	return &Translator{
		locale:       loc,
		ruler:        pluralRules[loc.PluralRule],
		translations: make(translations),
		groups:       make(groups),
	}
}

// Add registers a new translation to the Translator using the
// key for a unique translation and group allows for retrieving
// translations for dev purposes later i.e. load all translations
// for the "homepage" group
// plural is optional, can be blank, but allows adding a plural translation
// at the same time as a singular
func (t *Translator) Add(rule PluralRule, group string, key string, text string) {

	trans := &translation{
		text: text,
	}

	if _, ok := t.translations[rule]; !ok {
		t.translations[rule] = make(map[string]*translation)
	}

	if _, ok := t.translations[rule][key]; ok {
		panic(fmt.Sprintf("Translation with key '%s' already exists", key))
	}

	t.translations[rule][key] = trans

	if _, ok := t.groups[group]; !ok {
		t.groups[group] = make([]*translation, 0)
	}

	t.groups[group] = append(t.groups[group], trans)
}

// PrintPluralRules prints the supported rules for the
// given translators locale
func (t *Translator) PrintPluralRules() {
	rules := pluralPluralRules[t.locale.PluralRule]

	fmt.Println("Translator locale '" + t.locale.Locale + "' supported rules:")

	for _, rule := range rules {
		fmt.Println("- " + rule.String())
	}

	fmt.Println("")
}

func (t *Translator) pluralRule(count NumberValue) (rule PluralRule) {
	return t.ruler.FindRule(count)
}

// T translates the text associated with the given key with the
// arguments passed in
func (t *Translator) T(key string, a ...interface{}) string {
	return t.P(key, 1, a...)
}

// TSafe translates the text associated with the given key with the
// arguments passed in, if the key or rule cannot be found it returns an error
func (t *Translator) TSafe(key string, a ...interface{}) (string, error) {
	return t.PSafe(key, 1, a...)
}

// P translates the plural text associated with the given key with the
// arguments passed in
func (t *Translator) P(key string, count interface{}, a ...interface{}) string {

	trans, err := t.PSafe(key, count, a...)
	if err != nil {
		log.Println(err.Error())
		return err.Error()
	}

	return trans
}

// PSafe translates the plural text associated with the given key with the
// arguments passed in, if the key or rule cannot be found it returns an error
func (t *Translator) PSafe(key string, count interface{}, a ...interface{}) (string, error) {

	rule := t.ruler.FindRule(count)

	trans, ok := t.translations[rule][key]
	if !ok {
		return "", errors.New("***** WARNING:***** Translation Key " + key + " Not Found")
	}

	return fmt.Sprintf(trans.text, a...), nil
}

// FmtDateFullSafe formats the time with the current locales full date format
func (t *Translator) FmtDateFullSafe(dt time.Time) (string, error) {
	return t.locale.Calendar.FmtDateFullSafe(dt)
}

// FmtDateLongSafe formats the time with the current locales long date format
func (t *Translator) FmtDateLongSafe(dt time.Time) (string, error) {
	return t.locale.Calendar.FmtDateLongSafe(dt)
}

// FmtDateMediumSafe formats the time with the current locales medium date format
func (t *Translator) FmtDateMediumSafe(dt time.Time) (string, error) {
	return t.locale.Calendar.FmtDateMediumSafe(dt)
}

// FmtDateShortSafe formats the time with the current locales short date format
func (t *Translator) FmtDateShortSafe(dt time.Time) (string, error) {
	return t.locale.Calendar.FmtDateShortSafe(dt)
}

// FmtDateTimeFullSafe formats the time with the current locales full data & time format
func (t *Translator) FmtDateTimeFullSafe(dt time.Time) (string, error) {
	return t.locale.Calendar.FmtDateTimeFullSafe(dt)
}

// FmtDateTimeLongSafe formats the time with the current locales long data & time format
func (t *Translator) FmtDateTimeLongSafe(dt time.Time) (string, error) {
	return t.locale.Calendar.FmtDateTimeLongSafe(dt)
}

// FmtDateTimeMediumSafe formats the time with the current locales medium data & time format
func (t *Translator) FmtDateTimeMediumSafe(dt time.Time) (string, error) {
	return t.locale.Calendar.FmtDateTimeMediumSafe(dt)
}

// FmtDateTimeShortSafe formats the time with the current locales short data & time format
func (t *Translator) FmtDateTimeShortSafe(dt time.Time) (string, error) {
	return t.locale.Calendar.FmtDateTimeShortSafe(dt)
}

// FmtTimeFullSafe formats the time with the current locales full time format
func (t *Translator) FmtTimeFullSafe(dt time.Time) (string, error) {
	return t.locale.Calendar.FmtTimeFullSafe(dt)
}

// FmtTimeLongSafe formats the time with the current locales long time format
func (t *Translator) FmtTimeLongSafe(dt time.Time) (string, error) {
	return t.locale.Calendar.FmtTimeLongSafe(dt)
}

// FmtTimeMediumSafe formats the time with the current locales medium time format
func (t *Translator) FmtTimeMediumSafe(dt time.Time) (string, error) {
	return t.locale.Calendar.FmtTimeMediumSafe(dt)
}

// FmtTimeShortSafe formats the time with the current locales short time format
func (t *Translator) FmtTimeShortSafe(dt time.Time) (string, error) {
	return t.locale.Calendar.FmtTimeShortSafe(dt)
}

// FmtDateFull formats the time with the current locales full date format
func (t *Translator) FmtDateFull(dt time.Time) string {
	return t.locale.Calendar.FmtDateFull(dt)
}

// FmtDateLong formats the time with the current locales long date format
func (t *Translator) FmtDateLong(dt time.Time) string {
	return t.locale.Calendar.FmtDateLong(dt)
}

// FmtDateMedium formats the time with the current locales medium date format
func (t *Translator) FmtDateMedium(dt time.Time) string {
	return t.locale.Calendar.FmtDateMedium(dt)
}

// FmtDateShort formats the time with the current locales short date format
func (t *Translator) FmtDateShort(dt time.Time) string {
	return t.locale.Calendar.FmtDateShort(dt)
}

// FmtDateTimeFull formats the time with the current locales full data & time format
func (t *Translator) FmtDateTimeFull(dt time.Time) string {
	return t.locale.Calendar.FmtDateTimeFull(dt)
}

// FmtDateTimeLong formats the time with the current locales long data & time format
func (t *Translator) FmtDateTimeLong(dt time.Time) string {
	return t.locale.Calendar.FmtDateTimeLong(dt)
}

// FmtDateTimeMedium formats the time with the current locales medium data & time format
func (t *Translator) FmtDateTimeMedium(dt time.Time) string {
	return t.locale.Calendar.FmtDateTimeMedium(dt)
}

// FmtDateTimeShort formats the time with the current locales short data & time format
func (t *Translator) FmtDateTimeShort(dt time.Time) string {
	return t.locale.Calendar.FmtDateTimeShort(dt)
}

// FmtTimeFull formats the time with the current locales full time format
func (t *Translator) FmtTimeFull(dt time.Time) string {
	return t.locale.Calendar.FmtTimeFull(dt)
}

// FmtTimeLong formats the time with the current locales long time format
func (t *Translator) FmtTimeLong(dt time.Time) string {
	return t.locale.Calendar.FmtTimeLong(dt)
}

// FmtTimeMedium formats the time with the current locales medium time format
func (t *Translator) FmtTimeMedium(dt time.Time) string {
	return t.locale.Calendar.FmtTimeMedium(dt)
}

// FmtTimeShort formats the time with the current locales short time format
func (t *Translator) FmtTimeShort(dt time.Time) string {
	return t.locale.Calendar.FmtTimeShort(dt)
}

// FmtDateTimeSafe formats the time with the current locales short time format
func (t *Translator) FmtDateTimeSafe(dt time.Time, pattern string) (string, error) {
	return t.locale.Calendar.FormatSafe(dt, pattern)
}

// FmtDateTime formats the time with the current locales short time format
func (t *Translator) FmtDateTime(dt time.Time, pattern string) string {
	return t.locale.Calendar.Format(dt, pattern)
}

// FmtCurrencySafe takes a float number and a currency key and returns a string
// with a properly formatted currency amount with the correct currency symbol.
// If a symbol cannot be found for the reqested currency, the the key is used
// instead. If the currency key requested is not recognized, it is used as the
// symbol, and an error is returned with the formatted string.
func (t *Translator) FmtCurrencySafe(typ CurrencyType, currency string, number interface{}) (string, error) {
	return t.locale.Number.FmtCurrencySafe(typ, currency, toFloat64(number))
}

// FmtCurrencyWholeSafe does exactly what FormatCurrency does, but it leaves off
// any decimal places. AKA, it would return $100 rather than $100.00.
func (t *Translator) FmtCurrencyWholeSafe(typ CurrencyType, currency string, number interface{}) (string, error) {
	return t.locale.Number.FmtCurrencyWholeSafe(typ, currency, toFloat64(number))
}

// FmtCurrency takes a float number and a currency key and returns a string
// with a properly formatted currency amount with the correct currency symbol.
// If a symbol cannot be found for the reqested currency, this will panic, use
// FmtCurrencySafe for non panicing variant.
func (t *Translator) FmtCurrency(typ CurrencyType, currency string, number interface{}) string {
	return t.locale.Number.FmtCurrency(typ, currency, toFloat64(number))
}

// FmtCurrencyWhole does exactly what FormatCurrency does, but it leaves off
// any decimal places. AKA, it would return $100 rather than $100.00.
// If a symbol cannot be found for the reqested currency, this will panic, use
// FmtCurrencyWholeSafe for non panicing variant.
func (t *Translator) FmtCurrencyWhole(typ CurrencyType, currency string, number interface{}) string {
	return t.locale.Number.FmtCurrencyWhole(typ, currency, toFloat64(number))
}

// FmtNumber takes a float number and returns a properly formatted string
// representation of that number according to the locale's number format.
func (t *Translator) FmtNumber(number interface{}) string {
	return t.locale.Number.FmtNumber(toFloat64(number))
}

// FmtNumberWhole does exactly what FormatNumber does, but it leaves off any
// decimal places. AKA, it would return 100 rather than 100.01.
func (t *Translator) FmtNumberWhole(number interface{}) string {
	return t.locale.Number.FmtNumberWhole(toFloat64(number))
}

// FmtPercent takes a float number and returns a properly formatted string
// representation of that number as a percentage according to the locale's
// percentage format.
func (t *Translator) FmtPercent(number interface{}) string {
	return t.locale.Number.FmtPercent(toFloat64(number))
}
