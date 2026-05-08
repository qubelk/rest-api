package books

import "errors"

var InvalidYearArgument error = errors.New("Invalid Year Argument: year can't be future")
var InvalidTitleArgument error = errors.New("Invalid Title Argument: title can't be empty")
var InvalidAuthorArgument error = errors.New("Invalid Author Argument: author can't be empty")
