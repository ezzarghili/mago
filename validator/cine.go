package validator

import (
	"regexp"
	"strings"
)

// CINE CINE ID validation
type CINE struct {
	Location string // Provinces and prefectures ISO code
	Spacer   string // regex expression
}

// Provinces and prefectures
/*
AGD	Agadir-Ida-Ou-Tanane
AOU	Aousserd
ASZ	Assa-Zag
AZI	Azilal
BEM	Béni Mellal
BES	Benslimane
BER	Berkane
BRR	Berrechid
BOD	Boujdour
BOM	Boulemane
CAS	Casablanca
CHE	Chefchaouen
CHI	Chichaoua
CHT	Chtouka-Ait Baha
DRI	Driouch
ERR	Errachidia
ESM	Es-Semara
ESI	Essaouira
FAH	Fahs-Anjra
FES	Fès
FIG	Figuig
FQH	Fquih Ben Salah
GUE	Guelmim
GUF	Guercif
HAJ	El Hajeb
HAO	Al Haouz
HOC	Al Hoceïma
IFR	Ifrane
INE	Inezgane-Ait Melloul
JDI	El Jadida
JRA	Jerada
KES	El Kelâa des Sraghna
KEN	Kénitra
KHE	Khemisset
KHN	Khenifra
KHO	Khouribga
LAA	Laâyoune (EH)
LAR	Larache
MAR	Marrakech
MDF	M’diq-Fnideq
MED	Médiouna
MEK	Meknès
MID	Midelt
MOH	Mohammadia
MOU	Moulay Yacoub
NAD	Nador
NOU	Nouaceur
OUA	Ouarzazate
OUD	Oued Ed-Dahab
OUZ	Ouezzane
OUJ	Oujda-Angad
RAB	Rabat
REH	Rehamna
SAF	Safi
SAL	Salé
SEF	Sefrou
SET	Settat
SIB	Sidi Bennour
SIF	Sidi Ifni
SIK	Sidi Kacem
SIL	Sidi Slimane
SKH	Skhirate-Témara
TNT	Tan-Tan
TNG	Tanger-Assilah
TAO	Taounate
TAI	Taourirt
TAF	Tarfaya
TAR	Taroudant
TAT	Tata
TAZ	Taza
TET	Tétouan
TIN	Tinghir
TIZ	Tiznit
YUS	Youssoufia
ZAG	Zagora
*/

var cityMap = map[string][]string{
	"bes": []string{
		"ta",
		"tk",
	},
}

// Validate if CINE code is correct
func (validator *CINE) Validate(cine string) bool {
	var spacerReg string
	if validator.Spacer != "" {
		spacerReg = validator.Spacer
	} else {
		spacerReg = spacers
	}
	value := strings.ToLower(clean(cine, spacerReg))
	if validator.Location == "" {
		reg := regexp.MustCompile(`^[a-z]{1,2}\d+$`)
		return reg.MatchString(value)
	}
	combined := strings.Join(cityMap[strings.ToLower(validator.Location)], "|") // bes-> ta|tk
	if combined != "" {
		reg := regexp.MustCompile(`^(` + combined + `)\d+$`)
		return reg.MatchString(value)
	}
	return false
}
