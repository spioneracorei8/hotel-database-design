package constants

// # hotel_status -> Hotel
type HotelStatus string

const (
	HOTEL_OPEN   HotelStatus = "OPEN"
	HOTEL_CLOSED HotelStatus = "CLOSED"

	HOTEL_OPEN_TH   HotelStatus = "เปิด"
	HOTEL_CLOSED_TH HotelStatus = "ปิด"
)

// # hotel_rooms_status -> HotelRooms
type HotelRoomsStatus string

const (
	HOTEL_ROOM_BOOK  HotelRoomsStatus = "BOOK"
	HOTEL_ROOM_BLANK HotelRoomsStatus = "BLANK"

	HOTEL_ROOM_BOOK_TH  HotelRoomsStatus = "จอง"
	HOTEL_ROOM_BLANK_TH HotelRoomsStatus = "ว่าง"
)

// # hotel_room_type_status -> HotelRoomType
type HotelRoomTypeStatus string

const (
	HOTEL_ROOM_TYPE_INACTIVE    HotelRoomTypeStatus = "INACTIVE"
	HOTEL_ROOM_TYPE_ACTIVE      HotelRoomTypeStatus = "ACTIVE"
	HOTEL_ROOM_TYPE_MAINTENANCE HotelRoomTypeStatus = "MAINTENANCE"

	HOTEL_ROOM_TYPE_INACTIVE_TH    HotelRoomTypeStatus = "ไม่ได้ใช้งาน"
	HOTEL_ROOM_TYPE_ACTIVE_TH      HotelRoomTypeStatus = "ใช้งาน"
	HOTEL_ROOM_TYPE_MAINTENANCE_TH HotelRoomTypeStatus = "ซ่อมบำรุง"
)

// # bookings_status -> Bookings
type BookingsStatus string

const (
	BOOKINGS_CHECK_IN  BookingsStatus = "CHECK-IN"
	BOOKINGS_CHECK_OUT BookingsStatus = "CHECK-OUT"
	BOOKINGS_BLANK     BookingsStatus = "MAINTENANCE"

	BOOKINGS_CHECK_IN_TH  BookingsStatus = "เช็คอิน"
	BOOKINGS_CHECK_OUT_TH BookingsStatus = "เช็คเอาท์"
	BOOKINGS_BLANK_TH     BookingsStatus = "ว่าง"
)
