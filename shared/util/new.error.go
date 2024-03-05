package util

import "errors"

var (
	TitleIsRequired    = errors.New("Title is required")
	ActivityGroupId    = errors.New("Activity group id is required")
	ActivityIdNotFound = errors.New("Activity id not found")
	ActivityNotFound   = errors.New("Activity not found")
	TodoIdNotFound     = errors.New("Todo id not found")
	TodoNotFound       = errors.New("Todo not found")
)
