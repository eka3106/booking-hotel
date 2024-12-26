package fasilitashotel

import (
	"booking-hotel/databases"
	"booking-hotel/libs"
	"booking-hotel/modules/fasilitas"
	"booking-hotel/modules/hotel"
	"booking-hotel/modules/user"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func CreateFasilitasHotel(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	fasilitasHotel := FasilitasHotel{}
	if err := c.BodyParser(&fasilitasHotel); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	err := validate.Struct(fasilitasHotel)
	if err != nil {
		err := err.(validator.ValidationErrors)
		errors := map[string]string{}
		for _, err := range err {
			errors[err.Field()] = err.Tag()
		}
		return libs.ResponseError(c, errors, 400)
	}
	checkHotel := make(chan bool)
	checkFasilitas := make(chan bool)
	checkInFasilitasHotel := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(1)
    go checkFasilitasHotelIsExist(fasilitasHotel.Hotel_id, fasilitasHotel.Fasilitas_id, &wg, checkInFasilitasHotel)

    wg.Add(1)
    go checkIsFasilitasExist(fasilitasHotel.Fasilitas_id, &wg, checkFasilitas)

    wg.Add(1)
    go checkIsHotelExist(fasilitasHotel.Hotel_id, &wg, checkHotel)

	go func() {
        wg.Wait()
        close(checkHotel)
        close(checkFasilitas)
        close(checkInFasilitasHotel)
    }()

    hotelExists := <-checkHotel
    fasilitasExists := <-checkFasilitas
    inFasilitasHotel := <-checkInFasilitasHotel
	
	if !hotelExists && !fasilitasExists  {
		return libs.ResponseError(c, "Hotel and Fasilitas not found", 404)
	}
    if !hotelExists {
        return libs.ResponseError(c, "Hotel not found", 404)
    }
    if !fasilitasExists {
        return libs.ResponseError(c, "Fasilitas not found", 404)
    }
    if inFasilitasHotel {
        return libs.ResponseError(c, "Fasilitas hotel already exist", 400)
    }
	
	if err := databases.DB.Table("fasilitas_hotel").Create(&fasilitasHotel); err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	return libs.ResponseSuccess(c, "Success create fasilitas hotel", 201)
}

func checkFasilitasHotelIsExist(hotel_id int, fasilitas_id int, wg *sync.WaitGroup,ch chan bool)  {
	var fasilitasHotel []FasilitasHotel
	defer wg.Done()
	if err := databases.DB.Table("fasilitas_hotel").Where("hotel_id = ? AND fasilitas_id = ?", hotel_id, fasilitas_id).Find(&fasilitasHotel); err.Error != nil {
		ch <- false		
		return 
	} else if len(fasilitasHotel) != 0 {		
		ch <- true
		return
	} 
	
}

func checkIsHotelExist(hotel_id int, wg *sync.WaitGroup, ch chan bool) {
	var hotel hotel.Hotel
	defer wg.Done()
	if err := databases.DB.Table("hotel").First(&hotel, hotel_id); err.Error != nil {
		ch <- false
		return
	} else if err.RowsAffected == 0 {
		ch <- false
		return
	} else {
		ch <- true
		return
	}	
}

func checkIsFasilitasExist(fasilitas_id int, wg *sync.WaitGroup, ch chan bool) {
	var fasilitas fasilitas.Fasilitas
	defer wg.Done()
	if err := databases.DB.Table("fasilitas").First(&fasilitas, fasilitas_id); err.Error != nil {
		ch <- false
		return
	} else if err.RowsAffected == 0 {
		ch <- false
		return
	} else {
		ch <- true
		return
	}
}

func GetAllFasilitasHotel(c *fiber.Ctx) error {
	var fasilitasHotel []FasilitasHotel
	if err := databases.DB.Table("fasilitas_hotel").Find(&fasilitasHotel).Error; err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	return libs.ResponseSuccess(c, fasilitasHotel, 200)
}

func GetFasilitasHotelById(c *fiber.Ctx) error {
	fasilitasHotel := FasilitasHotel{}
	id := c.Params("id")
	err := databases.DB.Table("fasilitas_hotel").First(&fasilitasHotel, id)
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	return libs.ResponseSuccess(c, fasilitasHotel, 200)
}

func UpdateFasilitasHotel(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	fasilitasHotel := FasilitasHotel{}
	if err := c.BodyParser(&fasilitasHotel); err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	id := c.Params("id")

	checkHotel := make(chan bool)
	checkFasilitas := make(chan bool)
	checkInFasilitasHotel := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(1)
    go checkFasilitasHotelIsExist(fasilitasHotel.Hotel_id, fasilitasHotel.Fasilitas_id, &wg, checkInFasilitasHotel)

    wg.Add(1)
    go checkIsFasilitasExist(fasilitasHotel.Fasilitas_id, &wg, checkFasilitas)

    wg.Add(1)
    go checkIsHotelExist(fasilitasHotel.Hotel_id, &wg, checkHotel)

	go func() {
        wg.Wait()
        close(checkHotel)
        close(checkFasilitas)
        close(checkInFasilitasHotel)
    }()

    hotelExists := <-checkHotel
    fasilitasExists := <-checkFasilitas
    inFasilitasHotel := <-checkInFasilitasHotel

	if !hotelExists && !fasilitasExists  {
		return libs.ResponseError(c, "Hotel and Fasilitas not found", 404)
	}
    if !hotelExists {
        return libs.ResponseError(c, "Hotel not found", 404)
    }
    if !fasilitasExists {
        return libs.ResponseError(c, "Fasilitas not found", 404)
    }
    if inFasilitasHotel {
        return libs.ResponseError(c, "Fasilitas hotel already exist", 400)
    }

	err := databases.DB.Table("fasilitas_hotel").Where("fasilitas_hotel_id = ?", id).Updates(&fasilitasHotel)
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	return libs.ResponseSuccess(c, "Success update fasilitas hotel", 200)
}

func DeleteFasilitasHotel(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	id := c.Params("id")
	err := databases.DB.Table("fasilitas_hotel").Where("fasilitas_hotel_id = ?", id).Delete(&FasilitasHotel{})
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	return libs.ResponseSuccess(c, "Success delete fasilitas hotel", 200)
}
