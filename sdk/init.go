package sdk

import "regexp"

func init() {
	RePhone, _ = regexp.Compile(`(\d{3})\d*(\d{4})`)
	ReWhiteSpace, _ = regexp.Compile(`\s`)
	ReFieldWhiteList, _ = regexp.Compile(`^[A-Za-z0-9]+$`)
	ReUserName, _ = regexp.Compile("^[a-zA-Z0-9]+([-._][a-zA-Z0-9]+)*$")
	ReUserNameWithEmail, _ = regexp.Compile(`^([a-zA-Z0-9]+([-._][a-zA-Z0-9]+)*)|([a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,})$`) // Add support for email formats

	// used zh lang
	InitCasdoorsdkWithLang("zh")
}
