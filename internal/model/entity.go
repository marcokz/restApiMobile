package model

type Phone struct {
	Id              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	ModelYear       int     `json:"modelYear"`
	Diagonal        float64 `json:"diagonal"`
	MemoryStorage   int     `json:"memoryStorage"`
	Ram             int     `json:"ram"`
	Weight          int     `json:"weight"`
	BatteryCapacity int     `json:"batteryCapacity"`
	Color           string  `json:"color"`
	Price           int     `json:"price"`
}
