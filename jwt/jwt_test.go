package gosession

import "testing"

// method of running jwt
func testJWT(t *testing.T, kid string, m interface{}) {
	jwt, err := NewJWT("./") // init jwt
	if err != nil {
		t.Fatal(err)
	}
	token, err := jwt.Sign(kid, m, 72)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("token: %v\n", token)
	arg, err := jwt.Parse(token)
	if err != nil {
		t.Error(err)
		return
	} else {
		t.Log(arg)
	}
}

// the example for hmac key
func TestJWTExampleHmac(t *testing.T) {
	var m = make(map[string]interface{})
	m["uid"] = 19
	m["username"] = "the name of user"
	m["中文"] = "这是个测试123abc"
	testJWT(t, "hmac_test", m)
}

// the example for rsa key
func TestJWTExampleRSA(t *testing.T) {
	var m = make(map[string]interface{})
	m["uid"] = 21
	m["username"] = "the name of admin"
	m["中文"] = "这是个测试123abc"
	testJWT(t, "rsa_test", m)
}
