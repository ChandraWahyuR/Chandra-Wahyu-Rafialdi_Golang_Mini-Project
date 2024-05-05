package constant

import "errors"

// Insert
var ErrInsertDatabase error = errors.New("invalid Add Data in Database")

// Delete
var ErrDeleteData error = errors.New("Error, cannot delete data")

// Get By Id
var ErrFindData error = errors.New("Error, cannot find data by id")

//
var ErrGetDatabase error = errors.New("invalid data not found")
var ErrEmptyInput error = errors.New("Data Cannot be empty")

// Register
var ErrAddUsersEmail error = errors.New("Invalid in user email data")
var ErrAddUsersPassword error = errors.New("Invalid in user password data")

// Login
var ErrLogin error = errors.New("Invalid data, user cannot be found")

// Rent_confirm
var ErrEmptyAddress error = errors.New("Address Cannot be empty")
var ErrEmptyStatus error = errors.New("Status Cannot be empty")
