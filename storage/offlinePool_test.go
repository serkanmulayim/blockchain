package storage

import "testing"

func TestVerifyAddr(t *testing.T) {
	//ipv4
	ipport := "10.10.10.10:2334"
	if !verifyAddr(ipport) {
		t.Error("failed in test, expected true")
	}

	if ipport = "256.42.32.32:122"; verifyAddr(ipport) {
		t.Error("failed in test, expected false")
	}

	if ipport = "156.42.32.32:"; verifyAddr(ipport) {
		t.Error("failed in test, expected false")
	}

	if ipport = "156.42.32.32:abc"; verifyAddr(ipport) {
		t.Error("failed in test, expected false")
	}

	if ipport = ":"; verifyAddr(ipport) {
		t.Error("failed in test, expected false")
	}

	if ipport = ":3223"; verifyAddr(ipport) {
		t.Error("failed in test, expected false")
	}

	if ipport = "156.42.32.32:1222"; !verifyAddr(ipport) {
		t.Error("failed in test, expected true")
	}

	//ipv6
	if ipport = "2001:0db8:85a3:0000:0000:8a2e:0370:7334:1222"; !verifyAddr(ipport) {
		t.Error("failed in test, expected true")
	}

	if ipport = "2001:0db8:85a3:0000:0000:8a2e:0370:1222"; verifyAddr(ipport) {
		t.Error("failed in test, expected false")
	}
	if ipport = "2001:0db8:85a3:0000:0000:8a2e:0370"; verifyAddr(ipport) {
		t.Error("failed in test, expected false")
	}

	if ipport = "2001:0db8:85a3:000x:0000:8a2e:0370:7334:1222"; verifyAddr(ipport) {
		t.Error("failed in test, expected false")
	}

	if ipport = "2001:0db8:85a3:0000:0000:8a2e:0370:7334:66000"; verifyAddr(ipport) {
		t.Error("failed in test, expected false")
	}

	if ipport = "2001:0db8:85a3:0000:0000:8a2e:0370:7334:"; verifyAddr(ipport) {
		t.Error("failed in test, expected false")
	}
	if ipport = "2001:0db8:85a3:0000:0000:8a2e:0370:7334:-1"; verifyAddr(ipport) {
		t.Error("failed in test, expected false")
	}
	if ipport = ":11111"; verifyAddr(ipport) {
		t.Error("failed in test, expected false")
	}

}
