package user

import "time"

// User represents a user account
type User struct {
	// Uid is the User id
	ID     int
	Name   string
	Date   time.Time
	Status bool
}

func (this *User) SetupUser(id int, name string, date time.Time, status bool) {
	this.ID = id
	this.Name = name
	this.Date = date
	this.Status = status
}
