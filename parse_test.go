package apachelogparser

import (
	"fmt"
	"strings"
	"testing"
)

func TestParseLogRecord(t *testing.T) {

	tests := []struct {
		record string
		log    interface{}
		err    error
	}{
		{"127.0.0.1 - frank [10/Oct/2000:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326 xyz", nil, fmt.Errorf("Log record is neither Apache common nor combined log")},
		{"127.0.0.1 - - dda [10/Oct/2000:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326", nil, fmt.Errorf("Log record is neither Apache common nor combined log")},
		{"[10/Oct/2000:13:55:36 -0700] 200 2326", nil, fmt.Errorf("Log record is neither Apache common nor combined log")},
		{"[10/Oct/2000:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326", nil, fmt.Errorf("Log record is neither Apache common nor combined log")},
	}
	t.Log("===========================================")
	t.Log("Check non-standard log records (fewer or more fields)")
	t.Log("-------------------------------------------")
	for _, v := range tests {
		t.Log("checking ...", v.record)
		//record := strings.Join(v.record, " ")
		log, err := ParseLogRecord(v.record)
		if err != nil {
			k := strings.Compare(err.Error(), v.err.Error())
			if log != nil || k != 0 {
				t.Errorf("log record is neither Apache common nor combined log, no corresponding error returned")
			}
		} else {
			t.Errorf("log record is neither Apache common nor combined log, however error is nil")
		}

	}

	tests = []struct {
		record string
		log    interface{}
		err    error
	}{
		{"0.01 - rob [10/Oct/2000:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326", nil, fmt.Errorf("IP invalid")},
		{"1270.0.1 - rob [10/Oct/2000:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326", nil, fmt.Errorf("IP invalid")},
		{"abcdef - rob [10/Oct/2000:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326", nil, fmt.Errorf("IP invalid")},
	}
	t.Log("===========================================")
	t.Log("Check IP validity")
	t.Log("-------------------------------------------")
	for _, v := range tests {
		t.Log("checking ...", v.record)
		//record := strings.Join(v.record, " ")
		log, err := ParseLogRecord(v.record)
		if err != nil {
			k := strings.Compare(err.Error(), v.err.Error())
			if log != nil || k != 0 {
				t.Errorf("IP invalid, no corresponding error returned")
			}
		} else {
			t.Errorf("IP invalid, however error is nil")
		}

	}

}
