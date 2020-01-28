package main

import (
	"fmt"
	"testing"
)

type pingDataFormat struct {
	UserAccessToken          string `json:"userAccessToken"`
	UploadStartTimeInSeconds int    `json:"uploadStartTimeInSeconds"`
	UploadEndTimeInSeconds   int    `json:"uploadEndTimeInSeconds"`
	CallbackURL              string `json:"callbackURL"`
}

func TestDecode(t *testing.T) {
	pingJSON := make(map[string][]pingDataFormat)
	cases := []struct {
		inJson           string
		inTarget         interface{}
		wantObjectString string
	}{
		{`{"dailies":[{"userAccessToken":"acessToken","uploadStartTimeInSeconds":1499744832,"uploadEndTimeInSeconds":1499744832,"callbackURL":"callbackurl"}]}`, &pingJSON, `&map[dailies:[{UserAccessToken:acessToken UploadStartTimeInSeconds:1499744832 UploadEndTimeInSeconds:1499744832 CallbackURL:callbackurl}]]`},
	}
	for _, c := range cases {
		Decode(c.inJson, c.inTarget)
		gotObjectString := fmt.Sprintf("%+v", c.inTarget)
		if gotObjectString != c.wantObjectString {
			t.Errorf("Decode(%q) == %q, want %q", c.inJson, gotObjectString, c.wantObjectString)
		}
	}
}
