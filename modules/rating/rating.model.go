package rating

type Rating struct {
	Rating_id int `json:"rating_id" gorm:"primaryKey"`
	Hotel_id  int `json:"hotel_id" gorm:"foreignKey:hotel_id" validate:"required"`
	User_id   int `json:"user_id" gorm:"foreignKey:user_id" validate:"required"`
	Rating    int `json:"rating" gorm:"type:int" validate:"required, gte=1, lte=5"`
}

type RequestRating struct {
	Rating int `json:"rating" validate:"required, gte=1, lte=5"`
}

type ResponseRating struct {
	Hotel_id int `json:"hotel_id"`
	Rating   int `json:"rating"`
}
