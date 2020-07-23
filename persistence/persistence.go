package persistence

type DatabaseHandler interface {
	AddEvent(Event) (Event, error)
	FindEvent([]byte) (Event, error)
	FindEventByName(string) (Event, error)
	FindAllAvailableEvents() ([]Event, error)
}
