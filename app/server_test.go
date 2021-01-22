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
	data := "+PING\r\n"
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
