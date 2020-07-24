package persistence

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
)

type User struct {
	ID       bson.ObjectId `bson:"_id"`
	First    string        `json:"first"`
	Last     string        `json:"last"`
	Age      int           `json:"age"`
	Bookings []Booking     `json:"bookings"`
}

func (u *User) String() string {
	return fmt.Sprintf("id: %s, first_name: %s, last_name: %s, Age: %d, Bookings: %v", u.ID, u.First, u.Last, u.Age, u.Bookings)
}

type Booking struct {
	Date    int64  `json:"date"`
	EventID []byte `json:"event_id"`
	Seats   int    `json:"seats"`
}

type Event struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Name      string        `json:"name"`
	Duration  int           `json:"duration"`
	StartDate int64         `json:"start_date"`
	EndDate   int64         `json:"end_date"`
	Location  Location      `json:"location"`
}

type Location struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Name      string        `json:"name"`
	Address   string        `json:"address"`
	Country   string        `json:"country"`
	OpenTime  int           `json:"open_time"`
	CloseTime int           `json:"close_time"`
	Halls     []Hall        `json:"halls"`
}

type Hall struct {
	Name     string `json:"name"`
	Location string `json:"location,omitempty"`
	Capacity int    `json:"capacity"`
}
