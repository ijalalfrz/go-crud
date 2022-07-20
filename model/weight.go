package model

// WeightPayload is a model for weight http request
type WeightPayload struct {
	Date int64 `json:"date" validate:"required"`
	Max  int   `json:"max" validate:"required"`
	Min  int   `json:"min" validate:"required"`
}

type WeightResponse struct {
	List        []WeighDetailResponse `json:"list"`
	AverageMax  float32               `json:"averageMax"`
	AverageMin  float32               `json:"averageMin"`
	AverageDiff float32               `json:"averageDiff"`
}

type WeighDetailResponse struct {
	Date       int64  `json:"date"`
	DateString string `json:"-"`
	Max        int    `json:"max"`
	Min        int    `json:"min"`
	Diff       int    `json:"diff"`
}
