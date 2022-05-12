package utils

import (
	"regexp"
	"server/structs"
)


func GetRequestType(text string) structs.RequestType {
	var addRegex, _ = regexp.Compile("^-일정 +추가")
	var getRegex, _ = regexp.Compile("^-일정 +조회")
	var deleteRegex, _ = regexp.Compile("^-일정 +삭제")
	
	if addRegex.MatchString(text) { return structs.REQUEST_ADD }
	if getRegex.MatchString(text) { return structs.REQUEST_GET }
	if deleteRegex.MatchString(text) { return structs.REQUEST_DELETE }

	return structs.REQUEST_ERROR
}
