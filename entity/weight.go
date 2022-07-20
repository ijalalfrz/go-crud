package entity

// Weight is an entity to represent weight collection
type Weight struct {
	Date int64 `json:"date"`
	Max  int   `json:"max"`
	Min  int   `json:"min"`
	Diff int   `json:"diff"`
}
