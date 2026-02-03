package note

import "time"

type Note struct {
	ID	  string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Time time.Time `json:"time"`
}