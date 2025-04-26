package sdk

import (
	"github.com/abmpio/libx/slicex"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

const (
	// app signup item
	SignupItem_Key_ID              string = "ID"
	SignupItem_Key_UserName        string = "username"
	SignupItem_Key_DisplayName     string = "Display name"
	SignupItem_Key_Password        string = "Password"
	SignupItem_Key_ConfirmPassword string = "Confirm password"
	SignupItem_Key_Email           string = "Email"
	SignupItem_Key_Phone           string = "Phone"
	SignupItem_Key_Agreement       string = "Agreement"
	SignupItem_Key_SignupButton    string = "Signup button"

	// signup item rule for email,phone no verification
	SignupItemRule_NoVerification string = "No verification"
	SignupItemRule_Normal         string = "Normal"
)

// get application signup item by key
func GetSignupItem(app *casdoorsdk.Application, key string) *casdoorsdk.SignupItem {
	if len(app.SignupItems) <= 0 {
		return nil
	}
	return slicex.FindOne(app.SignupItems, func(item *casdoorsdk.SignupItem) bool {
		return item.Name == key
	})
}
