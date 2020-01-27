package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
	"time"
	"github.com/gobuffalo/validate/validators"
)
// Redirection is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Redirection struct {
    ID uuid.UUID `json:"id" db:"id"`
    Email string `json:"email" db:"email"`
    IsVerified bool `json:"is_verified" db:"is_verified"`
    Subdomain string `json:"subdomain" db:"subdomain"`
    Destination string `json:"destination" db:"destination"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (r Redirection) String() string {
	jr, _ := json.Marshal(r)
	return string(jr)
}

// Redirections is not required by pop and may be deleted
type Redirections []Redirection

// String is not required by pop and may be deleted
func (r Redirections) String() string {
	jr, _ := json.Marshal(r)
	return string(jr)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (r *Redirection) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: r.Email, Name: "Email"},
		&validators.StringIsPresent{Field: r.Subdomain, Name: "Subdomain"},
		&validators.StringIsPresent{Field: r.Destination, Name: "Destination"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (r *Redirection) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (r *Redirection) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
