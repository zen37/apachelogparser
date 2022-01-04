package apachelogparser

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

const (
	sizeCommonLog   = 10
	sizeCombinedLog = 12
	separator       = " "
)

var (
	ErrInvalidLog            = errors.New("log record is neither Apache common nor combined log")
	ErrInvalidIP             = errors.New("IP invalid")
	ErrInvalidFieldTimestamp = errors.New("timestamp field invalid, missing either opening '[' or closing ']' or both")
	ErrInvalidTimestamp      = errors.New("timestamp invalid")
	ErrInvalidStatus         = errors.New("status invalid")
	ErrInvalidSize           = errors.New("size invalid")
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

	var log CommonLog

	ip, err := getIP(s[IP])
	if err != nil {
		return nil, err
	}
	// second position is the UTC, 10/Oct/2000:13:55:36 -0700
	timestamp, err := getDateTime(s[Timestamp] + separator + s[TZ])
	if err != nil {
		return nil, err
	}
	status, err := getStatus(s[Status])
	if err != nil {
		return nil, err
	}

	size, err := getSize(s[Size])
	if err != nil {
		return nil, err
	}

	log.IP = ip
	log.Identity = s[Identity]
	log.User = s[User]
	log.Timestamp = timestamp

	r := request{s[Method], s[Resource], s[Protocol]}
	log.Request = r

	log.Status = status
	log.Size = size

	return &log, nil
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
		err = fmt.Errorf("parsing time error:[%s] %w", err, ErrInvalidTimestamp)
	}
	return
}

func getStatus(input string) (int, error) {
	i, err := strconv.Atoi(input)
	if err != nil {
		err = fmt.Errorf("parsing time error:[%s] %w", err, ErrInvalidStatus)
		return 0, err
	}
	return i, nil
}

func getSize(input string) (int64, error) {
	i, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		err = fmt.Errorf("parsing time error:[%s] %w", err, ErrInvalidSize)
		return 0, err
	}
	return i, nil
}
