package models

type Item struct {
	Id       int32  `json:"_id,omitempty"`
	Name     string `json:"name,omitempty"`
	NameFull string `json:"nameFull,omitempty"`
	Category string `json:"category,omitempty"`
	Desc     string `json:"desc,omitempty"`
	SlotName string `json:"slotName,omitempty"`
	Stats    Stats  `json:"stats,omitempty"`
}
