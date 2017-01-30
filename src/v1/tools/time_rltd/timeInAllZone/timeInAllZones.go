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
	utc := fmt.Sprintf("UTC :: %s", time.Now().UTC().Format("15:04:05"))
	hun := fmt.Sprintf("Hungary :: %s", timeIn("Hungary").Format("15:04:05"))
	eg := fmt.Sprintf("Egypt :: %s", timeIn("Egypt").Format("15:04:05"))
	ade := fmt.Sprintf("Aden :: %s", timeIn("Aden").Format("15:04:05"))
	alm := fmt.Sprintf("Almaty :: %s", timeIn("Almaty").Format("15:04:05"))
	amm := fmt.Sprintf("Amman :: %s", timeIn("Amman").Format("15:04:05"))
	ana := fmt.Sprintf("Anadyr :: %s", timeIn("Anadyr").Format("15:04:05"))
	aqt := fmt.Sprintf("Aqtobe :: %s", timeIn("Aqtobe").Format("15:04:05"))
	ash := fmt.Sprintf("Ashgabat :: %s", timeIn("Ashgabat").Format("15:04:05"))
	bag := fmt.Sprintf("Baghdad :: %s", timeIn("Baghdad").Format("15:04:05"))
	bah := fmt.Sprintf("Bahrain :: %s", timeIn("Bahrain").Format("15:04:05"))
	bak := fmt.Sprintf("Baku :: %s", timeIn("Baku").Format("15:04:05"))
	bar := fmt.Sprintf("Barnaul :: %s", timeIn("Barnaul").Format("15:04:05"))
	bru := fmt.Sprintf("Brunei :: %s", timeIn("Brunei").Format("15:04:05"))
	cho := fmt.Sprintf("Choibalsan :: %s", timeIn("Choibalsan").Format("15:04:05"))
	kol := fmt.Sprintf("Kolkata :: %s", timeIn("Kolkata").Format("15:04:05"))
	pon := fmt.Sprintf("Pontianak :: %s", timeIn("Pontianak").Format("15:04:05"))
	sak := fmt.Sprintf("Sakhalin :: %s", timeIn("Sakhalin").Format("15:04:05"))
	sam := fmt.Sprintf("Samarkand :: %s", timeIn("Samarkand").Format("15:04:05"))
	sin := fmt.Sprintf("Singapore :: %s", timeIn("Singapore").Format("15:04:05"))
	sre := fmt.Sprintf("Srednekolymsk :: %s", timeIn("Srednekolymsk").Format("15:04:05"))
	tas := fmt.Sprintf("Tashkent :: %s", timeIn("Tashkent").Format("15:04:05"))
	tbi := fmt.Sprintf("Tbilisi :: %s", timeIn("Tbilisi").Format("15:04:05"))
	teh := fmt.Sprintf("Tehran :: %s", timeIn("Tehran").Format("15:04:05"))
	tok := fmt.Sprintf("Tokyo :: %s", timeIn("Tokyo").Format("15:04:05"))
	ula := fmt.Sprintf("Ulaanbaatar :: %s", timeIn("Ulaanbaatar").Format("15:04:05"))
	uru := fmt.Sprintf("Urumqi :: %s", timeIn("Urumqi").Format("15:04:05"))
	vie := fmt.Sprintf("Vientiane :: %s", timeIn("Vientiane").Format("15:04:05"))
	vla := fmt.Sprintf("Vladivostok :: %s", timeIn("Vladivostok").Format("15:04:05"))
	yan := fmt.Sprintf("Yangon :: %s", timeIn("Yangon").Format("15:04:05"))
	yek := fmt.Sprintf("Yekaterinburg :: %s", timeIn("Yekaterinburg").Format("15:04:05"))
	yer := fmt.Sprintf("Yerevan :: %s", timeIn("Yerevan").Format("15:04:05"))
	fmt.Println("\n", utc, "\n", kol)
	fmt.Println("\n", ade, "\n", alm, "\n", amm, "\n", ana, "\n", aqt, "\n", aqt, "\n", ash, "\n", bag, "\n", bah, "\n", bak, "\n", bar, "\n", bru, "\n", cho, "\n", eg, "\n", hun, "\n", pon, "\n", sak, "\n", sam, "\n", sin, "\n", sre, "\n", tas, "\n", tbi, "\n", teh, "\n", tok, "\n", ula, "\n", uru, "\n", vie, "\n", vla, "\n", yan, "\n", yek, "\n", yer)
}
