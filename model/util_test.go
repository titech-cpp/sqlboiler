package model

import "testing"

func TestNameDetail(t *testing.T) {
	nd := new(NameDetail)
	result,err := nd.snakeToCamel("aaa_bbb", true)
	if err != nil {
		t.Fatalf("Unexpected SnakeToCamel Error: %#v", err)
	}
	if result != "AaaBbb" {
		t.Fatalf("Invalid SnakeToCamel Value %s, Expected AaaBbb", result)
	}

	result,err = nd.snakeToCamel("aaa_bbb", false)
	if err != nil {
		t.Fatalf("Unexpected SnakeToCamel Error: %#v", err)
	}
	if result != "aaaBbb" {
		t.Fatalf("Invalid SnakeToCamel Value %s, Expected aaaBbb", result)
	}

	result,err = nd.snakeToCamel("Aaa_bbb", true)
	if err == nil {
		t.Fatal("Unexpected No Error(Aaa_bbb)")
	}

	_,err = nd.snakeToCamel("Aaa__bbb", true)
	if err == nil {
		t.Fatal("Unexpected No Error(Aaa__bbb)")
	}

	nameDetail,err := NewNameDetail("aaa_bbb")
	if err != nil {
		t.Fatalf("Unexpected NewNameDetails Error: %#v", err)
	}
	if nameDetail.LowerCamel != "aaaBbb" {
		t.Fatalf("Unexpected LowerCamel Value: %#v", nameDetail)
	}
	if nameDetail.UpperCamel != "AaaBbb" {
		t.Fatalf("Unexpected UpperCamel Value: %#v", nameDetail)
	}
	if nameDetail.Snake != "aaa_bbb" {
		t.Fatalf("Unexpected Snake Value: %#v", nameDetail)
	}
}