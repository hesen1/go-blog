package util

import (
	"testing"
)

func TestJwt(t *testing.T) {
	payload := JwtClaims{
		Phone: "1078213694",
		Email: "test@example.com",
		Name: "Test",
	}

	tokenString, err := BuildJwt(payload, 120)

	if err != nil {
		t.Log(err)
		t.Fail()
	}
	// t.Fail()

	p, e := ParseJwt(tokenString)
	if e != nil {
		t.Log(err)
		t.Fail()
	}
	t.Log(p)
}
