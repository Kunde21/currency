// Copyright (c) 2020 Bojan Zivanovic and contributors
// SPDX-License-Identifier: MIT

//go:build ignore
// +build ignore

package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/bojanz/currency"
)

const assetDir = "raw"

const dataTemplate = `// Code generated by go generate; DO NOT EDIT.
//go:generate go run gen.go

package currency

// CLDRVersion is the CLDR version from which the data is derived.
const CLDRVersion = "{{ .CLDRVersion }}"

type numberingSystem uint8

const (
	numLatn numberingSystem = iota
	numArab
	numArabExt
	numBeng
	numDeva
	numMymr
)

type currencyInfo struct {
	numericCode string
	digits      uint8
}

type symbolInfo struct {
	symbol  string
	locales []string
}

type currencyFormat struct {
	standardPattern       string
	accountingPattern     string
	numberingSystem       numberingSystem
	minGroupingDigits     uint8
	primaryGroupingSize   uint8
	secondaryGroupingSize uint8
	decimalSeparator      string
	groupingSeparator     string
	plusSign              string
	minusSign             string
}

// Defined separately to ensure consistent ordering (G10, then others).
var currencyCodes = []string{
	// G10 currencies https://en.wikipedia.org/wiki/G10_currencies.
	{{ export .G10Currencies 10 "\t" }}

	// Other currencies.
	{{ export .OtherCurrencies 10 "\t" }}
}

var currencies = map[string]currencyInfo{
	{{ export .CurrencyInfo 3 "\t" }}
}

var currencySymbols = map[string][]symbolInfo{
	{{ export .SymbolInfo 1 "\t" }}
}

var currencyFormats = map[string]currencyFormat{
	{{ export .Formats 1 "\t" }}
}

var countryCurrencies = map[string]string{
	{{ export .CountryCurrencies 5 "\t" }}
}

var parentLocales = map[string]string{
	{{ export .ParentLocales 3 "\t" }}
}
`

type currencyInfo struct {
	numericCode string
	digits      uint8
}

func (c currencyInfo) GoString() string {
	return fmt.Sprintf("{%q, %d}", c.numericCode, int(c.digits))
}

type symbolInfo struct {
	symbol  string
	locales []string
}

func (s symbolInfo) GoString() string {
	return fmt.Sprintf("{%q, %#v}", s.symbol, s.locales)
}

type symbolInfoSlice []*symbolInfo

func (ss symbolInfoSlice) GoString() string {
	b := strings.Builder{}
	b.WriteString("{\n")
	for _, s := range ss {
		b.WriteString("\t\t")
		fmt.Fprintf(&b, "%#v,\n", s)
	}
	b.WriteString("\t}")

	return b.String()
}

type numberingSystem uint8

const (
	numLatn numberingSystem = iota
	numArab
	numArabExt
	numBeng
	numDeva
	numMymr
)

type currencyFormat struct {
	standardPattern       string
	accountingPattern     string
	numberingSystem       numberingSystem
	minGroupingDigits     uint8
	primaryGroupingSize   uint8
	secondaryGroupingSize uint8
	decimalSeparator      string
	groupingSeparator     string
	plusSign              string
	minusSign             string
}

func (f currencyFormat) GoString() string {
	return fmt.Sprintf("{%q, %q, %d, %d, %d, %d, %q, %q, %q, %q}", f.standardPattern, f.accountingPattern, f.numberingSystem, f.minGroupingDigits, f.primaryGroupingSize, f.secondaryGroupingSize, f.decimalSeparator, f.groupingSeparator, f.plusSign, f.minusSign)
}

func main() {
	err := os.Mkdir(assetDir, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(assetDir)

	log.Println("Fetching CLDR data...")
	CLDRVersion, err := fetchCLDR(assetDir)
	if err != nil {
		os.RemoveAll(assetDir)
		log.Fatal(err)
	}

	log.Println("Fetching ISO data...")
	currencies, err := fetchISO()
	if err != nil {
		os.RemoveAll(assetDir)
		log.Fatal(err)
	}

	log.Println("Processing...")
	err = replaceDigits(currencies, assetDir)
	if err != nil {
		os.RemoveAll(assetDir)
		log.Fatal(err)
	}
	locales, err := collectLocales(assetDir)
	if err != nil {
		os.RemoveAll(assetDir)
		log.Fatal(err)
	}
	symbols, err := generateSymbols(currencies, locales, assetDir)
	if err != nil {
		os.RemoveAll(assetDir)
		log.Fatal(err)
	}
	formats, err := generateFormats(locales, assetDir)
	if err != nil {
		os.RemoveAll(assetDir)
		log.Fatal(err)
	}
	countryCurrencies, err := generateCountryCurrencies(assetDir)
	if err != nil {
		os.RemoveAll(assetDir)
		log.Fatal(err)
	}
	parentLocales, err := generateParentLocales(locales, assetDir)
	if err != nil {
		os.RemoveAll(assetDir)
		log.Fatal(err)
	}

	var currencyCodes []string
	for currencyCode := range currencies {
		currencyCodes = append(currencyCodes, currencyCode)
	}
	sort.Strings(currencyCodes)

	g10Currencies := []string{
		"AUD", "CAD", "CHF", "EUR", "GBP", "JPY", "NOK", "NZD", "SEK", "USD",
	}
	var otherCurrencies []string
	for _, currencyCode := range currencyCodes {
		if !contains(g10Currencies, currencyCode) {
			otherCurrencies = append(otherCurrencies, currencyCode)
		}
	}

	os.Remove("data.go")
	f, err := os.Create("data.go")
	if err != nil {
		os.RemoveAll(assetDir)
		log.Fatal(err)
	}
	defer f.Close()

	funcMap := template.FuncMap{
		"export": export,
	}
	t, err := template.New("data").Funcs(funcMap).Parse(dataTemplate)
	if err != nil {
		os.RemoveAll(assetDir)
		log.Fatal(err)
	}
	t.Execute(f, struct {
		CLDRVersion       string
		G10Currencies     []string
		OtherCurrencies   []string
		CurrencyInfo      map[string]*currencyInfo
		SymbolInfo        map[string]symbolInfoSlice
		Formats           map[string]currencyFormat
		CountryCurrencies map[string]string
		ParentLocales     map[string]string
	}{
		CLDRVersion:       CLDRVersion,
		G10Currencies:     g10Currencies,
		OtherCurrencies:   otherCurrencies,
		CurrencyInfo:      currencies,
		SymbolInfo:        symbols,
		Formats:           formats,
		CountryCurrencies: countryCurrencies,
		ParentLocales:     parentLocales,
	})

	log.Println("Done.")
}

// fetchCLDR fetches the CLDR data from GitHub and returns its version.
//
// The JSON version of the data is used because it is more convenient
// to parse. See https://github.com/unicode-org/cldr-json for details.
func fetchCLDR(dir string) (string, error) {
	repo := "https://github.com/unicode-org/cldr-json.git"
	cmd := exec.Command("git", "clone", repo, "--depth", "1", dir)
	cmd.Stderr = os.Stderr
	_, err := cmd.Output()
	if err != nil {
		return "", err
	}

	data, err := os.ReadFile(dir + "/cldr-json/cldr-core/package.json")
	if err != nil {
		return "", fmt.Errorf("fetchCLDR: %w", err)
	}
	aux := struct {
		Version string
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return "", fmt.Errorf("fetchCLDR: %w", err)
	}

	return aux.Version, nil
}

// fetchISO fetches currency info from ISO.
//
// ISO data is needed because CLDR can't be used as a reliable source
// of numeric codes (e.g. BYR has no numeric code as of CLDR v36).
// Furthermore, CLDR includes both active and inactive currencies, while ISO
// includes only active ones, matching the needs of this package.
func fetchISO() (map[string]*currencyInfo, error) {
	data, err := fetchURL("https://www.six-group.com/dam/download/financial-information/data-center/iso-currrency/lists/list-one.xml")
	if err != nil {
		return nil, fmt.Errorf("fetchISO: %w", err)
	}
	aux := struct {
		Table []struct {
			Entry []struct {
				Code    string `xml:"Ccy"`
				Number  string `xml:"CcyNbr"`
				Digits  string `xml:"CcyMnrUnts"`
				Country string `xml:"CtryNm"`
				Name    struct {
					Value  string `xml:",chardata"`
					IsFund bool   `xml:"IsFund,attr"`
				} `xml:"CcyNm"`
			} `xml:"CcyNtry"`
		} `xml:"CcyTbl"`
	}{}
	if err := xml.Unmarshal(data, &aux); err != nil {
		return nil, fmt.Errorf("fetchISO: %w", err)
	}

	currencies := make(map[string]*currencyInfo, 170)
	for _, entry := range aux.Table[0].Entry {
		if entry.Code == "" || entry.Number == "" || entry.Digits == "N.A." {
			continue
		}

		// We use ISO digits here with a fallback to 2, but prefer CLDR
		// data when available. See replaceDigits() for the next step.
		digits := parseDigits(entry.Digits, 2)
		currencies[entry.Code] = &currencyInfo{entry.Number, digits}
	}

	return currencies, nil
}

func fetchURL(url string) ([]byte, error) {
	client := http.Client{Timeout: 15 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetchURL: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetchURL: Get %q: %v", url, resp.Status)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("fetchURL: Get %q: %w", url, err)
	}

	return data, nil
}

// collectLocales collects CLDR locales with "modern" coverage.
func collectLocales(dir string) ([]string, error) {
	data, err := os.ReadFile(dir + "/cldr-json/cldr-core/coverageLevels.json")
	if err != nil {
		return nil, fmt.Errorf("collectLocales: %w", err)
	}
	aux := struct {
		EffectiveCoverageLevels map[string]string
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return nil, fmt.Errorf("collectLocales: %w", err)
	}

	locales := make([]string, 0, len(aux.EffectiveCoverageLevels))
	for locale, level := range aux.EffectiveCoverageLevels {
		if !shouldIgnoreLocale(locale) && level == "modern" {
			locales = append(locales, locale)
		}
	}

	return locales, nil
}

// replaceDigits replaces currency digits with data from CLDR.
//
// CLDR data reflects real life usage more closely, specifying 0 digits
// (instead of 2 in ISO data) for ~14 currencies, such as ALL and RSD.
//
// Note that CLDR does not have data for every currency, in which ase
// the original ISO digits are kept.
func replaceDigits(currencies map[string]*currencyInfo, dir string) error {
	data, err := os.ReadFile(dir + "/cldr-json/cldr-core/supplemental/currencyData.json")
	if err != nil {
		return fmt.Errorf("replaceDigits: %w", err)
	}
	aux := struct {
		Supplemental struct {
			CurrencyData struct {
				Fractions map[string]map[string]string
			}
		}
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return fmt.Errorf("replaceDigits: %w", err)
	}

	for currencyCode := range currencies {
		fractions, ok := aux.Supplemental.CurrencyData.Fractions[currencyCode]
		if ok {
			currencies[currencyCode].digits = parseDigits(fractions["_digits"], 2)
		}
	}

	return nil
}

// generateCountryCurrencies generates the map of country codes to currency codes.
func generateCountryCurrencies(dir string) (map[string]string, error) {
	data, err := os.ReadFile(dir + "/cldr-json/cldr-core/supplemental/currencyData.json")
	if err != nil {
		return nil, fmt.Errorf("generateCountryCurrencies: %w", err)
	}

	aux := struct {
		Supplemental struct {
			CurrencyData struct {
				Region map[string][]map[string]struct {
					From   string `json:"_from"`
					To     string `json:"_to"`
					Tender string `json:"_tender"`
				}
			}
		}
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return nil, fmt.Errorf("generateCountryCurrencies: %w", err)
	}

	countryCurrencies := make(map[string]string)
	for countryCode, currencies := range aux.Supplemental.CurrencyData.Region {
		if contains([]string{"EA", "EU", "ZZ"}, countryCode) {
			// EA, EU and ZZ are not countries.
			continue
		}

		lastCurrencyCode := ""
		lastFrom := ""
		for _, currencyUsage := range currencies {
			for currencyCode, usageInfo := range currencyUsage {
				if usageInfo.To != "" || usageInfo.Tender == "false" {
					// Currency no longer in use, skip.
					continue
				}
				if lastFrom == "" || usageInfo.From > lastFrom {
					lastCurrencyCode = currencyCode
					lastFrom = usageInfo.From
				}
			}
		}

		if lastCurrencyCode != "" && lastCurrencyCode != "XXX" {
			countryCurrencies[countryCode] = lastCurrencyCode
		}
	}

	return countryCurrencies, nil
}

// generateSymbols generates currency symbols for all locales.
//
// Symbols are grouped by locale, and deduplicated by parent.
func generateSymbols(currencies map[string]*currencyInfo, locales []string, dir string) (map[string]symbolInfoSlice, error) {
	symbols := make(map[string]map[string][]string)
	for _, locale := range locales {
		localSymbols, err := readSymbols(currencies, dir, locale)
		if err != nil {
			return nil, fmt.Errorf("generateSymbols: %w", err)
		}

		for currencyCode, symbol := range localSymbols {
			if _, ok := symbols[currencyCode]; !ok {
				symbols[currencyCode] = make(map[string][]string)
			}
			symbols[currencyCode][symbol] = append(symbols[currencyCode][symbol], locale)
		}
	}

	for currencyCode, localSymbols := range symbols {
		if _, ok := localSymbols[currencyCode]; !ok {
			continue
		}
		// currency.GetSymbol always returns the currency code when no
		// symbol is available, so there is no need to store a currency
		// whose only symbol is the currency code.
		if len(localSymbols) == 1 {
			delete(symbols, currencyCode)
			continue
		}
		// Diverge from CLDR by preventing child locales from using the
		// currency code as a symbol when the "en" locale has a real one.
		// E.g. the "hr", "hu", "ro" locales should use $, £, € instead
		// of USD, GBP, EUR, just like the "en" locale does.
		// This noticeably decreases the size of generated data.
		for symbol, locales := range localSymbols {
			if contains(locales, "en") && currencyCode != symbol {
				symbols[currencyCode][symbol] = append(symbols[currencyCode][symbol], localSymbols[currencyCode]...)
				delete(symbols[currencyCode], currencyCode)
				break
			}
		}
		// The logic above results in "en-AU" using the same $ symbol for AUD and USD.
		// Related: https://unicode-org.atlassian.net/projects/CLDR/issues/CLDR-10710
		if currencyCode == "USD" {
			// Move en-AU from symbols["USD"]["$"] to symbols["USD"]["US$"].
			li := 0
			for i, locale := range symbols["USD"]["$"] {
				if locale == "en-AU" {
					li = i
					break
				}
			}
			symbols["USD"]["$"] = append(symbols["USD"]["$"][:li], symbols["USD"]["$"][li+1:]...)
			symbols["USD"]["US$"] = append(symbols["USD"]["US$"], "en-AU")
		}
	}

	// Child locales don't need to be listed if the parent is present.
	for currencyCode, localSymbols := range symbols {
		for symbol, locales := range localSymbols {
			var deleteLocales []string
			for _, localeID := range locales {
				locale := currency.NewLocale(localeID)
				parent := locale.GetParent()
				if contains(locales, parent.String()) {
					deleteLocales = append(deleteLocales, localeID)
				}
			}
			symbols[currencyCode][symbol] = []string{}
			for _, localeID := range locales {
				if !contains(deleteLocales, localeID) {
					symbols[currencyCode][symbol] = append(symbols[currencyCode][symbol], localeID)
				}
			}
			sort.Strings(symbols[currencyCode][symbol])
		}
	}

	// Convert to the final data structure.
	currencySymbols := make(map[string]symbolInfoSlice)
	for currencyCode, localSymbols := range symbols {
		// Always put the "en" symbol first, then the other sorted symbols.
		for symbol, locales := range localSymbols {
			if contains(locales, "en") {
				currencySymbols[currencyCode] = append(currencySymbols[currencyCode], &symbolInfo{symbol, locales})
				break
			}
		}
		var symbolKeys []string
		for symbol := range localSymbols {
			symbolKeys = append(symbolKeys, symbol)
		}
		sort.Strings(symbolKeys)
		for _, symbol := range symbolKeys {
			locales := symbols[currencyCode][symbol]
			if !contains(locales, "en") {
				currencySymbols[currencyCode] = append(currencySymbols[currencyCode], &symbolInfo{symbol, locales})
			}
		}
	}

	return currencySymbols, nil
}

// readSymbols reads the given locale's currency symbols from CLDR data.
//
// Discards symbols belonging to inactive currencies.
func readSymbols(currencies map[string]*currencyInfo, dir string, locale string) (map[string]string, error) {
	filename := fmt.Sprintf("%v/cldr-json/cldr-numbers-full/main/%v/currencies.json", dir, locale)
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("readSymbols: %w", err)
	}

	type cldrData struct {
		Numbers struct {
			Currencies map[string]map[string]string
		}
	}
	aux := struct {
		Main map[string]cldrData
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return nil, fmt.Errorf("readSymbols: %w", err)
	}

	symbols := make(map[string]string)
	for currencyCode, data := range aux.Main[locale].Numbers.Currencies {
		if _, ok := currencies[currencyCode]; ok {
			symbols[currencyCode] = data["symbol"]
			// CLDR omits the symbol when it matches the currency code.
			if symbols[currencyCode] == "" {
				symbols[currencyCode] = currencyCode
			}
		}
	}

	return symbols, nil
}

// generateFormats generates currency formats from CLDR data.
//
// Formats are deduplicated by parent.
func generateFormats(locales []string, dir string) (map[string]currencyFormat, error) {
	formats := make(map[string]currencyFormat)
	for _, locale := range locales {
		format, err := readFormat(dir, locale)
		if err != nil {
			return nil, fmt.Errorf("generateFormats: %w", err)
		}
		formats[locale] = format
	}

	// Remove formats which are identical to their parents.
	var deleteLocales []string
	for localeID, format := range formats {
		locale := currency.NewLocale(localeID)
		parentID := locale.GetParent().String()
		if parentID != "" && format == formats[parentID] {
			deleteLocales = append(deleteLocales, localeID)
		}
	}
	for _, localeID := range deleteLocales {
		delete(formats, localeID)
	}

	return formats, nil
}

// readFormat reads the given locale's currency format from CLDR data.
func readFormat(dir string, locale string) (currencyFormat, error) {
	filename := fmt.Sprintf("%v/cldr-json/cldr-numbers-full/main/%v/numbers.json", dir, locale)
	data, err := os.ReadFile(filename)
	if err != nil {
		return currencyFormat{}, fmt.Errorf("readFormat: %w", err)
	}

	type cldrPattern struct {
		Standard   string
		Accounting string
	}
	type cldrData struct {
		Numbers struct {
			MinimumGroupingDigits  string
			DefaultNumberingSystem string
			PatternLatn            cldrPattern       `json:"currencyFormats-numberSystem-latn"`
			PatternArab            cldrPattern       `json:"currencyFormats-numberSystem-arab"`
			PatternArabExt         cldrPattern       `json:"currencyFormats-numberSystem-arabext"`
			PatternBeng            cldrPattern       `json:"currencyFormats-numberSystem-beng"`
			PatternDeva            cldrPattern       `json:"currencyFormats-numberSystem-deva"`
			PatternMymr            cldrPattern       `json:"currencyFormats-numberSystem-mymr"`
			SymbolsLatn            map[string]string `json:"symbols-numberSystem-latn"`
			SymbolsArab            map[string]string `json:"symbols-numberSystem-arab"`
			SymbolsArabExt         map[string]string `json:"symbols-numberSystem-arabext"`
			SymbolsBeng            map[string]string `json:"symbols-numberSystem-beng"`
			SymbolsDeva            map[string]string `json:"symbols-numberSystem-deva"`
			SymbolsMymr            map[string]string `json:"symbols-numberSystem-mymr"`
		}
	}
	aux := struct {
		Main map[string]cldrData
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return currencyFormat{}, fmt.Errorf("readFormat: %w", err)
	}

	var numSystem numberingSystem
	var standardPattern string
	var accountingPattern string
	var symbols map[string]string
	extFormat := aux.Main[locale].Numbers
	switch extFormat.DefaultNumberingSystem {
	case "latn":
		numSystem = numLatn
		standardPattern = extFormat.PatternLatn.Standard
		accountingPattern = extFormat.PatternLatn.Accounting
		symbols = extFormat.SymbolsLatn
	case "arab":
		numSystem = numArab
		standardPattern = extFormat.PatternArab.Standard
		accountingPattern = extFormat.PatternArab.Accounting
		symbols = extFormat.SymbolsArab
	case "arabext":
		numSystem = numArabExt
		standardPattern = extFormat.PatternArabExt.Standard
		accountingPattern = extFormat.PatternArabExt.Accounting
		symbols = extFormat.SymbolsArabExt
	case "beng":
		numSystem = numBeng
		standardPattern = extFormat.PatternBeng.Standard
		accountingPattern = extFormat.PatternBeng.Accounting
		symbols = extFormat.SymbolsBeng
	case "deva":
		numSystem = numDeva
		standardPattern = extFormat.PatternDeva.Standard
		accountingPattern = extFormat.PatternDeva.Accounting
		symbols = extFormat.SymbolsDeva
	case "mymr":
		numSystem = numMymr
		standardPattern = extFormat.PatternMymr.Standard
		accountingPattern = extFormat.PatternMymr.Accounting
		symbols = extFormat.SymbolsMymr
	default:
		return currencyFormat{}, fmt.Errorf("readFormat: unknown numbering system %q in locale %q", extFormat.DefaultNumberingSystem, locale)
	}
	primaryGroupingSize := 0
	secondaryGroupingSize := 0
	patternParts := strings.Split(standardPattern, ";")
	if strings.Contains(patternParts[0], ",") {
		r, _ := regexp.Compile("#+0")
		primaryGroup := r.FindString(patternParts[0])
		primaryGroupingSize = len(primaryGroup)
		secondaryGroupingSize = primaryGroupingSize
		numberGroups := strings.Split(patternParts[0], ",")
		if len(numberGroups) > 2 {
			// This pattern has a distinct secondary group size.
			secondaryGroupingSize = len(numberGroups[1])
		}
	}
	// Strip the grouping info from the patterns, now that it is available separately.
	standardPattern = processPattern(standardPattern)
	accountingPattern = processPattern(accountingPattern)
	// To save memory, the accounting pattern is stored
	// only if it's different from the standard pattern.
	if accountingPattern == standardPattern {
		accountingPattern = ""
	}
	decimalSeparator := symbols["decimal"]
	groupingSeparator := symbols["group"]
	// Most locales use the same separators for decimal and currency
	// formatting, with the exception of de-AT and fr-CH (as of CLDR v37).
	if _, ok := symbols["currencyDecimal"]; ok {
		decimalSeparator = symbols["currencyDecimal"]
	}
	if _, ok := symbols["currencyGroup"]; ok {
		groupingSeparator = symbols["currencyGroup"]
	}

	format := currencyFormat{}
	format.standardPattern = standardPattern
	format.accountingPattern = accountingPattern
	format.numberingSystem = numSystem
	format.minGroupingDigits = parseDigits(extFormat.MinimumGroupingDigits, 1)
	format.primaryGroupingSize = uint8(primaryGroupingSize)
	format.secondaryGroupingSize = uint8(secondaryGroupingSize)
	format.decimalSeparator = decimalSeparator
	format.groupingSeparator = groupingSeparator
	format.plusSign = symbols["plusSign"]
	format.minusSign = symbols["minusSign"]

	return format, nil
}

// processPattern processes the pattern.
func processPattern(pattern string) string {
	// Strip the grouping info.
	pattern = strings.ReplaceAll(pattern, "#,##,##", "")
	pattern = strings.ReplaceAll(pattern, "#,##", "")

	return pattern
}

// generateParentLocales generates parent locales from CLDR data.
//
// Ensures ignored locales are skipped.
// Replaces "und" with "en", since this package treats them as equivalent.
func generateParentLocales(locales []string, dir string) (map[string]string, error) {
	data, err := os.ReadFile(dir + "/cldr-json/cldr-core/supplemental/parentLocales.json")
	if err != nil {
		return nil, fmt.Errorf("generateParentLocales: %w", err)
	}
	aux := struct {
		Supplemental struct {
			ParentLocales struct {
				ParentLocale map[string]string
			}
		}
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return nil, fmt.Errorf("generateParentLocales: %w", err)
	}

	parentLocales := make(map[string]string)
	for locale, parent := range aux.Supplemental.ParentLocales.ParentLocale {
		// Avoid exposing the concept of "und" to users.
		if parent == "und" {
			parent = "en"
		}
		if slices.Contains(locales, locale) && !shouldIgnoreLocale(locale) {
			parentLocales[locale] = parent
		}
	}
	// Dsrt and Shaw are made up scripts.
	delete(parentLocales, "en-Dsrt")
	delete(parentLocales, "en-Shaw")

	return parentLocales, nil
}

func shouldIgnoreLocale(locale string) bool {
	ignoredLocales := []string{
		// English is our fallback, we don't need another.
		"und",
		// Esperanto, Interlingua, Volapuk are made up languages.
		"eo", "ia", "vo",
		// Belarus (Classical orthography), Church Slavic, Manx, Prussian are historical.
		"be-tarask", "cu", "gv", "prg",
		// Valencian differs from its parent only by a single character (è/é).
		"ca-ES-valencia",
	}
	localeParts := strings.Split(locale, "-")
	ignore := false
	for _, ignoredLocale := range ignoredLocales {
		if ignoredLocale == locale || ignoredLocale == localeParts[0] {
			ignore = true
			break
		}
	}

	return ignore
}

func contains(a []string, x string) bool {
	for _, v := range a {
		if v == x {
			return true
		}
	}
	return false
}

func parseDigits(n string, fallback uint8) uint8 {
	digits, err := strconv.Atoi(n)
	if err != nil {
		return fallback
	}
	return uint8(digits)
}

func export(i interface{}, width int, indent string) string {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Map:
		return exportMap(v, width, indent)
	case reflect.Slice:
		return exportSlice(v, width, indent)
	default:
		return fmt.Sprintf("%#v", i)
	}
}

func exportMap(v reflect.Value, width int, indent string) string {
	var maxKeyLen int
	var keys []string
	for _, k := range v.MapKeys() {
		key := k.Interface().(string)
		keys = append(keys, key)
		if len(key) > maxKeyLen {
			maxKeyLen = len(key)
		}
	}
	sort.Strings(keys)

	b := strings.Builder{}
	i := 0
	for _, key := range keys {
		spaceWidth := 1
		if width == 1 {
			spaceWidth += maxKeyLen - len(key)
		}
		space := strings.Repeat(" ", spaceWidth)
		value := v.MapIndex(reflect.ValueOf(key))
		fmt.Fprintf(&b, `%q:%v%#v,`, key, space, value)
		if i+1 != v.Len() {
			if (i+1)%width == 0 {
				b.WriteString("\n")
				b.WriteString(indent)
			} else {
				b.WriteString(" ")
			}
		}
		i++
	}

	return b.String()
}

func exportSlice(v reflect.Value, width int, indent string) string {
	b := strings.Builder{}
	for i := 0; i < v.Len(); i++ {
		fmt.Fprintf(&b, `%#v,`, v.Index(i).Interface())
		if i+1 != v.Len() {
			if (i+1)%width == 0 {
				b.WriteString("\n")
				b.WriteString(indent)
			} else {
				b.WriteString(" ")
			}
		}
	}

	return b.String()
}
