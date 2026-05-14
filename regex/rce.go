package regex

import "regexp"

// 开发中
var RceList = []*regexp.Regexp{

	regexp.MustCompile(`(?i)\b(eval|assert|preg_replace|call_user_func|call_user_func_array)\b`),
	regexp.MustCompile(`(?i)\b(system|exec|shell_exec|passthru|popen|proc_open)\b`),
	regexp.MustCompile(`(;|&&|\|\||\|)`),
	regexp.MustCompile(`(?i)\$IFS|bin\/bash|bin\/sh|tac|cat.*\*|echo.*>.*\.php`),
}

func RCE(Full_req string) (bool, string) {
	for _, reg := range RceList {
		if reg.MatchString(Full_req) {
			return true, "命令执行"
		}
	}
	return false, ""
}
