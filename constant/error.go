package constant

import "errors"

var ErrFetchData error = errors.New("error, cannot fetch data from Database")

// Insert
var ErrInsertDatabase error = errors.New("invalid Add Data in Database")
var ErrDataNotFound error = errors.New("Data not found in Database")

// image
var ErrGetImage error = errors.New("Error, cannot get image")
var ErrFetchImage error = errors.New("Error, cannot fetch image")

// Delete
var ErrDeleteData error = errors.New("Error, cannot delete data")

// Get By Id
var ErrFindData error = errors.New("Error, cannot find data by id")
var ErrById error = errors.New("Error, cannot find id")
var ErrGetDataFromId error = errors.New("Error, id didnt have any data")
var ErrUpdateRentData error = errors.New("cannot update data, data has been confirmed")
var ErrGetDataID error = errors.New("An error occurred while searching for data.")

//
var ErrUpdateData error = errors.New("error, cannot update data")
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

// Info Rental
var ErrInvalidConfirmationStatus error = errors.New("Invalid confirmation status")
var ErrInvalidReturnConfirmation error = errors.New("cannot return confirmation status")
var ErrRentalNotAccepted error = errors.New("error, data has been return")
