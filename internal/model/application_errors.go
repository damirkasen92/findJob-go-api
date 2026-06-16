package model

import "errors"

var (
	ErrAlreadyApplied = errors.New(
		"already applied",
	)
	ErrOwnVacancy = errors.New(
		"cannot apply to own vacancy",
	)
	ErrForeignResume = errors.New(
		"resume does not belong to user",
	)
)
