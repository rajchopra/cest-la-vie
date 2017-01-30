package main

import "fmt"
import "time"

var countryTz = map[string]string{
	"Hungary":       "Europe/Budapest",
	"Egypt":         "Africa/Cairo",
	"Aden":          "Asia/Aden",
	"Almaty":        "Asia/Almaty",
	"Amman":         "Asia/Amman",
	"Anadyr":        "Asia/Anadyr",
	"Aqtau":         "Asia/Aqtau",
	"Aqtobe":        "Asia/Aqtobe",
	"Ashgabat":      "Asia/Ashgabat",
	"Atyrau":        "Asia/Atyrau",
	"Baghdad":       "Asia/Baghdad",
	"Bahrain":       "Asia/Bahrain",
	"Baku":          "Asia/Baku",
	"Bangkok":       "Asia/Bangkok",
	"Barnaul":       "Asia/Barnaul",
	"Beirut":        "Asia/Beirut",
	"Bishkek":       "Asia/Bishkek",
	"Brunei":        "Asia/Brunei",
	"Chita":         "Asia/Chita",
	"Choibalsan":    "Asia/Choibalsan",
	"Colombo":       "Asia/Colombo",
	"Damascus":      "Asia/Damascus",
	"Dhaka":         "Asia/Dhaka",
	"Dili":          "Asia/Dili",
	"Dubai":         "Asia/Dubai",
	"Dushanbe":      "Asia/Dushanbe",
	"Famagusta":     "Asia/Famagusta",
	"Gaza":          "Asia/Gaza",
	"Hebron":        "Asia/Hebron",
	"Ho_Chi_Minh":   "Asia/Ho_Chi_Minh",
	"Hong_Kong":     "Asia/Hong_Kong",
	"Hovd":          "Asia/Hovd",
	"Irkutsk":       "Asia/Irkutsk",
	"Jakarta":       "Asia/Jakarta",
	"Jayapura":      "Asia/Jayapura",
	"Jerusalem":     "Asia/Jerusalem",
	"Kabul":         "Asia/Kabul",
	"Kamchatka":     "Asia/Kamchatka",
	"Karachi":       "Asia/Karachi",
	"Kathmandu":     "Asia/Kathmandu",
	"Khandyga":      "Asia/Khandyga",
	"Kolkata":       "Asia/Kolkata",
	"Krasnoyarsk":   "Asia/Krasnoyarsk",
	"Kuala_Lumpur":  "Asia/Kuala_Lumpur",
	"Kuching":       "Asia/Kuching",
	"Kuwait":        "Asia/Kuwait",
	"Macau":         "Asia/Macau",
	"Magadan":       "Asia/Magadan",
	"Makassar":      "Asia/Makassar",
	"Manila":        "Asia/Manila",
	"Muscat":        "Asia/Muscat",
	"Nicosia":       "Asia/Nicosia",
	"Novokuznetsk":  "Asia/Novokuznetsk",
	"Novosibirsk":   "Asia/Novosibirsk",
	"Omsk":          "Asia/Omsk",
	"Oral":          "Asia/Oral",
	"Phnom_Penh":    "Asia/Phnom_Penh",
	"Pontianak":     "Asia/Pontianak",
	"Pyongyang":     "Asia/Pyongyang",
	"Qatar":         "Asia/Qatar",
	"Qyzylorda":     "Asia/Qyzylorda",
	"Riyadh":        "Asia/Riyadh",
	"Sakhalin":      "Asia/Sakhalin",
	"Samarkand":     "Asia/Samarkand",
	"Seoul":         "Asia/Seoul",
	"Shanghai":      "Asia/Shanghai",
	"Singapore":     "Asia/Singapore",
	"Srednekolymsk": "Asia/Srednekolymsk",
	"Taipei":        "Asia/Taipei",
	"Tashkent":      "Asia/Tashkent",
	"Tbilisi":       "Asia/Tbilisi",
	"Tehran":        "Asia/Tehran",
	"Thimphu":       "Asia/Thimphu",
	"Tokyo":         "Asia/Tokyo",
	"Tomsk":         "Asia/Tomsk",
	"Ulaanbaatar":   "Asia/Ulaanbaatar",
	"Urumqi":        "Asia/Urumqi",
	"Nera":          "Ust-Nera",
	"Vientiane":     "Asia/Vientiane",
	"Vladivostok":   "Asia/Vladivostok",
	"Yakutsk":       "Asia/Yakutsk",
	"Yangon":        "Asia/Yangon",
	"Yekaterinburg": "Asia/Yekaterinburg",
	"Yerevan":       "Asia/Yerevan",
}

func timeIn(name string) time.Time {
	loc, err := time.LoadLocation(countryTz[name])
	if err != nil {
		panic(err)
	}
	return time.Now().In(loc)
}

func main() {
	utc := time.Now().UTC().Format("15:04")
	hun := timeIn("Hungary").Format("15:04")
	eg := timeIn("Egypt").Format("15:04")
	ade := timeIn("Aden").Format("22:11")
	alm := timeIn("Almaty").Format("22:11")
	amm := timeIn("Amman").Format("22:11")
	ana := timeIn("Anadyr").Format("22:11")
	aqt := timeIn("Aqtau").Format("22:11")
	aqt := timeIn("Aqtobe").Format("22:11")
	ash := timeIn("Ashgabat").Format("22:11")
	bag := timeIn("Baghdad").Format("22:11")
	bah := timeIn("Bahrain").Format("22:11")
	bak := timeIn("Baku").Format("22:11")
	bar := timeIn("Barnaul").Format("22:11")
	bru := timeIn("Brunei").Format("22:11")
	cho := timeIn("Choibalsan").Format("22:11")
	pon := timeIn("Pontianak").Format("22:11")
	sak := timeIn("Sakhalin").Format("22:11")
	sam := timeIn("Samarkand").Format("22:11")
	sin := timeIn("Singapore").Format("22:11")
	sre := timeIn("Srednekolymsk").Format("22:11")
	tas := timeIn("Tashkent").Format("22:11")
	tbi := timeIn("Tbilisi").Format("22:11")
	teh := timeIn("Tehran").Format("22:11")
	tok := timeIn("Tokyo").Format("22:11")
	ula := timeIn("Ulaanbaatar").Format("22:11")
	uru := timeIn("Urumqi").Format("22:11")
	ner := timeIn("Nera").Format("22:11")
	vie := timeIn("Vientiane").Format("22:11")
	vla := timeIn("Vladivostok").Format("22:11")
	yan := timeIn("Yangon").Format("22:11")
	yek := timeIn("Yekaterinburg").Format("22:11")
	yer := timeIn("Yerevan").Format("22:11")
	fmt.Println(utc, hun, eg)
}
