package note

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	notes []Note
}

func NewService() *Service {
	return &Service{make([]Note, 0)}
}

func (s *Service) Create(title, description string) (Note, error) {
	// Validating the title
	if title == "" {
		return Note{}, fmt.Errorf("Create(): Title cannot be empty.")
	}

	// validating the description
	if description == "" {
		return Note{}, fmt.Errorf("Create(): Description cannot be empty")
	}

	// Creating a new id for note
	id := uuid.NewString()

	// Creating the note
	newNote := Note{ID: id, Title: title, Description: description, Time: time.Now()}

	// Appending the note into notes
	s.notes = append(s.notes, newNote)

	return newNote, nil

}


// GetAllNotes returns copy of all the notes 
func (s *Service) GetAllNotes() []Note {
	result := make([]Note, 0, len(s.notes))

	for _, note := range s.notes {
		result = append(result, note)
	}
	return result
}
