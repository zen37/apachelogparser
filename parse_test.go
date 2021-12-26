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
		{"[10/Oct/2000:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326", nil, fmt.Errorf("Log record is neither Apache common nor combined log")},
		{"[10/Oct/2000:13:55:36 -0700] 200 2326", nil, fmt.Errorf("Log record is neither Apache common nor combined log")},
	}
	t.Log("===========================================")
	t.Log("Check non-standard log records - additional or fewer fields")
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

	/* 	tests = []struct {
	   		record []string
	   		log    interface{}
	   		err    error
	   	}{
	   		{[]string{"1270.1 - frank [10/Oct/2000:13:55:36 -0700]", "GET /apache_pb.gif HTTP/1.0", "200 2326"}, nil, fmt.Errorf("IP not valid")},
	   		{[]string{"127.0.01 - - dda [10/Oct/2000:13:55:36 -0700]", "GET /apache_pb.gif HTTP/1.0", "200 2326"}, nil, fmt.Errorf("IP not valid")},
	   	}
	   	t.Log("===========================================")
	   	t.Log("Checking IP validity")
	   	t.Log("-------------------------------------------")
	   	for _, v := range tests {
	   		record := strings.Join(v.record, " ")
	   		log, err := ParseLogRecord(record)
	   		if log != nil || err != v.err {
	   			t.Errorf("IP invalid, no corresponding error returned")
	   		}
	   	} */

}
