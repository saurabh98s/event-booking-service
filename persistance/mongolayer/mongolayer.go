package persistance

import (
	"github.com/geeks/cloud-native/persistance"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Event struct {
	ID        bson.ObjectId `bson:"_id"`
	Name      string
	Duration  int
	StartDate int64
	EndDate   int64
	Location  Location
}

type Location struct {
	Name      string
	Address   string
	Country   string
	OpenTime  int
	CloseTime int
	Halls     []Hall
}
type Hall struct {
	Name     string `json:"name"`
	Location string `json:"location,omitempty"`
	Capacity int    `json:"capacity"`
}

const (
	DB     = "myevents"
	USERS  = "users"
	EVENTS = "events"
)

type DatabaseHandler interface {
	AddEvent(Event) ([]byte, error)
	FindEvent([]byte) (Event, error)
	FindEventByName(string) (Event, error)
	FindAllAvailableEvents() ([]Event, error)
}

type MongoDBLayer struct {
	session *mgo.Session
}

func NewMongoDBLayer(connection string) (*MongoDBLayer, error) {
	s, err := mgo.Dial(connection)
	if err != nil {
		return nil, err
	}
	return &MongoDBLayer{
		session: s,
	}, err
}

func (mgoLayer *MongoDBLayer) getFreshSession() *mgo.Session {
	return mgoLayer.session.Copy()
}

func (mgoLayer *MongoDBLayer) AddEvent(e persistance.Event) ([]byte, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	if !e.ID.Valid() {

		e.ID = bson.NewObjectId()
	} //let's assume the method below checks if the ID is valid for the location object of the event
	if !e.Location.ID.Valid() {
		e.Location.ID = bson.NewObjectId()

	}
	return []byte(e.ID), s.DB(DB).C(EVENTS).Insert(e)
}

func (mgoLayer *MongoDBLayer) FindEventByName(name string) (persistance.Event, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	e := persistance.Event{}
	err := s.DB(DB).C(EVENTS).Find(bson.M{"name": name}).One(&e)

	return e, err
}

func (mgoLayer *MongoDBLayer) FindAllAvailableEvents() ([]persistance.Event, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	events := []persistance.Event{}
	err := s.DB(DB).C(EVENTS).Find(nil).All(&events)
	return events, err
}

func (mgoLayer *MongoDBLayer) FindEvent(id []byte) (persistance.Event, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	e := persistance.Event{}
	err := s.DB(DB).C(EVENTS).FindId(bson.ObjectId(id)).One(&e)
	return e, err
}
