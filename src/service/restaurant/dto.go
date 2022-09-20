package restaurant

import "time"

type RestaurantDto struct {
	OwnerId          uint64                 `json:"owner_id"`
	Name             string                 `json:"name"`
	Address          string                 `json:"address"`
	CityId           uint64                 `json:"city_id"`
	Lat              float64                `json:"lat"`
	Lng              float64                `json:"lng"`
	Logo             map[string]interface{} `json:"logo"`
	ShippingFeePerKm float64                `json:"shipping_fee_per_km"`
	Status           int8                   `json:"status"`
	CreatedAt        time.Time              `json:"created_at"`
	UpdatedAt        time.Time              `json:"updated_at"`
}

type FullResponseObject struct {
	RestaurantDto
}

type AllRestaurantDto struct {
	Restaurants []*RestaurantDto
}
