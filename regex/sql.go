package regex

import "regexp"

// 开发中
type Rule struct {
	Name string
	Reg  *regexp.Regexp
}

// re := regexp.MustCompile(`'";#--`\\%_()|&<>@\[\]`)

// re := regexp.MustCompile(`%&?+@*\[\]`)

// re := regexp.MustCompile(`(?i)%09|%0A|%0D|%0B|%0C|/\*\*/`)

// re := regexp.MustCompile(`(?i)union|select|insert|update|delete|drop|alter|create|rename|load_file|into|outfile`)

// regexp.MustCompile(`(?i)(and|or|not|xor|[\(\)\|&^!=<>]|>=|<=)`)

var SQLList = []*regexp.Regexp{
	regexp.MustCompile(`[%&?+@*\[\]]`),
	regexp.MustCompile(`%09|%0A|%0B|%0C|%0D|/\*\*/`),
	regexp.MustCompile(`(?i)union|select|insert|update|delete`),
	regexp.MustCompile(`(?i)and|or|not|xor`),
	regexp.MustCompile(`[\(\)\|&^!=<>]|>=|<=`),
}

func SQL_Injection(Full_req string) (bool, string) {
	for _, reg := range SQLList {
		if reg.MatchString(Full_req) {
			return true, "SQL注入"
		}
	}
	return false, ""
}
