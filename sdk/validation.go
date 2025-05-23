// Copyright 2021 The Casdoor Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sdk

import (
	"fmt"
	"net/mail"
	"regexp"
	"strings"

	"github.com/nyaruka/phonenumbers"
)

var (
	RePhone             *regexp.Regexp
	ReWhiteSpace        *regexp.Regexp
	ReFieldWhiteList    *regexp.Regexp
	ReUserName          *regexp.Regexp
	ReUserNameWithEmail *regexp.Regexp
)

func IsEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsPhoneValid(phone string, countryCode string) bool {
	phoneNumber, err := phonenumbers.Parse(phone, countryCode)
	if err != nil {
		return false
	}
	return phonenumbers.IsValidNumber(phoneNumber)
}

func IsPhoneAllowInRegin(countryCode string, allowRegions []string) bool {
	if ContainsString(allowRegions, "All") {
		return true
	}
	return ContainsString(allowRegions, countryCode)
}

func IsRegexp(s string) (bool, error) {
	if _, err := regexp.Compile(s); err != nil {
		return false, err
	}
	return regexp.QuoteMeta(s) != s, nil
}

func IsInvitationCodeMatch(pattern string, invitationCode string) (bool, error) {
	if !strings.HasPrefix(pattern, "^") {
		pattern = "^" + pattern
	}
	if !strings.HasSuffix(pattern, "$") {
		pattern = pattern + "$"
	}
	return regexp.MatchString(pattern, invitationCode)
}

func GetE164Number(phone string, countryCode string) (string, bool) {
	phoneNumber, _ := phonenumbers.Parse(phone, countryCode)
	return phonenumbers.Format(phoneNumber, phonenumbers.E164), phonenumbers.IsValidNumber(phoneNumber)
}

func GetCountryCode(prefix string, phone string) (string, error) {
	if prefix == "" || phone == "" {
		return "", nil
	}

	phoneNumber, err := phonenumbers.Parse(fmt.Sprintf("+%s%s", prefix, phone), "")
	if err != nil {
		return "", err
	}

	countryCode := phonenumbers.GetRegionCodeForNumber(phoneNumber)
	if countryCode == "" {
		return "", fmt.Errorf("country code not found for phone prefix: %s", prefix)
	}

	return countryCode, nil
}

func FilterField(field string) bool {
	return ReFieldWhiteList.MatchString(field)
}
