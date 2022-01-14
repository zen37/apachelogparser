package apachelogparser

import (
	"net"
	"time"
)

const (
	StandardEnglishFormat string = "02/Jan/2006:15:04:05 -0700"
	sizeCommonLog                = 10
	sizeCombinedLog              = 7
	separator                    = " "
)

//position of the log entry
const (
	IP = iota
	Identity
	User
	Timestamp
	TZ
	Method
	Resource
	Protocol
	Status
	Size
	Referer
	UserAgent
)

// Example "GET /apache_pb.gif HTTP/1.0"
type request struct {
	Method   string
	Resource string //path
	Protocol string
}

// LogFormat "%h %l %u %t \"%r\" %>s %b" common
type CommonLog struct {
	IP        net.IP
	Identity  string
	User      string
	Timestamp time.Time
	Request   request
	Status    int
	Size      int64
}

// LogFormat "%h %l %u %t \"%r\" %>s %b \"%{Referer}i\" \"%{User-agent}i\"" combined
type CombinedLog struct {
	Common    CommonLog
	Referer   string
	UserAgent string
}
