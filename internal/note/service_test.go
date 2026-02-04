package note

import (
	"testing"
)
// Testing create service 
func TestCreate_Service(t *testing.T) {
	tests := []struct {
		name        string
		title       string
		description string
		wantErr     bool
	}{
		{
			name:        "Empty title",
			title:       "",
			description: "Valid description",
			wantErr:     true,
		}, {
			name:        "Empty description",
			title:       "Valid title",
			description: "",
			wantErr:     true,
		}, {
			name:        "Valid note",
			title:       "Valid title",
			description: "Valid description",
			wantErr:     false,
		},
	}

	service := NewService()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			note, err := service.Create(tt.title, tt.description)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("Expected error but got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("Unexpected error %q", err)
				return
			}

			if tt.title != note.Title {
				t.Errorf("Expected title %q but got %q", tt.title, note.Title)
			}

			if tt.description != note.Description {
				t.Errorf("Expected description %q but got %q", tt.description, note.Description)
			}

			if note.ID == "" {
				t.Errorf("Expected id to be set")
			}
		})
	}
}


// Tests that internal state of serive is immutable
func TestGetAllNotes_Service(t *testing.T){

}