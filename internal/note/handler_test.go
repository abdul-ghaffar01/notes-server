package note

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// Creating a new gin engine
func setupRouter(h *Handler) *gin.Engine {
	router := gin.New()
	router.POST("/create", h.Create)
	router.GET("/notes", h.GetAllNotes)
	return router
}

func TestCreate_Handler(t *testing.T) {
	service := NewService()
	handler := NewHandler(service)
	router := setupRouter(handler)

	// defining all the tests
	tests := []struct {
		name      string
		payload   map[string]string
		wantCode  int
		wantError bool
	}{
		{
			name: "Valid note",
			payload: map[string]string{
				"title":       "This is the title",
				"description": "This is the description",
			},
			wantCode:  http.StatusCreated,
			wantError: false,
		}, {
			name: "empty title",
			payload: map[string]string{
				"title":       "",
				"description": "No title",
			},
			wantCode:  http.StatusBadRequest,
			wantError: true,
		},
		{
			name:    "invalid JSON",
			payload: map[string]string{
				// invalid json
			},
			wantCode:  http.StatusBadRequest,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var bodyBytes []byte
			var err error

			// invalid json case handling
			if tt.name == "invalid JSON" {
				bodyBytes = []byte(`{invalid json}`)
			} else {
				bodyBytes, err = json.Marshal(tt.payload)
				if err != nil {
					t.Errorf("Failed to marshal payload %v", err)
				}
			}

			req := httptest.NewRequest(http.MethodPost, "/create", bytes.NewReader(bodyBytes))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			if rec.Code != tt.wantCode {
				t.Errorf("Expected code %d but got %d", rec.Code, tt.wantCode)
			}

			if tt.wantError {
				var resp map[string]interface{}

				if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
					t.Errorf("Got invalid json response")
				}

				if _, ok := resp["error"]; !ok {
					t.Errorf("Expected error field in json but got %s", rec.Body.String())
				}

			}
		})
	}
}

// Testing /notes handler
func TestGetAllNotes_Handler(t *testing.T) {
	service := NewService()
	service.Create("The very first note", "The description of the note")
	service.Create("The second note", "The description of the note 2")
	handler := NewHandler(service)
	router := setupRouter(handler)

	req := httptest.NewRequest(http.MethodGet, "/notes", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	var resp []Note

	// Checking for return code
	if rec.Code != http.StatusOK{
		t.Errorf("Expected code %d but got %d", http.StatusOK, rec.Code)
	}

	// checking if the returning json is valid array of notes
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Errorf("Got unexpected json response")
	}

	if len(resp) != 2 {
		t.Errorf("Expected 2 notes but got %d", len(resp))
	}

}
