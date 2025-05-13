package sdk

import (
	"github.com/abmpio/libx/slicex"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
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
