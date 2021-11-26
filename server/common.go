package server

import "errors"

var CANNOT_PARSE_BODY = errors.New("cannot parse request body")
var INVALID_OBJECT_ID = errors.New("invalid object id")
