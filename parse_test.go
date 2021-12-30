package apachelogparser

import (
	"errors"
	"reflect"
	"testing"
)

func TestParseLogRecord(t *testing.T) {

	tests := []struct {
		record string
		log    interface{}
	}{
		{"127.0.0.1 - frank [10/Oct/2000:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326 xyz", nil},
		{"127.0.0.1 - - dda [10/Oct/2000:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326", nil},
		{"[10/Oct/2000:13:55:36 -0700] 200 2326", nil},
		{"[10/Oct/2000:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326", nil},
	}
	t.Log("===========================================")
	t.Log("Check non-standard log records (fewer or more fields)")
	t.Log("-------------------------------------------")
	for _, v := range tests {
		t.Log("checking ...", v.record)
		rec, err := ParseLogRecord(v.record)
		if !isNil(rec) {
			t.Errorf("FAIL: returned log is not nil")
		}
		if !errors.Is(err, ErrInvalidLog) {
			t.Errorf("FAIL: log is not marked as invalid")
		}

	}

	tests = []struct {
		record string
		log    interface{}
	}{
		{"0.01 - rob [10/Oct/2000:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326", nil},
		{"1270.0.1 - rob [10/Oct/2000:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326", nil},
		{"abcdef - rob [10/Oct/2000:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326", nil},
	}
	t.Log("===========================================")
	t.Log("Check whether IP is invalid")
	t.Log("-------------------------------------------")
	for _, v := range tests {
		t.Log("checking ...", v.record)
		_, err := ParseLogRecord(v.record)
		if !errors.Is(err, ErrInvalidIP) {
			t.Errorf("FAIL: IP not marked as invalid")
		}
	}
}

func isNil(v interface{}) bool {
	return v == nil || (reflect.ValueOf(v).Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil())
}
