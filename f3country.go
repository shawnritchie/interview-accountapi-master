package currency

import (
	"fmt"
	"strings"
)

type Country string
var zeroValueCountry = Country("")

func NewCountry(cc string) Country {
	return Country(strings.ToUpper(cc))
}

func (c *Country) IsValid() error {
	if _, ok := Countries[string(*c)]; !ok {
		return &InvalidCountry{*c}
	}
	return nil
}

func (c *Country) IsZeroValue() bool {
	return *c == zeroValueCountry
}

type InvalidCountry struct {
	country Country
}

func (e *InvalidCountry) Error() string {
	return fmt.Sprintf("invalid Country code %q. supported ISO 3166-1 formats e.g. 'GB', 'FR'", e.country)
}


var Countries = map[string]Country{
	"AF": Country("AF"),
	"AX": Country("AX"),
	"AL": Country("AL"),
	"DZ": Country("DZ"),
	"AS": Country("AS"),
	"AD": Country("AD"),
	"AO": Country("AO"),
	"AI": Country("AI"),
	"AQ": Country("AQ"),
	"AG": Country("AG"),
	"AR": Country("AR"),
	"AM": Country("AM"),
	"AW": Country("AW"),
	"AU": Country("AU"),
	"AT": Country("AT"),
	"AZ": Country("AZ"),
	"BS": Country("BS"),
	"BH": Country("BH"),
	"BD": Country("BD"),
	"BB": Country("BB"),
	"BY": Country("BY"),
	"BE": Country("BE"),
	"BZ": Country("BZ"),
	"BJ": Country("BJ"),
	"BM": Country("BM"),
	"BT": Country("BT"),
	"BO": Country("BO"),
	"BA": Country("BA"),
	"BW": Country("BW"),
	"BV": Country("BV"),
	"BR": Country("BR"),
	"IO": Country("IO"),
	"BN": Country("BN"),
	"BG": Country("BG"),
	"BF": Country("BF"),
	"BI": Country("BI"),
	"KH": Country("KH"),
	"CM": Country("CM"),
	"CA": Country("CA"),
	"CV": Country("CV"),
	"KY": Country("KY"),
	"CF": Country("CF"),
	"TD": Country("TD"),
	"CL": Country("CL"),
	"CN": Country("CN"),
	"CX": Country("CX"),
	"CC": Country("CC"),
	"CO": Country("CO"),
	"KM": Country("KM"),
	"CG": Country("CG"),
	"CD": Country("CD"),
	"CK": Country("CK"),
	"CR": Country("CR"),
	"CI": Country("CI"),
	"HR": Country("HR"),
	"CU": Country("CU"),
	"CY": Country("CY"),
	"CZ": Country("CZ"),
	"DK": Country("DK"),
	"DJ": Country("DJ"),
	"DM": Country("DM"),
	"DO": Country("DO"),
	"EC": Country("EC"),
	"EG": Country("EG"),
	"SV": Country("SV"),
	"GQ": Country("GQ"),
	"ER": Country("ER"),
	"EE": Country("EE"),
	"ET": Country("ET"),
	"FK": Country("FK"),
	"FO": Country("FO"),
	"FJ": Country("FJ"),
	"FI": Country("FI"),
	"FR": Country("FR"),
	"GF": Country("GF"),
	"PF": Country("PF"),
	"TF": Country("TF"),
	"GA": Country("GA"),
	"GM": Country("GM"),
	"GE": Country("GE"),
	"DE": Country("DE"),
	"GH": Country("GH"),
	"GI": Country("GI"),
	"GR": Country("GR"),
	"GL": Country("GL"),
	"GD": Country("GD"),
	"GP": Country("GP"),
	"GU": Country("GU"),
	"GT": Country("GT"),
	"GG": Country("GG"),
	"GN": Country("GN"),
	"GW": Country("GW"),
	"GY": Country("GY"),
	"HT": Country("HT"),
	"HM": Country("HM"),
	"VA": Country("VA"),
	"HN": Country("HN"),
	"HK": Country("HK"),
	"HU": Country("HU"),
	"IS": Country("IS"),
	"IN": Country("IN"),
	"ID": Country("ID"),
	"IR": Country("IR"),
	"IQ": Country("IQ"),
	"IE": Country("IE"),
	"IM": Country("IM"),
	"IL": Country("IL"),
	"IT": Country("IT"),
	"JM": Country("JM"),
	"JP": Country("JP"),
	"JE": Country("JE"),
	"JO": Country("JO"),
	"KZ": Country("KZ"),
	"KE": Country("KE"),
	"KI": Country("KI"),
	"KR": Country("KR"),
	"KW": Country("KW"),
	"KG": Country("KG"),
	"LA": Country("LA"),
	"LV": Country("LV"),
	"LB": Country("LB"),
	"LS": Country("LS"),
	"LR": Country("LR"),
	"LY": Country("LY"),
	"LI": Country("LI"),
	"LT": Country("LT"),
	"LU": Country("LU"),
	"MO": Country("MO"),
	"MK": Country("MK"),
	"MG": Country("MG"),
	"MW": Country("MW"),
	"MY": Country("MY"),
	"MV": Country("MV"),
	"ML": Country("ML"),
	"MT": Country("MT"),
	"MH": Country("MH"),
	"MQ": Country("MQ"),
	"MR": Country("MR"),
	"MU": Country("MU"),
	"YT": Country("YT"),
	"MX": Country("MX"),
	"FM": Country("FM"),
	"MD": Country("MD"),
	"MC": Country("MC"),
	"MN": Country("MN"),
	"ME": Country("ME"),
	"MS": Country("MS"),
	"MA": Country("MA"),
	"MZ": Country("MZ"),
	"MM": Country("MM"),
	"NA": Country("NA"),
	"NR": Country("NR"),
	"NP": Country("NP"),
	"NL": Country("NL"),
	"AN": Country("AN"),
	"NC": Country("NC"),
	"NZ": Country("NZ"),
	"NI": Country("NI"),
	"NE": Country("NE"),
	"NG": Country("NG"),
	"NU": Country("NU"),
	"NF": Country("NF"),
	"MP": Country("MP"),
	"NO": Country("NO"),
	"OM": Country("OM"),
	"PK": Country("PK"),
	"PW": Country("PW"),
	"PS": Country("PS"),
	"PA": Country("PA"),
	"PG": Country("PG"),
	"PY": Country("PY"),
	"PE": Country("PE"),
	"PH": Country("PH"),
	"PN": Country("PN"),
	"PL": Country("PL"),
	"PT": Country("PT"),
	"PR": Country("PR"),
	"QA": Country("QA"),
	"RE": Country("RE"),
	"RO": Country("RO"),
	"RU": Country("RU"),
	"RW": Country("RW"),
	"BL": Country("BL"),
	"SH": Country("SH"),
	"KN": Country("KN"),
	"LC": Country("LC"),
	"MF": Country("MF"),
	"PM": Country("PM"),
	"VC": Country("VC"),
	"WS": Country("WS"),
	"SM": Country("SM"),
	"ST": Country("ST"),
	"SA": Country("SA"),
	"SN": Country("SN"),
	"RS": Country("RS"),
	"SC": Country("SC"),
	"SL": Country("SL"),
	"SG": Country("SG"),
	"SK": Country("SK"),
	"SI": Country("SI"),
	"SB": Country("SB"),
	"SO": Country("SO"),
	"ZA": Country("ZA"),
	"GS": Country("GS"),
	"ES": Country("ES"),
	"LK": Country("LK"),
	"SD": Country("SD"),
	"SR": Country("SR"),
	"SJ": Country("SJ"),
	"SZ": Country("SZ"),
	"SE": Country("SE"),
	"CH": Country("CH"),
	"SY": Country("SY"),
	"TW": Country("TW"),
	"TJ": Country("TJ"),
	"TZ": Country("TZ"),
	"TH": Country("TH"),
	"TL": Country("TL"),
	"TG": Country("TG"),
	"TK": Country("TK"),
	"TO": Country("TO"),
	"TT": Country("TT"),
	"TN": Country("TN"),
	"TR": Country("TR"),
	"TM": Country("TM"),
	"TC": Country("TC"),
	"TV": Country("TV"),
	"UG": Country("UG"),
	"UA": Country("UA"),
	"AE": Country("AE"),
	"GB": Country("GB"),
	"US": Country("US"),
	"UM": Country("UM"),
	"UY": Country("UY"),
	"UZ": Country("UZ"),
	"VU": Country("VU"),
	"VE": Country("VE"),
	"VN": Country("VN"),
	"VG": Country("VG"),
	"VI": Country("VI"),
	"WF": Country("WF"),
	"EH": Country("EH"),
	"YE": Country("YE"),
	"ZM": Country("ZM"),
	"ZW": Country("ZW"),
}
