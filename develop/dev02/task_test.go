package main

import "testing"

func TestUnpackStrEmptyString(t *testing.T) {
	res, err := UnpackStr("")

	if res != "" || err != nil {
		t.Fatalf(`unpack("") = %q, %v, expected "", nil`, res, err)
	}
}

func TestUnpackStrCorrectStringNoEscape(t *testing.T) {
	s := "a10bc1d5efgh2"
	expected := "aaaaaaaaaabcdddddefghh"

	res, err := UnpackStr(s)

	if res != expected || err != nil {
		t.Fatalf(`unpack(%q) = %q, %v, expected %q, nil`, s, res, err, expected)
	}
}

func TestUnpackStrCorrectStringWithEscape(t *testing.T) {
	s := `qwe\4\5\03\\4`
	expected := `qwe45000\\\\`

	res, err := UnpackStr(s)

	if res != expected || err != nil {
		t.Fatalf(`unpack(%v) = %v, %v, expected %v, nil`, s, res, err, expected)
	}
}

func TestUnpackStrStartsWithANumber(t *testing.T) {
	res, err := UnpackStr("45")

	if res != "" || err == nil {
		t.Fatalf(`unpack("45") = %q, %v, expected "", error`, res, err)
	}
}

func TestUnpackStrWithZero(t *testing.T) {
	res, err := UnpackStr("qw0e")
	expected := "qe"

	if res != expected || err != nil {
		t.Fatalf(`unpack("qw0e") = %q, %v, expected "", error`, res, err)
	}
}

func TestUnpackStrWrongEscape(t *testing.T) {
	res, err := UnpackStr(`qwe\\\`)

	if res != "" || err == nil {
		t.Fatalf(`unpack("qwe\\\") = %v, %v, expected "", error`, res, err)
	}
}
