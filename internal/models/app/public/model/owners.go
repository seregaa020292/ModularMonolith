//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"github.com/google/uuid"
	"time"
)

type Owners struct {
	ID            uuid.UUID `sql:"primary_key"`
	FullName      string
	LicenseNumber string
	Phone         string
	Email         string
	CreatedAt     time.Time
	UpdatedAt     *time.Time
}
