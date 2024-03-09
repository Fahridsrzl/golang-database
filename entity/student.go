package entity

import "time"

type Student struct {
	Id        int
	Name      string
	Email     string
	Address   string
	BirthDate time.Time
	Gender    string
}
