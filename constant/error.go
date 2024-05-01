package constant

import "errors"

var ErrInsertDatabase error = errors.New("invalid Add Data in Database")
var ErrGetDatabase error = errors.New("invalid data not found")
var ErrEmptyInput error = errors.New("Data Cannot be empty")

// Register
var ErrAddUsersEmail error = errors.New("Invalid in user email data")
var ErrAddUsersPassword error = errors.New("Invalid in user password data")

// Login
var ErrLogin error = errors.New("Invalid data, user cannot be found")
