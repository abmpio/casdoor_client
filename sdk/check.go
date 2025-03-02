package sdk

const (
	ErrorMsg_EmptyUsername     = "用户名不可为空"
	ErrorMsg_UserNameIsTooLong = "用户名过长(最大允许长度为39个字符)"
	ErrorMsg_UserNameInvalid   = "用户名只能包含字母数字字符、下划线或连字符，不能有连续的连字符或下划线，也不能以连字符或下划线开头或结尾"
)

func CheckUsername(username string) string {
	if username == "" {
		return ErrorMsg_EmptyUsername
	} else if len(username) > 39 {
		return ErrorMsg_UserNameIsTooLong
	}

	// https://stackoverflow.com/questions/58726546/github-username-convention-using-regex

	if !ReUserName.MatchString(username) {
		return ErrorMsg_UserNameInvalid
	}

	return ""
}
