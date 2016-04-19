package gosession

import "testing"

func testJWT(t *testing.T, kid string, m interface{}) {
	jwt, err := NewJWT("./")
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

func TestJWTExampleHmac(t *testing.T) {
	var m = make(map[string]interface{})
	m["uid"] = 19
	m["username"] = "the name of user"
	m["中文"] = "这是个测试123abc"
	testJWT(t, "hmac_demo", m)
}

func TestJWTExampleRSA(t *testing.T) {
	var m = make(map[string]interface{})
	m["uid"] = 21
	m["username"] = "the name of admin"
	m["中文"] = "这是个测试123abc"
	testJWT(t, "rsa_demo", m)
}
