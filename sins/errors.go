package sins

import "errors"

var InvalidYearArgument error = errors.New("Invalid Year Argument: year can't be future")
var InvalidTitleArgument error = errors.New("Invalid Title Argument: title can't be empty")
var InvalidAuthorArgument error = errors.New("Invalid Author Argument: author can't be empty")
var InvalidQueryAuthorArgument error = errors.New("Invalid Author Query Arguments: query can't be empty")
var BookAlreadyExists error = errors.New("Book Already Exists: can't add new book while have same")
var BookNotExists error = errors.New("Book Not Exists: can't find book in library")
