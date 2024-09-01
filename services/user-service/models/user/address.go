package user

import "time"

type (
	AddressEntity struct {
		UserAddressId string    `json:"user_address_id" db:"user_address_id"`
		UserId        string    `json:"user_id" db:"user_id"`
		ProvinceId    string    `json:"province_id" db:"province_id"`
		CityId        string    `json:"city_id" db:"city_id"`
		DistrictId    string    `json:"district_id" db:"district_id"`
		SubDistrictId string    `json:"sub_district_id" db:"sub_district_id"`
		Zipcode       string    `json:"province" db:"zipcode"`
		Address       string    `json:"address" db:"address"`
		Note          *string   `json:"note" db:"note"`
		GoogleMap     *string   `json:"google_map" db:"google_map"`
		IsMain        bool      `json:"is_main" db:"is_main"`
		CreatedAt     time.Time `json:"created_at" db:"created_at"`
	}
)
