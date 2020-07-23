package mongolayer

import (
	"cloud-native/persistence"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

const (
	DB = "myevents" //nolint:gofmt
	USERS = "users"
	EVENTS = "events"
)

type MongoDBLayer struct {
	session *mgo.Session
}

//NewMongoDBLayer creates a new layer for the mongolayer
func NewMongoDBLayer(connection string) (persistence.DatabaseHandler,error) {
	s,err:=mgo.Dial(connection)
	return &MongoDBLayer{
		session: s,
	},err
}

func (mgoLayer *MongoDBLayer) AddEvent (e persistence.Event) ( persistence.Event,error){
	s:=mgoLayer.getFreshSession()
	defer s.Close()
	e.ID=bson.NewObjectId()
	e.Location.ID=bson.NewObjectId()
	err:=s.DB(DB).C(EVENTS).Insert(e)
	if err != nil {
		return e,err
	}
	return e,nil
}

func (mgoLayer *MongoDBLayer) FindEvent (id []byte) (persistence.Event,error) {
	s:=mgoLayer.getFreshSession()
	defer s.Close()
	e:=persistence.Event{}

	err:=s.DB(DB).C(EVENTS).FindId(bson.ObjectId(id)).One(&e)
	return e,err
}

func(mgoLayer *MongoDBLayer) FindEventByName (name string) (persistence.Event,error) {
	s:=mgoLayer.getFreshSession()
	defer s.Close()
	e:=persistence.Event{}
	err:=s.DB(DB).C(EVENTS).FindId(bson.ObjectId(name)).One(&e)
	return e,err
}

func (mgoLayer *MongoDBLayer) FindAllAvailableEvents() ([]persistence.Event, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	var events []persistence.Event
	err := s.DB(DB).C(EVENTS).Find(nil).All(&events)
	return events, err
}


func (mgoLayer *MongoDBLayer) getFreshSession() *mgo.Session {
	return mgoLayer.session.Copy()
}
