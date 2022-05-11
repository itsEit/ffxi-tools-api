package models

type Stats struct {
	Def              int `json:"def,omitempty"`
	Hp               int `json:"hp,omitempty"`
	Mp               int `json:"mp,omitempty"`
	Str              int `json:"str,omitempty"`
	Dex              int `json:"dex,omitempty"`
	Vit              int `json:"vit,omitempty"`
	Agi              int `json:"agi,omitempty"`
	Int              int `json:"int,omitempty"`
	Mnd              int `json:"mnd,omitempty"`
	Chr              int `json:"chr,omitempty"`
	MagicAccuracy    int `json:"magicAcc,omitempty"`
	PhysicalAccuracy int `json:"physicalAcc,omitempty"`
	MagicEvasion     int `json:"magicEvasion,omitempty"`
	PhysicalEvasion  int `json:"physicalEvasion,omitempty"`
}
