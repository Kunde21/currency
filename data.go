// Code generated by go generate; DO NOT EDIT.
//go:generate go run gen.go

package currency

// CLDRVersion is the CLDR version from which the data is derived.
const CLDRVersion = "39.0.0"

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
	pattern               string
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
	"AUD", "CAD", "CHF", "EUR", "GBP", "JPY", "NOK", "NZD", "SEK", "USD",

	// Other currencies.
	"AED", "AFN", "ALL", "AMD", "ANG", "AOA", "ARS", "AWG", "AZN", "BAM",
	"BBD", "BDT", "BGN", "BHD", "BIF", "BMD", "BND", "BOB", "BRL", "BSD",
	"BTN", "BWP", "BYN", "BZD", "CDF", "CLP", "CNY", "COP", "CRC", "CUC",
	"CUP", "CVE", "CZK", "DJF", "DKK", "DOP", "DZD", "EGP", "ERN", "ETB",
	"FJD", "FKP", "GEL", "GHS", "GIP", "GMD", "GNF", "GTQ", "GYD", "HKD",
	"HNL", "HRK", "HTG", "HUF", "IDR", "ILS", "INR", "IQD", "IRR", "ISK",
	"JMD", "JOD", "KES", "KGS", "KHR", "KMF", "KPW", "KRW", "KWD", "KYD",
	"KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LYD", "MAD", "MDL", "MGA",
	"MKD", "MMK", "MNT", "MOP", "MRU", "MUR", "MVR", "MWK", "MXN", "MYR",
	"MZN", "NAD", "NGN", "NIO", "NPR", "OMR", "PAB", "PEN", "PGK", "PHP",
	"PKR", "PLN", "PYG", "QAR", "RON", "RSD", "RUB", "RWF", "SAR", "SBD",
	"SCR", "SDG", "SGD", "SHP", "SLL", "SOS", "SRD", "SSP", "STN", "SVC",
	"SYP", "SZL", "THB", "TJS", "TMT", "TND", "TOP", "TRY", "TTD", "TWD",
	"TZS", "UAH", "UGX", "UYU", "UYW", "UZS", "VES", "VND", "VUV", "WST",
	"XAF", "XCD", "XOF", "XPF", "YER", "ZAR", "ZMW", "ZWL",
}

var currencies = map[string]currencyInfo{
	"AED": {"784", 2}, "AFN": {"971", 0}, "ALL": {"008", 0},
	"AMD": {"051", 2}, "ANG": {"532", 2}, "AOA": {"973", 2},
	"ARS": {"032", 2}, "AUD": {"036", 2}, "AWG": {"533", 2},
	"AZN": {"944", 2}, "BAM": {"977", 2}, "BBD": {"052", 2},
	"BDT": {"050", 2}, "BGN": {"975", 2}, "BHD": {"048", 3},
	"BIF": {"108", 0}, "BMD": {"060", 2}, "BND": {"096", 2},
	"BOB": {"068", 2}, "BRL": {"986", 2}, "BSD": {"044", 2},
	"BTN": {"064", 2}, "BWP": {"072", 2}, "BYN": {"933", 2},
	"BZD": {"084", 2}, "CAD": {"124", 2}, "CDF": {"976", 2},
	"CHF": {"756", 2}, "CLP": {"152", 0}, "CNY": {"156", 2},
	"COP": {"170", 2}, "CRC": {"188", 2}, "CUC": {"931", 2},
	"CUP": {"192", 2}, "CVE": {"132", 2}, "CZK": {"203", 2},
	"DJF": {"262", 0}, "DKK": {"208", 2}, "DOP": {"214", 2},
	"DZD": {"012", 2}, "EGP": {"818", 2}, "ERN": {"232", 2},
	"ETB": {"230", 2}, "EUR": {"978", 2}, "FJD": {"242", 2},
	"FKP": {"238", 2}, "GBP": {"826", 2}, "GEL": {"981", 2},
	"GHS": {"936", 2}, "GIP": {"292", 2}, "GMD": {"270", 2},
	"GNF": {"324", 0}, "GTQ": {"320", 2}, "GYD": {"328", 2},
	"HKD": {"344", 2}, "HNL": {"340", 2}, "HRK": {"191", 2},
	"HTG": {"332", 2}, "HUF": {"348", 2}, "IDR": {"360", 2},
	"ILS": {"376", 2}, "INR": {"356", 2}, "IQD": {"368", 0},
	"IRR": {"364", 0}, "ISK": {"352", 0}, "JMD": {"388", 2},
	"JOD": {"400", 3}, "JPY": {"392", 0}, "KES": {"404", 2},
	"KGS": {"417", 2}, "KHR": {"116", 2}, "KMF": {"174", 0},
	"KPW": {"408", 0}, "KRW": {"410", 0}, "KWD": {"414", 3},
	"KYD": {"136", 2}, "KZT": {"398", 2}, "LAK": {"418", 0},
	"LBP": {"422", 0}, "LKR": {"144", 2}, "LRD": {"430", 2},
	"LSL": {"426", 2}, "LYD": {"434", 3}, "MAD": {"504", 2},
	"MDL": {"498", 2}, "MGA": {"969", 0}, "MKD": {"807", 2},
	"MMK": {"104", 0}, "MNT": {"496", 2}, "MOP": {"446", 2},
	"MRU": {"929", 2}, "MUR": {"480", 2}, "MVR": {"462", 2},
	"MWK": {"454", 2}, "MXN": {"484", 2}, "MYR": {"458", 2},
	"MZN": {"943", 2}, "NAD": {"516", 2}, "NGN": {"566", 2},
	"NIO": {"558", 2}, "NOK": {"578", 2}, "NPR": {"524", 2},
	"NZD": {"554", 2}, "OMR": {"512", 3}, "PAB": {"590", 2},
	"PEN": {"604", 2}, "PGK": {"598", 2}, "PHP": {"608", 2},
	"PKR": {"586", 2}, "PLN": {"985", 2}, "PYG": {"600", 0},
	"QAR": {"634", 2}, "RON": {"946", 2}, "RSD": {"941", 0},
	"RUB": {"643", 2}, "RWF": {"646", 0}, "SAR": {"682", 2},
	"SBD": {"090", 2}, "SCR": {"690", 2}, "SDG": {"938", 2},
	"SEK": {"752", 2}, "SGD": {"702", 2}, "SHP": {"654", 2},
	"SLL": {"694", 0}, "SOS": {"706", 0}, "SRD": {"968", 2},
	"SSP": {"728", 2}, "STN": {"930", 2}, "SVC": {"222", 2},
	"SYP": {"760", 0}, "SZL": {"748", 2}, "THB": {"764", 2},
	"TJS": {"972", 2}, "TMT": {"934", 2}, "TND": {"788", 3},
	"TOP": {"776", 2}, "TRY": {"949", 2}, "TTD": {"780", 2},
	"TWD": {"901", 2}, "TZS": {"834", 2}, "UAH": {"980", 2},
	"UGX": {"800", 0}, "USD": {"840", 2}, "UYU": {"858", 2},
	"UYW": {"927", 4}, "UZS": {"860", 2}, "VES": {"928", 2},
	"VND": {"704", 0}, "VUV": {"548", 0}, "WST": {"882", 2},
	"XAF": {"950", 0}, "XCD": {"951", 2}, "XOF": {"952", 0},
	"XPF": {"953", 0}, "YER": {"886", 0}, "ZAR": {"710", 2},
	"ZMW": {"967", 2}, "ZWL": {"932", 2},
}

var currencySymbols = map[string][]symbolInfo{
	"AED": {
		{"AED", []string{"en"}},
		{"د.إ.\u200f", []string{"ar"}},
	},
	"AFN": {
		{"AFN", []string{"en"}},
		{"؋", []string{"fa", "ps"}},
	},
	"ALL": {
		{"ALL", []string{"en"}},
		{"Lekë", []string{"sq"}},
	},
	"AMD": {
		{"AMD", []string{"en"}},
		{"֏", []string{"hy"}},
	},
	"ANG": {
		{"ANG", []string{"en"}},
		{"NAf", []string{"my"}},
		{"NAf.", []string{"en-SX", "nl-CW", "nl-SX"}},
	},
	"AOA": {
		{"AOA", []string{"en"}},
		{"Kz", []string{"pt-AO"}},
	},
	"ARS": {
		{"ARS", []string{"en", "fr-CA"}},
		{"$", []string{"es-AR"}},
		{"$AR", []string{"fr"}},
	},
	"AUD": {
		{"A$", []string{"en"}},
		{"$", []string{"en-AU", "en-CC", "en-CX", "en-KI", "en-NF", "en-NR", "en-TV"}},
		{"$AU", []string{"fr"}},
		{"$\u00a0AU", []string{"fr-CA"}},
		{"AU$", []string{"am", "ar", "ca", "cs", "da", "de", "et", "id", "ko", "lv", "nl", "pt", "th", "tr", "vi", "yue", "zh", "zh-Hant"}},
	},
	"AWG": {
		{"AWG", []string{"en"}},
		{"Afl", []string{"my"}},
		{"Afl.", []string{"nl-AW"}},
	},
	"AZN": {
		{"AZN", []string{"en"}},
		{"₼", []string{"az"}},
	},
	"BAM": {
		{"BAM", []string{"en"}},
		{"KM", []string{"bs", "hr-BA", "sr-Latn"}},
		{"КМ", []string{"sr"}},
	},
	"BBD": {
		{"BBD", []string{"en"}},
		{"$", []string{"en-BB"}},
		{"Bds$", []string{"sv"}},
		{"DBB", []string{"so"}},
	},
	"BDT": {
		{"BDT", []string{"en"}},
		{"৳", []string{"bn"}},
	},
	"BGN": {
		{"BGN", []string{"en"}},
		{"лв.", []string{"bg"}},
	},
	"BHD": {
		{"BHD", []string{"en"}},
		{"د.ب.\u200f", []string{"ar"}},
	},
	"BIF": {
		{"BIF", []string{"en"}},
		{"FBu", []string{"en-BI", "fr-BI"}},
	},
	"BMD": {
		{"BMD", []string{"en", "fr-CA"}},
		{"$", []string{"en-BM"}},
		{"$BM", []string{"fr"}},
		{"BM$", []string{"sv"}},
	},
	"BND": {
		{"BND", []string{"en", "fr-CA"}},
		{"$", []string{"ms-BN"}},
		{"$BN", []string{"fr"}},
	},
	"BOB": {
		{"BOB", []string{"en"}},
		{"Bs", []string{"es-BO"}},
	},
	"BRL": {
		{"R$", []string{"en"}},
		{"BR$", []string{"sv"}},
	},
	"BSD": {
		{"BSD", []string{"en"}},
		{"$", []string{"en-BS"}},
		{"BS$", []string{"sv"}},
	},
	"BWP": {
		{"BWP", []string{"en"}},
		{"P", []string{"en-BW"}},
	},
	"BYN": {
		{"BYN", []string{"en"}},
		{"Br", []string{"be", "ru-BY"}},
	},
	"BZD": {
		{"BZD", []string{"en", "fr-CA"}},
		{"$", []string{"en-BZ", "es-BZ"}},
		{"$BZ", []string{"fr"}},
		{"BZ$", []string{"sv"}},
	},
	"CAD": {
		{"CA$", []string{"en"}},
		{"$", []string{"en-CA"}},
		{"$CA", []string{"fa", "fr"}},
		{"$\u00a0CA", []string{"fr-CA"}},
		{"C$", []string{"nl"}},
	},
	"CDF": {
		{"CDF", []string{"en"}},
		{"FC", []string{"fr-CD", "sw-CD"}},
	},
	"CLP": {
		{"CLP", []string{"en", "fr-CA"}},
		{"$", []string{"es-CL"}},
		{"$CL", []string{"fr"}},
	},
	"CNY": {
		{"CN¥", []string{"en", "zh-Hans-HK", "zh-Hans-MO", "zh-Hans-SG"}},
		{"¥", []string{"zh"}},
		{"¥CN", []string{"fa"}},
		{"\u200eCN¥\u200e", []string{"he"}},
		{"元", []string{"ja"}},
	},
	"COP": {
		{"COP", []string{"en", "fr-CA"}},
		{"$", []string{"es-CO"}},
		{"$CO", []string{"fr"}},
	},
	"CRC": {
		{"CRC", []string{"en"}},
		{"₡", []string{"es-CR"}},
	},
	"CUP": {
		{"CUP", []string{"en"}},
		{"$", []string{"es-CU"}},
	},
	"CVE": {
		{"CVE", []string{"en"}},
		{"\u200b", []string{"pt-CV"}},
	},
	"CZK": {
		{"CZK", []string{"en"}},
		{"Kč", []string{"cs"}},
	},
	"DJF": {
		{"DJF", []string{"en"}},
		{"Fdj", []string{"ar-DJ", "fr-DJ", "so-DJ"}},
	},
	"DKK": {
		{"DKK", []string{"en"}},
		{"Dkr", []string{"sv"}},
		{"kr.", []string{"da", "en-DK"}},
	},
	"DOP": {
		{"DOP", []string{"en"}},
		{"RD$", []string{"es-DO", "sv"}},
	},
	"DZD": {
		{"DZD", []string{"en"}},
		{"DA", []string{"fr-DZ"}},
		{"د.ج.\u200f", []string{"ar"}},
	},
	"EGP": {
		{"EGP", []string{"en"}},
		{"EG£", []string{"sv"}},
		{"ج.م.\u200f", []string{"ar"}},
	},
	"ERN": {
		{"ERN", []string{"en"}},
		{"Nfk", []string{"ar-ER", "en-ER"}},
	},
	"ETB": {
		{"ETB", []string{"en"}},
		{"Br", []string{"so-ET"}},
		{"ብር", []string{"am"}},
	},
	"EUR": {
		{"€", []string{"en"}},
	},
	"FJD": {
		{"FJD", []string{"en", "fr-CA"}},
		{"$", []string{"en-FJ"}},
		{"$FJ", []string{"fr"}},
		{"FJ$", []string{"nl"}},
	},
	"FKP": {
		{"FKP", []string{"en", "fr-CA"}},
		{"£", []string{"en-FK"}},
		{"£FK", []string{"fr"}},
	},
	"GBP": {
		{"£", []string{"en", "fr-CA"}},
		{"GB£", []string{"ar-SS", "en-FK", "en-GI", "en-MT", "en-SH", "en-SS"}},
		{"UK£", []string{"ar"}},
		{"£GB", []string{"fr"}},
	},
	"GEL": {
		{"GEL", []string{"en"}},
		{"₾", []string{"ka"}},
	},
	"GHS": {
		{"GHS", []string{"en"}},
		{"GH₵", []string{"en-GH"}},
	},
	"GIP": {
		{"GIP", []string{"en", "fr-CA"}},
		{"£", []string{"en-GI"}},
		{"£GI", []string{"fr"}},
	},
	"GMD": {
		{"GMD", []string{"en"}},
		{"D", []string{"en-GM"}},
	},
	"GNF": {
		{"GNF", []string{"en"}},
		{"FG", []string{"fr-GN"}},
	},
	"GTQ": {
		{"GTQ", []string{"en"}},
		{"Q", []string{"es-GT"}},
	},
	"GYD": {
		{"GYD", []string{"en"}},
		{"$", []string{"en-GY"}},
	},
	"HKD": {
		{"HK$", []string{"en"}},
		{"$HK", []string{"fa"}},
		{"$\u00a0HK", []string{"fr-CA"}},
	},
	"HNL": {
		{"HNL", []string{"en"}},
		{"L", []string{"es-HN"}},
	},
	"HRK": {
		{"HRK", []string{"en"}},
		{"kn", []string{"bs"}},
	},
	"HTG": {
		{"HTG", []string{"en"}},
		{"G", []string{"fr-HT", "my"}},
	},
	"HUF": {
		{"HUF", []string{"en"}},
		{"Ft", []string{"hu"}},
	},
	"IDR": {
		{"IDR", []string{"en"}},
		{"Rp", []string{"id", "ms-ID"}},
	},
	"ILS": {
		{"₪", []string{"en"}},
		{"NIS", []string{"sk"}},
	},
	"INR": {
		{"₹", []string{"en"}},
		{"Rs", []string{"id"}},
	},
	"IQD": {
		{"IQD", []string{"en"}},
		{"د.ع.\u200f", []string{"ar"}},
	},
	"IRR": {
		{"IRR", []string{"en"}},
		{"ر.إ.", []string{"ar"}},
		{"ریال", []string{"fa"}},
	},
	"ISK": {
		{"ISK", []string{"en"}},
		{"Ikr", []string{"sv"}},
	},
	"JMD": {
		{"JMD", []string{"en"}},
		{"$", []string{"en-JM"}},
		{"JM$", []string{"sv"}},
	},
	"JOD": {
		{"JOD", []string{"en"}},
		{"د.أ.\u200f", []string{"ar"}},
	},
	"JPY": {
		{"¥", []string{"en", "en-AU"}},
		{"JP¥", []string{"af", "am", "ar", "as", "az", "bn", "cs", "cy", "da", "el", "en-001", "eu", "gl", "gu", "hi", "hy", "id", "is", "kk", "km", "ko", "ky", "lo", "mn", "mr", "ms", "my", "ne", "nl", "pa", "ps", "pt", "si", "so", "sq", "sw", "te", "tk", "ur", "uz", "zh", "zu"}},
		{"￥", []string{"ja"}},
	},
	"KES": {
		{"KES", []string{"en"}},
		{"Ksh", []string{"en-KE", "so-KE", "sw"}},
	},
	"KGS": {
		{"KGS", []string{"en"}},
		{"сом", []string{"ky", "ru-KG"}},
	},
	"KHR": {
		{"KHR", []string{"en"}},
		{"៛", []string{"km"}},
	},
	"KMF": {
		{"KMF", []string{"en"}},
		{"CF", []string{"ar-KM", "fr-KM"}},
	},
	"KRW": {
		{"₩", []string{"en", "zh-Hant-HK"}},
		{"￦", []string{"yue", "zh", "zh-Hant"}},
	},
	"KWD": {
		{"KWD", []string{"en"}},
		{"د.ك.\u200f", []string{"ar"}},
	},
	"KYD": {
		{"KYD", []string{"en"}},
		{"$", []string{"en-KY"}},
	},
	"KZT": {
		{"KZT", []string{"en"}},
		{"₸", []string{"kk", "ru-KZ"}},
	},
	"LAK": {
		{"LAK", []string{"en"}},
		{"₭", []string{"lo"}},
	},
	"LBP": {
		{"LBP", []string{"en", "fr-CA"}},
		{"£LB", []string{"fr"}},
		{"ل.ل.\u200f", []string{"ar"}},
	},
	"LKR": {
		{"LKR", []string{"en"}},
		{"Rs.", []string{"ta-LK"}},
		{"රු.", []string{"si"}},
	},
	"LRD": {
		{"LRD", []string{"en"}},
		{"$", []string{"en-LR"}},
	},
	"LYD": {
		{"LYD", []string{"en"}},
		{"د.ل.\u200f", []string{"ar"}},
	},
	"MAD": {
		{"MAD", []string{"en"}},
		{"د.م.\u200f", []string{"ar"}},
	},
	"MDL": {
		{"MDL", []string{"en"}},
		{"L", []string{"ro-MD", "ru-MD"}},
	},
	"MGA": {
		{"MGA", []string{"en"}},
		{"Ar", []string{"en-MG", "fr-MG"}},
	},
	"MKD": {
		{"MKD", []string{"en"}},
		{"den", []string{"sq-MK"}},
		{"ден.", []string{"mk"}},
	},
	"MMK": {
		{"MMK", []string{"en"}},
		{"K", []string{"my"}},
	},
	"MNT": {
		{"MNT", []string{"en"}},
		{"₮", []string{"mn"}},
	},
	"MOP": {
		{"MOP", []string{"en"}},
		{"MOP$", []string{"en-MO", "pt-MO", "zh-Hans-MO", "zh-Hant-MO"}},
	},
	"MRU": {
		{"MRU", []string{"en"}},
		{"UM", []string{"es-MX", "fr-MR"}},
		{"أ.م.", []string{"ar"}},
	},
	"MUR": {
		{"MUR", []string{"en"}},
		{"Rs", []string{"en-MU", "fr-MU"}},
	},
	"MWK": {
		{"MWK", []string{"en"}},
		{"MK", []string{"en-MW"}},
	},
	"MXN": {
		{"MX$", []string{"en", "fr-CA"}},
		{"$", []string{"es-MX"}},
		{"$MX", []string{"fa", "fr", "gl"}},
	},
	"MYR": {
		{"MYR", []string{"en"}},
		{"RM", []string{"en-MY", "ms", "ta-MY", "ta-SG"}},
	},
	"MZN": {
		{"MZN", []string{"en"}},
		{"MTn", []string{"pt-MZ"}},
	},
	"NAD": {
		{"NAD", []string{"en", "fr-CA"}},
		{"$", []string{"af-NA", "en-NA"}},
		{"$NA", []string{"fr"}},
	},
	"NGN": {
		{"NGN", []string{"en"}},
		{"₦", []string{"en-NG"}},
	},
	"NIO": {
		{"NIO", []string{"en"}},
		{"C$", []string{"es-NI"}},
	},
	"NOK": {
		{"NOK", []string{"en"}},
		{"Nkr", []string{"sv"}},
		{"kr", []string{"nb", "nn", "no"}},
	},
	"NPR": {
		{"NPR", []string{"en"}},
		{"नेरू", []string{"ne"}},
	},
	"NZD": {
		{"NZ$", []string{"en"}},
		{"$", []string{"en-CK", "en-NU", "en-NZ", "en-PN", "en-TK"}},
		{"$NZ", []string{"fa", "fr"}},
		{"$\u00a0NZ", []string{"fr-CA"}},
	},
	"OMR": {
		{"OMR", []string{"en"}},
		{"ر.ع.\u200f", []string{"ar"}},
	},
	"PAB": {
		{"PAB", []string{"en"}},
		{"B/.", []string{"es-PA", "my"}},
	},
	"PEN": {
		{"PEN", []string{"en"}},
		{"S/", []string{"es-PE"}},
	},
	"PGK": {
		{"PGK", []string{"en"}},
		{"K", []string{"en-PG"}},
	},
	"PHP": {
		{"₱", []string{"en"}},
	},
	"PKR": {
		{"PKR", []string{"en", "ur-IN"}},
		{"Rs", []string{"en-PK", "ps-PK", "ur"}},
	},
	"PLN": {
		{"PLN", []string{"en"}},
		{"zł", []string{"pl"}},
	},
	"PYG": {
		{"PYG", []string{"en"}},
		{"Gs.", []string{"es-PY"}},
	},
	"QAR": {
		{"QAR", []string{"en"}},
		{"ر.ق.\u200f", []string{"ar"}},
	},
	"RSD": {
		{"RSD", []string{"en"}},
		{"din.", []string{"bs"}},
	},
	"RUB": {
		{"RUB", []string{"en"}},
		{"₽", []string{"be", "kk", "ru"}},
	},
	"RWF": {
		{"RWF", []string{"en"}},
		{"RF", []string{"en-RW", "fr-RW"}},
	},
	"SAR": {
		{"SAR", []string{"en"}},
		{"ر.س.\u200f", []string{"ar"}},
	},
	"SBD": {
		{"SBD", []string{"en", "fr-CA"}},
		{"$", []string{"en-SB"}},
		{"$SB", []string{"fr"}},
		{"SI$", []string{"nl"}},
	},
	"SCR": {
		{"SCR", []string{"en"}},
		{"Rs", []string{"en-AU"}},
		{"SR", []string{"en-SC", "fr-SC"}},
	},
	"SDG": {
		{"SDG", []string{"ar-LB", "en"}},
		{"ج.س.", []string{"ar"}},
	},
	"SEK": {
		{"SEK", []string{"en"}},
		{"kr", []string{"en-SE", "sv"}},
	},
	"SGD": {
		{"SGD", []string{"en"}},
		{"$", []string{"en-SG", "ms-SG", "ta-SG", "zh-Hans-SG"}},
		{"$SG", []string{"fr"}},
		{"$\u00a0SG", []string{"fr-CA"}},
		{"S$", []string{"ta-MY"}},
	},
	"SHP": {
		{"SHP", []string{"en"}},
		{"£", []string{"en-SH"}},
	},
	"SLL": {
		{"SLL", []string{"en"}},
		{"Le", []string{"en-SL"}},
	},
	"SOS": {
		{"SOS", []string{"en"}},
		{"S", []string{"ar-SO", "so"}},
	},
	"SRD": {
		{"SRD", []string{"en", "fr-CA"}},
		{"$", []string{"nl-SR"}},
		{"$SR", []string{"fr"}},
	},
	"SSP": {
		{"SSP", []string{"en"}},
		{"£", []string{"ar-SS", "en-SS"}},
	},
	"STN": {
		{"STN", []string{"en"}},
		{"Db", []string{"pt-ST"}},
	},
	"SYP": {
		{"SYP", []string{"en"}},
		{"LS", []string{"fr-SY"}},
		{"ل.س.\u200f", []string{"ar"}},
	},
	"SZL": {
		{"SZL", []string{"en"}},
		{"E", []string{"en-SZ"}},
	},
	"THB": {
		{"THB", []string{"en", "es-419"}},
		{"฿", []string{"af", "am", "ar", "az", "bn", "bs", "ca", "cy", "da", "de", "el", "es", "et", "eu", "fa", "fil", "ga", "gl", "gu", "he", "hi", "hy", "id", "it", "kk", "km", "ky", "lo", "lv", "mn", "mr", "my", "ne", "nl", "pa", "pt", "ru", "si", "sq", "sw", "ta", "te", "th", "tr", "ur", "vi", "zu"}},
	},
	"TMT": {
		{"TMT", []string{"en"}},
		{"ТМТ", []string{"ru"}},
	},
	"TND": {
		{"TND", []string{"en"}},
		{"DT", []string{"fr-TN"}},
		{"د.ت.\u200f", []string{"ar"}},
	},
	"TOP": {
		{"TOP", []string{"en"}},
		{"T$", []string{"en-TO"}},
	},
	"TRY": {
		{"TRY", []string{"en"}},
		{"₺", []string{"tr"}},
	},
	"TTD": {
		{"TTD", []string{"en", "fr-CA"}},
		{"$", []string{"en-TT"}},
		{"$TT", []string{"fr"}},
		{"TT$", []string{"my"}},
	},
	"TWD": {
		{"NT$", []string{"en", "zh-Hant-HK"}},
		{"$", []string{"zh-Hant"}},
	},
	"TZS": {
		{"TZS", []string{"en"}},
		{"TSh", []string{"en-TZ", "sw"}},
	},
	"UAH": {
		{"UAH", []string{"en"}},
		{"₴", []string{"ru", "uk"}},
	},
	"UGX": {
		{"UGX", []string{"en"}},
		{"USh", []string{"en-UG", "sw-UG"}},
	},
	"USD": {
		{"$", []string{"en", "en-IN", "es-419", "nl-BQ", "sw-KE"}},
		{"$US", []string{"fr"}},
		{"$\u00a0US", []string{"fr-CA"}},
		{"US$", []string{"am", "ar", "as", "az", "bn", "cs", "cy", "da", "en-001", "es", "es-AR", "es-CL", "es-CO", "es-CU", "es-DO", "es-UY", "eu", "gu", "id", "ka", "ko", "lo", "mk", "my", "ne", "nl", "pa", "pt", "si", "so", "sq", "sr", "sr-Latn", "sv", "sw", "ta-SG", "th", "tk", "uz", "vi", "yue", "zh", "zh-Hant"}},
		{"щ.д.", []string{"bg"}},
	},
	"UYU": {
		{"UYU", []string{"en", "fr-CA"}},
		{"$", []string{"es-UY"}},
		{"$UY", []string{"fr"}},
	},
	"UYW": {
		{"UYW", []string{"en"}},
		{"UP", []string{"es-UY"}},
	},
	"UZS": {
		{"UZS", []string{"en"}},
		{"soʻm", []string{"uz"}},
	},
	"VES": {
		{"VES", []string{"en"}},
		{"Bs.S", []string{"es-VE"}},
	},
	"VND": {
		{"₫", []string{"en"}},
	},
	"VUV": {
		{"VUV", []string{"en"}},
		{"VT", []string{"en-VU", "fr-VU"}},
	},
	"WST": {
		{"WST", []string{"en", "fr-CA"}},
		{"$WS", []string{"fr"}},
		{"WS$", []string{"en-WS"}},
	},
	"XAF": {
		{"FCFA", []string{"en"}},
	},
	"XCD": {
		{"EC$", []string{"en"}},
		{"$", []string{"en-AG", "en-AI", "en-DM", "en-GD", "en-KN", "en-LC", "en-MS", "en-VC"}},
		{"$EC", []string{"fa"}},
	},
	"XOF": {
		{"F\u202fCFA", []string{"en"}},
		{"සිෆ්එ", []string{"si"}},
	},
	"XPF": {
		{"CFPF", []string{"en", "fr-CA"}},
		{"CFP", []string{"en-AU"}},
		{"FCFP", []string{"fr"}},
	},
	"YER": {
		{"YER", []string{"en"}},
		{"ر.ي.\u200f", []string{"ar"}},
	},
	"ZAR": {
		{"ZAR", []string{"en"}},
		{"R", []string{"af", "en-LS", "en-ZA", "zu"}},
	},
	"ZMW": {
		{"ZMW", []string{"en"}},
		{"K", []string{"en-ZM"}},
	},
}

var currencyFormats = map[string]currencyFormat{
	"af":      {"¤0.00", 0, 1, 3, 3, ",", "\u00a0", "+", "-"},
	"ar":      {"0.00\u00a0¤", 1, 1, 3, 3, "٫", "٬", "\u061c+", "\u061c-"},
	"ar-DZ":   {"¤\u00a00.00", 0, 1, 3, 3, ",", ".", "\u200e+", "\u200e-"},
	"ar-EH":   {"¤\u00a00.00", 0, 1, 3, 3, ".", ",", "\u200e+", "\u200e-"},
	"ar-LY":   {"¤\u00a00.00", 0, 1, 3, 3, ",", ".", "\u200e+", "\u200e-"},
	"ar-MA":   {"¤\u00a00.00", 0, 1, 3, 3, ",", ".", "\u200e+", "\u200e-"},
	"ar-TN":   {"¤\u00a00.00", 0, 1, 3, 3, ",", ".", "\u200e+", "\u200e-"},
	"as":      {"¤\u00a00.00", 3, 1, 3, 2, ".", ",", "+", "-"},
	"az":      {"0.00\u00a0¤", 0, 1, 3, 3, ",", ".", "+", "-"},
	"be":      {"0.00\u00a0¤", 0, 2, 3, 3, ",", "\u00a0", "+", "-"},
	"bg":      {"0.00\u00a0¤", 0, 2, 0, 0, ",", "\u00a0", "+", "-"},
	"bn":      {"0.00¤", 3, 1, 3, 2, ".", ",", "+", "-"},
	"bs":      {"0.00\u00a0¤", 0, 1, 3, 3, ",", ".", "+", "-"},
	"ca":      {"0.00\u00a0¤", 0, 1, 3, 3, ",", ".", "+", "-"},
	"cs":      {"0.00\u00a0¤", 0, 1, 3, 3, ",", "\u00a0", "+", "-"},
	"da":      {"0.00\u00a0¤", 0, 1, 3, 3, ",", ".", "+", "-"},
	"de":      {"0.00\u00a0¤", 0, 1, 3, 3, ",", ".", "+", "-"},
	"de-AT":   {"¤\u00a00.00", 0, 1, 3, 3, ",", ".", "+", "-"},
	"de-CH":   {"¤\u00a00.00;¤-0.00", 0, 1, 3, 3, ".", "’", "+", "-"},
	"de-LI":   {"¤\u00a00.00", 0, 1, 3, 3, ".", "’", "+", "-"},
	"el":      {"0.00\u00a0¤", 0, 1, 3, 3, ",", ".", "+", "-"},
	"en":      {"¤0.00", 0, 1, 3, 3, ".", ",", "+", "-"},
	"en-150":  {"0.00\u00a0¤", 0, 1, 3, 3, ".", ",", "+", "-"},
	"en-AT":   {"¤\u00a00.00", 0, 1, 3, 3, ",", ".", "+", "-"},
	"en-BE":   {"0.00\u00a0¤", 0, 1, 3, 3, ",", ".", "+", "-"},
	"en-CH":   {"¤\u00a00.00;¤-0.00", 0, 1, 3, 3, ".", "’", "+", "-"},
	"en-DE":   {"0.00\u00a0¤", 0, 1, 3, 3, ",", ".", "+", "-"},
	"en-DK":   {"0.00\u00a0¤", 0, 1, 3, 3, ",", ".", "+", "-"},
	"en-FI":   {"0.00\u00a0¤", 0, 1, 3, 3, ",", "\u00a0", "+", "-"},
	"en-IN":   {"¤0.00", 0, 1, 3, 2, ".", ",", "+", "-"},
	"en-NL":   {"¤\u00a00.00;¤\u00a0-0.00", 0, 1, 3, 3, ",", ".", "+", "-"},
	"en-SE":   {"0.00\u00a0¤", 0, 1, 3, 3, ",", "\u00a0", "+", "-"},
	"en-SI":   {"0.00\u00a0¤", 0, 1, 3, 3, ",", ".", "+", "-"},
	"en-ZA":   {"¤0.00", 0, 1, 3, 3, ",", "\u00a0", "+", "-"},
	"es":      {"0.00\u00a0¤", 0, 2, 3, 3, ",", ".", "+", "-"},
	"es-419":  {"¤0.00", 0, 1, 3, 3, ".", ",", "+", "-"},
	"es-AR":   {"¤\u00a00.00", 0, 1, 3, 3, ",", ".", "+", "-"},
	"es-BO":   {"¤0.00", 0, 1, 3, 3, ",", ".", "+", "-"},
	"es-CL":   {"¤0.00;¤-0.00", 0, 1, 3, 3, ",", ".", "+", "-"},
	"es-CO":   {"¤\u00a00.00", 0, 1, 3, 3, ",", ".", "+", "-"},
	"es-CR":   {"¤0.00", 0, 1, 3, 3, ",", "\u00a0", "+", "-"},
	"es-EC":   {"¤0.00;¤-0.00", 0, 1, 3, 3, ",", ".", "+", "-"},
	"es-GQ":   {"¤0.00", 0, 2, 3, 3, ",", ".", "+", "-"},
	"es-PE":   {"¤\u00a00.00", 0, 1, 3, 3, ".", ",", "+", "-"},
	"es-PY":   {"¤\u00a00.00;¤\u00a0-0.00", 0, 1, 3, 3, ",", ".", "+", "-"},
	"es-UY":   {"¤\u00a00.00", 0, 1, 3, 3, ",", ".", "+", "-"},
	"es-VE":   {"¤0.00;¤-0.00", 0, 1, 3, 3, ",", ".", "+", "-"},
	"et":      {"0.00\u00a0¤", 0, 2, 3, 3, ",", "\u00a0", "+", "−"},
	"eu":      {"0.00\u00a0¤", 0, 1, 3, 3, ",", ".", "+", "−"},
	"fa":      {"\u200e¤0.00", 2, 1, 3, 3, "٫", "٬", "\u200e+", "\u200e−"},
	"fa-AF":   {"¤\u00a00.00", 2, 1, 3, 3, "٫", "٬", "\u200e+", "\u200e−"},
	"fi":      {"0.00\u00a0¤", 0, 1, 3, 3, ",", "\u00a0", "+", "−"},
	"fr":      {"0.00\u00a0¤", 0, 1, 3, 3, ",", "\u202f", "+", "-"},
	"fr-CA":   {"0.00\u00a0¤", 0, 1, 3, 3, ",", "\u00a0", "+", "-"},
	"fr-CH":   {"0.00\u00a0¤", 0, 1, 3, 3, ".", "\u202f", "+", "-"},
	"fr-LU":   {"0.00\u00a0¤", 0, 1, 3, 3, ",", ".", "+", "-"},
	"fr-MA":   {"0.00\u00a0¤", 0, 1, 3, 3, ",", ".", "+", "-"},
	"gl":      {"0.00\u00a0¤", 0, 1, 3, 3, ",", ".", "+", "-"},
	"gu":      {"¤0.00", 0, 1, 3, 2, ".", ",", "+", "-"},
	"he":      {"\u200f0.00\u00a0¤;\u200f-0.00\u00a0¤", 0, 1, 3, 3, ".", ",", "\u200e+", "\u200e-"},
	"hi":      {"¤0.00", 0, 1, 3, 2, ".", ",", "+", "-"},
	"hr":      {"0.00\u00a0¤", 0, 1, 3, 3, ",", ".", "+", "−"},
	"hu":      {"0.00\u00a0¤", 0, 1, 3, 3, ",", "\u00a0", "+", "-"},
	"hy":      {"0.00\u00a0¤", 0, 1, 3, 3, ",", "\u00a0", "+", "-"},
	"id":      {"¤0.00", 0, 1, 3, 3, ",", ".", "+", "-"},
	"is":      {"0.00\u00a0¤", 0, 1, 3, 3, ",", ".", "+", "-"},
	"it":      {"0.00\u00a0¤", 0, 1, 3, 3, ",", ".", "+", "-"},
	"it-CH":   {"¤\u00a00.00;¤-0.00", 0, 1, 3, 3, ".", "’", "+", "-"},
	"ka":      {"0.00\u00a0¤", 0, 2, 3, 3, ",", "\u00a0", "+", "-"},
	"kk":      {"0.00\u00a0¤", 0, 1, 3, 3, ",", "\u00a0", "+", "-"},
	"km":      {"0.00¤", 0, 1, 3, 3, ",", ".", "+", "-"},
	"ky":      {"0.00\u00a0¤", 0, 1, 3, 3, ",", "\u00a0", "+", "-"},
	"lo":      {"¤0.00;¤-0.00", 0, 1, 3, 3, ",", ".", "+", "-"},
	"lt":      {"0.00\u00a0¤", 0, 1, 3, 3, ",", "\u00a0", "+", "−"},
	"lv":      {"0.00\u00a0¤", 0, 2, 3, 3, ",", "\u00a0", "+", "-"},
	"mk":      {"0.00\u00a0¤", 0, 1, 3, 3, ",", ".", "+", "-"},
	"mn":      {"¤\u00a00.00", 0, 1, 3, 3, ".", ",", "+", "-"},
	"mr":      {"¤0.00", 4, 1, 3, 3, ".", ",", "+", "-"},
	"ms-BN":   {"¤\u00a00.00", 0, 1, 3, 3, ",", ".", "+", "-"},
	"ms-ID":   {"¤0.00", 0, 1, 3, 3, ",", ".", "+", "-"},
	"my":      {"0.00\u00a0¤", 5, 1, 3, 3, ".", ",", "+", "-"},
	"nb":      {"¤\u00a00.00", 0, 1, 3, 3, ",", "\u00a0", "+", "−"},
	"ne":      {"¤\u00a00.00", 4, 1, 3, 2, ".", ",", "+", "-"},
	"nl":      {"¤\u00a00.00;¤\u00a0-0.00", 0, 1, 3, 3, ",", ".", "+", "-"},
	"nn":      {"0.00\u00a0¤", 0, 1, 3, 3, ",", "\u00a0", "+", "−"},
	"no":      {"¤\u00a00.00", 0, 1, 3, 3, ",", "\u00a0", "+", "−"},
	"pa":      {"¤\u00a00.00", 0, 1, 3, 2, ".", ",", "+", "-"},
	"pl":      {"0.00\u00a0¤", 0, 2, 3, 3, ",", "\u00a0", "+", "-"},
	"ps":      {"0.00\u00a0¤", 2, 1, 3, 3, "٫", "٬", "\u200e+\u200e", "\u200e-\u200e"},
	"pt":      {"¤\u00a00.00", 0, 1, 3, 3, ",", ".", "+", "-"},
	"pt-AO":   {"0.00\u00a0¤", 0, 1, 3, 3, ",", "\u00a0", "+", "-"},
	"pt-PT":   {"0.00\u00a0¤", 0, 2, 3, 3, ",", "\u00a0", "+", "-"},
	"ro":      {"0.00\u00a0¤", 0, 1, 3, 3, ",", ".", "+", "-"},
	"ru":      {"0.00\u00a0¤", 0, 1, 3, 3, ",", "\u00a0", "+", "-"},
	"ru-UA":   {"0.00\u00a0¤", 0, 2, 3, 3, ",", "\u00a0", "+", "-"},
	"sk":      {"0.00\u00a0¤", 0, 1, 3, 3, ",", "\u00a0", "+", "-"},
	"sl":      {"0.00\u00a0¤", 0, 1, 3, 3, ",", ".", "+", "−"},
	"sq":      {"0.00\u00a0¤", 0, 2, 3, 3, ",", "\u00a0", "+", "-"},
	"sr":      {"0.00\u00a0¤", 0, 1, 3, 3, ",", ".", "+", "-"},
	"sr-Latn": {"0.00\u00a0¤", 0, 1, 3, 3, ",", ".", "+", "-"},
	"sv":      {"0.00\u00a0¤", 0, 1, 3, 3, ",", "\u00a0", "+", "−"},
	"sw":      {"¤\u00a00.00", 0, 1, 3, 3, ".", ",", "+", "-"},
	"sw-CD":   {"¤\u00a00.00", 0, 1, 3, 3, ",", ".", "+", "-"},
	"ta":      {"¤\u00a00.00", 0, 1, 3, 2, ".", ",", "+", "-"},
	"ta-MY":   {"¤\u00a00.00", 0, 1, 3, 3, ".", ",", "+", "-"},
	"ta-SG":   {"¤\u00a00.00", 0, 1, 3, 3, ".", ",", "+", "-"},
	"te":      {"¤0.00", 0, 1, 3, 2, ".", ",", "+", "-"},
	"tk":      {"0.00\u00a0¤", 0, 1, 3, 3, ",", "\u00a0", "+", "-"},
	"tr":      {"¤0.00", 0, 1, 3, 3, ",", ".", "+", "-"},
	"uk":      {"0.00\u00a0¤", 0, 1, 3, 3, ",", "\u00a0", "+", "-"},
	"ur":      {"¤\u00a00.00", 0, 1, 3, 3, ".", ",", "\u200e+", "\u200e-"},
	"ur-IN":   {"¤\u00a00.00", 2, 1, 3, 2, "٫", "٬", "\u200e+\u200e", "\u200e-\u200e"},
	"uz":      {"0.00\u00a0¤", 0, 1, 3, 3, ",", "\u00a0", "+", "-"},
	"vi":      {"0.00\u00a0¤", 0, 1, 3, 3, ",", ".", "+", "-"},
}

var parentLocales = map[string]string{
	"az-Arab": "en", "az-Cyrl": "en", "blt-Latn": "en",
	"bs-Cyrl": "en", "en-150": "en-001", "en-AG": "en-001",
	"en-AI": "en-001", "en-AT": "en-150", "en-AU": "en-001",
	"en-BB": "en-001", "en-BE": "en-150", "en-BM": "en-001",
	"en-BS": "en-001", "en-BW": "en-001", "en-BZ": "en-001",
	"en-CA": "en-001", "en-CC": "en-001", "en-CH": "en-150",
	"en-CK": "en-001", "en-CM": "en-001", "en-CX": "en-001",
	"en-CY": "en-001", "en-DE": "en-150", "en-DG": "en-001",
	"en-DK": "en-150", "en-DM": "en-001", "en-ER": "en-001",
	"en-FI": "en-150", "en-FJ": "en-001", "en-FK": "en-001",
	"en-FM": "en-001", "en-GB": "en-001", "en-GD": "en-001",
	"en-GG": "en-001", "en-GH": "en-001", "en-GI": "en-001",
	"en-GM": "en-001", "en-GY": "en-001", "en-HK": "en-001",
	"en-IE": "en-001", "en-IL": "en-001", "en-IM": "en-001",
	"en-IN": "en-001", "en-IO": "en-001", "en-JE": "en-001",
	"en-JM": "en-001", "en-KE": "en-001", "en-KI": "en-001",
	"en-KN": "en-001", "en-KY": "en-001", "en-LC": "en-001",
	"en-LR": "en-001", "en-LS": "en-001", "en-MG": "en-001",
	"en-MO": "en-001", "en-MS": "en-001", "en-MT": "en-001",
	"en-MU": "en-001", "en-MW": "en-001", "en-MY": "en-001",
	"en-NA": "en-001", "en-NF": "en-001", "en-NG": "en-001",
	"en-NL": "en-150", "en-NR": "en-001", "en-NU": "en-001",
	"en-NZ": "en-001", "en-PG": "en-001", "en-PH": "en-001",
	"en-PK": "en-001", "en-PN": "en-001", "en-PW": "en-001",
	"en-RW": "en-001", "en-SB": "en-001", "en-SC": "en-001",
	"en-SD": "en-001", "en-SE": "en-150", "en-SG": "en-001",
	"en-SH": "en-001", "en-SI": "en-150", "en-SL": "en-001",
	"en-SS": "en-001", "en-SX": "en-001", "en-SZ": "en-001",
	"en-TC": "en-001", "en-TK": "en-001", "en-TO": "en-001",
	"en-TT": "en-001", "en-TV": "en-001", "en-TZ": "en-001",
	"en-UG": "en-001", "en-VC": "en-001", "en-VG": "en-001",
	"en-VU": "en-001", "en-WS": "en-001", "en-ZA": "en-001",
	"en-ZM": "en-001", "en-ZW": "en-001", "es-AR": "es-419",
	"es-BO": "es-419", "es-BR": "es-419", "es-BZ": "es-419",
	"es-CL": "es-419", "es-CO": "es-419", "es-CR": "es-419",
	"es-CU": "es-419", "es-DO": "es-419", "es-EC": "es-419",
	"es-GT": "es-419", "es-HN": "es-419", "es-MX": "es-419",
	"es-NI": "es-419", "es-PA": "es-419", "es-PE": "es-419",
	"es-PR": "es-419", "es-PY": "es-419", "es-SV": "es-419",
	"es-US": "es-419", "es-UY": "es-419", "es-VE": "es-419",
	"hi-Latn": "en", "iu-Latn": "en", "kk-Arab": "en",
	"ks-Deva": "en", "ku-Arab": "en", "ky-Arab": "en",
	"ky-Latn": "en", "mn-Mong": "en", "mni-Mtei": "en",
	"ms-Arab": "en", "nb": "no", "nn": "no",
	"pa-Arab": "en", "pt-AO": "pt-PT", "pt-CH": "pt-PT",
	"pt-CV": "pt-PT", "pt-FR": "pt-PT", "pt-GQ": "pt-PT",
	"pt-GW": "pt-PT", "pt-LU": "pt-PT", "pt-MO": "pt-PT",
	"pt-MZ": "pt-PT", "pt-ST": "pt-PT", "pt-TL": "pt-PT",
	"so-Arab": "en", "sr-Latn": "en", "sw-Arab": "en",
	"tg-Arab": "en", "ug-Cyrl": "en", "uz-Arab": "en",
	"uz-Cyrl": "en", "yue-Hans": "en", "zh-Hant": "en",
	"zh-Hant-MO": "zh-Hant-HK",
}
