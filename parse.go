package apachelogparser

import (
	"errors"
	"fmt"
	"net"
	"strings"
	"time"
)

const (
	sizeCommonLog         = 10
	sizeCombinedLog       = 12
	StandardEnglishFormat = "02/Jan/2006:15:04:05 -0700"
)

var (
	ErrInvalidLog            = errors.New("Log record is neither Apache common nor combined log")
	ErrInvalidIP             = errors.New("IP invalid")
	ErrInvalidFieldTimestamp = errors.New("Timestamp field invalid, missing either opening '[' or closing ']' or both")
)

func ParseLogRecord(r string) (interface{}, error) {

	var log *CommonLog
	var err error

	s := strings.Fields(r)
	l := len(s)

	switch l {
	case sizeCommonLog, sizeCombinedLog:

		log, err = getCommonFields(s)
		if err != nil {
			return nil, err
		}

	/* case sizeCommonLog:

		*log = CommonLog{
			IP:        net.ParseIP(s[IP]),
			User:      s[User],
			Identity:  s[Identity],
			Timestamp: s[Timestamp],
			Request:   s[Request],
			Status:    s[Status],
			Size:      s[Size],
		}

	case sizeCombinedLog:

		*log = CommonLog{
			IP:        net.ParseIP(s[IP]),
			User:      s[User],
			Identity:  s[Identity],
			Timestamp: s[Timestamp],
			Request:   s[Request],
			Status:    s[Status],
			Size:      s[Size],
			Referer:   s[Request],
			UserAgent: s[UserAgent],
		}
	*/
	default:
		return nil, fmt.Errorf("validate log size, %w", ErrInvalidLog)
	}

	return log, nil

}

func getCommonFields(s []string) (*CommonLog, error) {

	var log *CommonLog

	ip, err := getIP(s[IP])
	if err != nil {
		return nil, err
	}

	timestamp, err := getDateTime(s[Timestamp])
	if err != nil {
		return nil, err
	}

	log.IP = ip
	log.Timestamp = timestamp

	return log, nil
}

func getIP(input string) (net.IP, error) {
	ip := net.ParseIP(input)
	if ip == nil {
		return nil, fmt.Errorf("validate IP, %w", ErrInvalidIP)
	}
	return ip, nil
}

func getDateTime(input string) (timestamp time.Time, err error) {

	if input[0] != '[' {
		//s := fmt.Sprintf("got %q, want '['", input[0])
		err = fmt.Errorf("missing opening '[' %w", ErrInvalidFieldTimestamp)
		return
	}
	idx := strings.Index(input, "]")
	if idx == -1 {
		err = fmt.Errorf("missing closing ']', %w", ErrInvalidFieldTimestamp)
		return
	}
	if timestamp, err = time.Parse(StandardEnglishFormat, input[1:idx]); err != nil {
		err = errors.New("failed to parse datetime: " + err.Error())
	}
	return
}
