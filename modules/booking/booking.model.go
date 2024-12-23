package booking

type Booking struct {
	Booking_id       int
	User_id          int
	Kamar_id         int
	Tanggal_check_in  string
	Tanggal_check_out string
	Total_biaya      int
	Status_booking   string
}