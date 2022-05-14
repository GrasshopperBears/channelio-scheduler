package utils

import (
	"regexp"
	"server/structs"
	"server/texts"
)

var helpRegex, _ = regexp.Compile("^" + texts.SCHEDULER_PREFIX + " +" + texts.SCHEDULER_HELP + " *$")
var addRegex, _ = regexp.Compile("^" + texts.SCHEDULER_PREFIX + " +" + texts.SCHEDULER_ADD)
var getRegex, _ = regexp.Compile("^" + texts.SCHEDULER_PREFIX + " +" + texts.SCHEDULER_GET + " *$")
var deleteRegex, _ = regexp.Compile("^" + texts.SCHEDULER_PREFIX + " +" + texts.SCHEDULER_DELETE)

func GetRequestType(text string) structs.RequestType {
	if helpRegex.MatchString(text) { return structs.REQUEST_HELP }
	if addRegex.MatchString(text) { return structs.REQUEST_ADD }
	if getRegex.MatchString(text) { return structs.REQUEST_GET }
	if deleteRegex.MatchString(text) { return structs.REQUEST_DELETE }

	return structs.REQUEST_ERROR
}
