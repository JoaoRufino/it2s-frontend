package models

type DENM struct {
	ID        int64 `json:"ID"`
	Lat       int64 `json:"Lat"`
	Lon       int64 `json:"Lon"`
	Timestamp int64 `json:"Timestamp"`
}

func NewDENM(ID, Lat, Long, Timestamp int64) *DENM {
	return &DENM{ID, Lat, Long, Timestamp}
}
