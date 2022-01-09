package apachelogparser

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestParseLogRecord(t *testing.T) {

	tests := []struct {
		record string
		log    interface{}
		err    error
	}{
		{"127.0.0.1 - frank [10/Oct/2000:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326 xyz", nil, ErrInvalidLog},
		{"127.0.0.1 - - dda [10/Oct/2000:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326", nil, ErrInvalidLog},
		{"[10/Oct/2000:13:55:36 -0700] 200 2326", nil, ErrInvalidLog},
		{"[10/Oct/2000:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326", nil, ErrInvalidLog},
	}
	t.Log("===========================================")
	t.Log("Check non-standard log records (fewer or more fields than Common Log Format or Combined Log Format)")
	t.Log("-------------------------------------------")
	for _, v := range tests {
		t.Log("checking ...", v.record)
		rec, err := ParseLogRecord(v.record)
		if !isNil(rec) {
			t.Errorf("FAIL: returned log is not nil")
		}
		if !errors.Is(err, v.err) {
			t.Errorf("FAIL: log is not marked as invalid")
		}

	}

	tests = []struct {
		record string
		log    interface{}
		err    error
	}{
		{"0.01 - rob [10/Oct/2000:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326", nil, ErrInvalidIP},
		{"1270.0.1 - rob [10/Oct/2000:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326", nil, ErrInvalidIP},
		{"abcdef - rob [10/Oct/2000:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326", nil, ErrInvalidIP},
	}
	t.Log("===========================================")
	t.Log("Check invalid IP")
	t.Log("-------------------------------------------")
	for _, v := range tests {
		t.Log("checking ...", v.record)
		_, err := ParseLogRecord(v.record)
		if !errors.Is(err, v.err) {
			t.Errorf("FAIL: IP not marked as invalid")
		}
	}

	tests = []struct {
		record string
		log    interface{}
		err    error
	}{
		{"127.0.110.1  - rob 10/Oct/2000:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326", nil, ErrInvalidFieldTimestamp},
		{"127.0.0.1 - rob [10/Oct/2000:13:55:36 -0700 \"GET /apache_pb.gif HTTP/1.0\" 200 2326", nil, ErrInvalidFieldTimestamp},
		{"127.255.0.123 - rob &10/Oct/2000:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326", nil, ErrInvalidFieldTimestamp},
		{"127.1.1.1 - rob [10/Oct/2000:13:55:36 -0700) \"GET /apache_pb.gif HTTP/1.0\" 200 2326", nil, ErrInvalidFieldTimestamp},
	}
	t.Log("===========================================")
	t.Log("Check Timestamp field starts with [ and ends with ]")
	t.Log("-------------------------------------------")
	for _, v := range tests {
		t.Log("checking ...", v.record)
		_, err := ParseLogRecord(v.record)
		if !errors.Is(err, v.err) {
			t.Errorf("FAIL: Timestamp field not marked as invalid")
		}
	}

	tests = []struct {
		record string
		log    interface{}
		err    error
	}{
		{"127.0.110.1  - rob [x10/Oct/2000:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326", nil, ErrInvalidTimestamp},
		{"127.0.0.1 - rob [y10/Dec/2000:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326", nil, ErrInvalidTimestamp},
		{"127.255.0.123 - rob [10/Oct/2222:13:55:36 -250] \"GET /apache_pb.gif HTTP/1.0\" 200 2326", nil, ErrInvalidTimestamp},
		{"127.1.1.1 - - [10/Dec/2000:13:55:36 &800] \"GET /apache_pb.gif HTTP/1.0\" 200 2326", nil, ErrInvalidTimestamp},
	}
	t.Log("===========================================")
	t.Log("Check invalid Timestamp ")
	t.Log("-------------------------------------------")
	for _, v := range tests {
		t.Log("checking ...", v.record)
		_, err := ParseLogRecord(v.record)
		if !errors.Is(err, v.err) {
			t.Errorf("FAIL: Timestamp not marked as invalid")
		}
	}

	tests = []struct {
		record string
		log    interface{}
		err    error
	}{
		{"127.0.110.1  - rob [10/Oct/2000:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" abc1 2326", nil, ErrInvalidStatus},
		{"40.77.189.89 - - [23/Dec/2021:06:55:20 -0600] \"GET /assets/front/assets/vendor/bootstrap/css/bootstrap.min.css HTTP/2\" 20x0 114625", nil, ErrInvalidStatus},
		{"127.0.0.1 - rob [10/Dec/2000:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" abcd 200$", nil, ErrInvalidStatus},
		{"127.255.0.123 - rob [10/Oct/2222:13:55:36 -0250] \"GET /apache_pb.gif HTTP/1.0\" #200 2326", nil, ErrInvalidStatus},
		{"127.1.1.1 - - [10/Dec/2000:13:55:36 -0800] \"GET /apache_pb.gif HTTP/1.0\" 20222206423688288883 2326", nil, ErrInvalidStatus},
	}
	t.Log("===========================================")
	t.Log("Check invalid Status")
	t.Log("-------------------------------------------")
	for _, v := range tests {
		t.Log("checking ...", v.record)
		_, err := ParseLogRecord(v.record)
		if !errors.Is(err, v.err) {
			t.Errorf("FAIL: Status not marked as invalid")
		}
	}

	tests = []struct {
		record string
		log    interface{}
		err    error
	}{
		{"127.0.110.1  - rob [10/Oct/2000:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2a326", nil, ErrInvalidSize},
		{"40.77.189.89 - - [23/Dec/2021:06:55:20 -0600] \"GET /assets/front/assets/vendor/bootstrap/css/bootstrap.min.css HTTP/2\" 200 x114625", nil, ErrInvalidSize},
		{"127.0.0.1 - rob [10/Dec/2000:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326#", nil, ErrInvalidSize},
		{"127.255.0.123 - rob [10/Oct/2222:13:55:36 -0250] \"GET /apache_pb.gif HTTP/1.0\" 200 abcde", nil, ErrInvalidSize},
		{"127.1.1.1 xyz - [10/Dec/2000:13:55:36 -0800] \"GET /apache_pb.gif HTTP/1.0\" 200 $2326", nil, ErrInvalidSize},
	}
	t.Log("===========================================")
	t.Log("Check invalid Size")
	t.Log("-------------------------------------------")
	for _, v := range tests {
		t.Log("checking ...", v.record)
		_, err := ParseLogRecord(v.record)
		if !errors.Is(err, v.err) {
			t.Errorf("FAIL: Size not marked as invalid")
		}
	}

	testCommonLog := []struct {
		record    []string
		IP        string
		Identity  string
		User      string
		Timestamp string
		Request   request
		Status    int
		Size      int64
		err       error
	}{
		{[]string{"127.0.0.1", "-", "rob", "[10/Oct/2000:13:55:36 -0700]", "GET /apache_pb.gif HTTP/1.0", "200", "2326"},
			"127.0.0.1", "-", "rob", "2000-10-10 13:55:36 -0700 PDT", request{"GET", "/apache_pb.gif", "HTTP/1.0"}, 200, 2326, nil},
		{[]string{"154.53.43.164", "-", "-", "[23/Dec/2021:14:41:47 -0600]", "GET /new/installer.php HTTP/1.1", "301", "707"},
			"154.53.43.164", "-", "-", "2021-12-23 14:41:47 -0600 -0600", request{"GET", "/new/installer.php", "HTTP/1.1"}, 301, 707, nil},
	}
	t.Log("===========================================")
	t.Log("Check Common Log")
	t.Log("-------------------------------------------")
	for _, v := range testCommonLog {
		t.Log("checking ...", v.record)
		log, err := ParseLogRecord(strings.Join(v.record, " "))
		if !errors.Is(err, v.err) {
			t.Errorf("FAIL: Error is not nil")
			return
		}
		if log == nil {
			t.Errorf("FAIL: log not returned")
			return
		}

		l := log.(*CommonLog)

		g, w := string(l.IP.String()), v.IP
		compareGotWant(t, "IP", g, w)

		g, w = l.Identity, v.Identity
		compareGotWant(t, "Identity", g, w)

		g, w = l.User, v.User
		compareGotWant(t, "User", g, w)

		g, w = string(l.Timestamp.String()), v.Timestamp
		compareGotWant(t, "Timestamp", g, w)

		g_struct, w_struct := l.Request, v.Request
		compareGotWant(t, "Request", g_struct, w_struct)

		g_int, w_int := l.Status, v.Status
		compareGotWant(t, "Status", g_int, w_int)

		g_int64, w_int64 := l.Size, v.Size
		compareGotWant(t, "Size", g_int64, w_int64)

	}
}

func compareGotWant(t *testing.T, f string, g interface{}, w interface{}) {
	if g != w {
		t.Errorf("FAIL: correct %v not returned, got: %v, want: %v", f, g, w)
	}
}

func isNil(v interface{}) bool {
	return v == nil || (reflect.ValueOf(v).Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil())
}
