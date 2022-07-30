package main

import "regexp"

// sanitiseLeafName replaces any illegal (under Windows) filename character
// and replaces it with an underscore
func sanitiseLeafName(filename string) string {
	re := regexp.MustCompile(`[/\\?%*:|"<>]`)
	return re.ReplaceAllString(filename, "_")
}
