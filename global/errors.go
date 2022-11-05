package global

import "errors"

var ErrorNewRedisError = errors.New("error for new redis client")
var ErrorNilDBName = errors.New("error for nil db name")
