package main

import (
	"container/list"
	"testing"
)

func TestParse_should_parse_empty_arr(t *testing.T) {
	ret, err := parse("*0\r\n")
	if err != nil {
		t.Errorf("Error - %+v", err.Error())
	}
	l, _ := ret.(list.List)

	if l.Len() != 0 {
		t.Errorf("Return length is not 0, %d", l.Len())
	}
}

func TestParse_should_parse_bulk_str_arr(t *testing.T) {
	ret, err := parse("*2$3\r\nfoo\r\n$3\r\nbar\r\n")
	if err != nil {
		t.Errorf("Error - %+v", err.Error())
	}

	l, _ := ret.(*list.List)

	if l.Len() != 2 {
		t.Errorf("Return length is not 2, %d", l.Len())
	}

	elem := l.Front()
	if elem.Value != "foo" || elem.Next().Value != "bar" {
		t.Errorf("Content error, %+v", ret)
	}
}

func TestGenerateResponse_should_return_pong_for_ping(t *testing.T) {
	data := "*1\r\n$4\r\nping\r\n"
	res, err := generateResponse([]byte(data))
	if err != nil {
		t.Errorf("Failed to generate, %+v", err)
	}

	strRes := string(res)
	if strRes != "+PONG\r\n" {
		t.Errorf("Wrong response")
	}
}

func TestGenerateResponse_should_return_echo(t *testing.T) {
	data := "*2\r\n$4\r\nECHO\r\n$3\r\nhey\r\n"
	res, err := generateResponse([]byte(data))
	if err != nil {
		t.Errorf("Failed to generate, %+v", err)
	}

	strRes := string(res)
	if strRes != "+hey\r\n" {
		t.Errorf("Wrong response")
	}
}

func TestGet_should_return_nil_for_no_key(t *testing.T) {
	data := "*2\r\n$3\r\nGET\r\n$5\r\nhello\r\n"
	res, err := generateResponse([]byte(data))
	if err != nil {
		t.Errorf("Failed to get %+v", err)
	}

	strRes := string(res)
	if strRes != "$-1\r\n" {
		t.Errorf("Wrong response for get")
	}
}

func TestSet_should_return_ok(t *testing.T) {
	data := "*3\r\n$3\r\nSET\r\n$5\r\nhello\r\n$5\r\nworld\r\n"
	res, err := generateResponse([]byte(data))
	if err != nil {
		t.Errorf("Failed to set %+v", err)
	}

	strRes := string(res)
	if strRes != "+OK\r\n" {
		t.Errorf("Wrong response")
	}
}

func TestSet_with_px_should_return_ok(t *testing.T) {
	data := "*3\r\n$3\r\nSET\r\n$5\r\nhello\r\n$5\r\nworld\r\n$2\r\npx\r\n:300\r\n"
	res, err := generateResponse([]byte(data))
	if err != nil {
		t.Errorf("Failed to set %+v", err)
	}

	strRes := string(res)
	if strRes != "+OK\r\n" {
		t.Errorf("Wrong response")
	}
}

func TestGet_should_return_ok(t *testing.T) {
	data := "*3\r\n$3\r\nSET\r\n$5\r\nhello\r\n$5\r\nworld\r\n"
	res, err := generateResponse([]byte(data))
	if err != nil {
		t.Errorf("Failed to set %+v", err)
	}

	strRes := string(res)
	if strRes != "+OK\r\n" {
		t.Errorf("Wrong response for set")
	}

	data = "*2\r\n$3\r\nGET\r\n$5\r\nhello\r\n"
	res, err = generateResponse([]byte(data))
	if err != nil {
		t.Errorf("Failed to get %+v", err)
	}

	strRes = string(res)
	if strRes != "$5\r\nworld\r\n" {
		t.Errorf("Wrong response for get")
	}
}
