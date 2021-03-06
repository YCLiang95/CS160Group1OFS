package protocal

type Merchandise struct {
	Mid      int64   `xorm:"pk 'mid'" json:"mid"`
	Name     string  `json:"name"`
	Weight   float64 `json:"weight"`
	Price    float64 `json:"price"`
	Quantity int64   `json:"quantify"`
	ImageUrl string  `json:"imagerl"`
}


type MerchandisRequest struct {
	Name     string  `json:"name"`
	Weight   float64 `json:"weight"`
	Price    float64 `json:"price"`
	Quantity int64   `json:"quantify"`
	ImageUrl string  `json:"imagerl"`
}

type MerchandisUpdateRequest struct {
	Mid      int64   `json:"mid"`
	Name     string  `json:"name"`
	Weight   float64 `json:"weight"`
	Price    float64 `json:"price"`
	Quantity int64   `json:"quantify"`
	ImageUrl string  `json:"imagerl"`
}