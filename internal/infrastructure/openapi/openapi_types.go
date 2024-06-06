// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package openapi

import (
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Error defines model for Error.
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// Fine defines model for Fine.
type Fine struct {
	Amount      float32             `json:"amount"`
	CreatedAt   *time.Time          `json:"created_at,omitempty"`
	Description *string             `json:"description,omitempty"`
	DueDate     openapi_types.Date  `json:"due_date"`
	Id          *openapi_types.UUID `json:"id,omitempty" validate:"required,uuid"`
	IssueDate   openapi_types.Date  `json:"issue_date"`
	Status      string              `json:"status"`
	UpdatedAt   *time.Time          `json:"updated_at,omitempty"`
	VehicleId   openapi_types.UUID  `json:"vehicle_id"`
}

// Notification defines model for Notification.
type Notification struct {
	CreatedAt          *time.Time          `json:"created_at,omitempty"`
	FineId             openapi_types.UUID  `json:"fine_id"`
	Id                 *openapi_types.UUID `json:"id,omitempty"`
	NotificationStatus string              `json:"notification_status"`
	NotificationType   string              `json:"notification_type"`
	OwnerId            openapi_types.UUID  `json:"owner_id"`
	SentAt             *time.Time          `json:"sent_at,omitempty"`
	UpdatedAt          *time.Time          `json:"updated_at,omitempty"`
}

// Owner defines model for Owner.
type Owner struct {
	CreatedAt     *time.Time          `json:"created_at,omitempty"`
	Email         string              `json:"email"`
	FullName      string              `json:"full_name"`
	Id            *openapi_types.UUID `json:"id,omitempty"`
	LicenseNumber string              `json:"license_number"`
	Phone         string              `json:"phone"`
	UpdatedAt     *time.Time          `json:"updated_at,omitempty"`
}

// Payment defines model for Payment.
type Payment struct {
	Amount        float32             `json:"amount"`
	CreatedAt     *time.Time          `json:"created_at,omitempty"`
	FineId        openapi_types.UUID  `json:"fine_id"`
	Id            *openapi_types.UUID `json:"id,omitempty"`
	PaidDate      *time.Time          `json:"paid_date,omitempty"`
	PaymentMethod string              `json:"payment_method"`
	UpdatedAt     *time.Time          `json:"updated_at,omitempty"`
}

// Vehicle defines model for Vehicle.
type Vehicle struct {
	CreatedAt    *time.Time          `json:"created_at,omitempty"`
	Id           *openapi_types.UUID `json:"id,omitempty"`
	LicensePlate string              `json:"license_plate"`
	Model        string              `json:"model"`
	OwnerId      openapi_types.UUID  `json:"owner_id"`
	UpdatedAt    *time.Time          `json:"updated_at,omitempty"`
}

// N400 defines model for 400.
type N400 = Error

// CreateFineJSONRequestBody defines body for CreateFine for application/json ContentType.
type CreateFineJSONRequestBody = Fine

// CreateNotificationJSONRequestBody defines body for CreateNotification for application/json ContentType.
type CreateNotificationJSONRequestBody = Notification

// CreateOwnerJSONRequestBody defines body for CreateOwner for application/json ContentType.
type CreateOwnerJSONRequestBody = Owner

// CreatePaymentJSONRequestBody defines body for CreatePayment for application/json ContentType.
type CreatePaymentJSONRequestBody = Payment

// CreateVehicleJSONRequestBody defines body for CreateVehicle for application/json ContentType.
type CreateVehicleJSONRequestBody = Vehicle
