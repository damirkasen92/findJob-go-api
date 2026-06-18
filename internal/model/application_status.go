package model

import "errors"

type ApplicationStatus string

const (
	ApplicationPending   ApplicationStatus = "pending"
	ApplicationReviewing ApplicationStatus = "reviewing"
	ApplicationRejected  ApplicationStatus = "rejected"
	ApplicationAccepted  ApplicationStatus = "accepted"
)

var ErrInvalidApplicationStatus = errors.New("invalid application status")

func (s ApplicationStatus) IsValid() bool {
	switch s {
	case ApplicationPending,
		ApplicationReviewing,
		ApplicationRejected,
		ApplicationAccepted:
		return true
	}
	return false
}
