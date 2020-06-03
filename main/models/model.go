package models

import "time"

type Gender struct {
	Gender string
}
type Profile struct {
	ID                               uint64    `gorm:"primary_key" json:"id,omitempty"`
	HotelID                          uint64    `gorm:"primary_key" json:"hotel_id,omitempty"`
	FirstName                        string    `json:"first_name,omitempty"`
	MiddleName                       string    `json:"middle_name,omitempty"`
	LastName                         string    `json:"last_name,omitempty"`
	ProfilePicURL                    string    `json:"profile_pic_url,omitempty"`
	Locale                           string    `json:"locale,omitempty"`
	Timezone                         float64   `json:"timezone,omitempty"`
	Gender                           string    `json:"gender,omitempty"`
	ClientID                         uint64    `json:"client_id,omitempty"`
	RoomNo                           string    `json:"room_no,omitempty"`
	MainMarket                       string    `json:"main_market,omitempty"` // Yeni
	Market                           string    `json:"market,omitempty"`      // Yeni
	AgentName                        string    `json:"agent_name,omitempty"`  // Yeni
	SourceName                       string    `json:"source_name,omitempty"` // Yeni
	SalesDate                        time.Time `json:"sales_date,omitempty"`  // Yeni
	EntryDate                        time.Time `json:"entry_date,omitempty"`
	ReleaseDate                      time.Time `json:"release_date,omitempty"`
	PhoneCountryID                   string    `json:"phone_country_id,omitempty"`
	PhonePrefixID                    string    `json:"phone_prefix_id,omitempty"`
	PhoneNo                          string    `json:"phone_no,omitempty"`
	MobileCountryID                  string    `json:"mobile_country_id,omitempty"`
	MobilePrefixID                   string    `json:"mobile_prefix_id,omitempty"`
	MobileNo                         string    `json:"mobile_no,omitempty"`
	Address                          string    `json:"address,omitempty"`
	City                             string    `json:"city,omitempty"`
	CountyCode                       string    `json:"county_code,omitempty"` // Yeni
	Country                          string    `json:"country,omitempty"`
	VoucherOrBookingNo               string    `json:"voucher_or_booking_no,omitempty"`
	BookingDate                      time.Time `json:"booking_date,omitempty"`
	DateOfBirth                      time.Time `json:"date_of_birth,omitempty"`
	BirthPlace                       string    `json:"birth_place,omitempty"`
	Nationality                      string    `json:"nationality,omitempty"`
	TravelDocType                    string    `json:"travel_doc_type,omitempty"`
	TravelDocID                      string    `json:"travel_doc_id,omitempty"`
	DocumentIssuedOn                 time.Time `json:"document_issued_on,omitempty"`
	DocumentIssuedIn                 string    `json:"document_issued_in,omitempty"`
	EMail                            string    `json:"e_mail,omitempty"`
	LiveSupportIsActiveUntilThisTime time.Time `json:"live_support_is_active_until_this_time"`
	GetExternalInput                 bool      `json:"get_external_input,omitempty"`
	ActivePayload                    string    `json:"active_payload,omitempty"`
	NextPayload                      string    `json:"next_payload,omitempty"`
	CheckRoomNo                      bool      `json:"check_room_no"`
	NextPayloadAfterCheckRoom        string    `json:"next_payload_after_check_room"`
	CreatedAt                        time.Time `json:"created_at,omitempty"`
	UpdatedAt                        time.Time `json:"updated_at,omitempty"`
	Password                         string    `json:"password,omitempty"`
	Verify                           bool      `json:"verify,omitempty"`
	PrivacyPolicy                    bool      `json:"privacy_policy"`
	PersonalData                     bool      `json:"personal_data"`
}
