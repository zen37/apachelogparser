package apachelogparser

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

const (
	sizeCommonLog   = 10
	sizeCombinedLog = 12
)

var (
	ErrInvalidLog = errors.New("Log record is neither Apache common nor combined log")
	ErrInvalidIP  = errors.New("IP invalid")
)

func ParseLogRecord(r string) (*Log, error) {

	var log *Log

	s := strings.Fields(r)
	l := len(s)

	switch l {
	case sizeCommonLog, sizeCombinedLog:

		err := checkFields(s)
		if err != nil {
			return nil, err
		}

		*log = CommonLog{
			IP: net.ParseIP(s[IP]),
		}

	default:
		return nil, fmt.Errorf("validate log size, %w", ErrInvalidLog)
	}

	return log, nil

}

func checkFields(s []string) error {
	err := checkIP(s[IP])
	if err != nil {
		return err
	}
	return nil
}

func checkIP(v string) error {
	ip := net.ParseIP(v)
	if ip == nil {
		return fmt.Errorf("validate IP, %w", ErrInvalidIP)
	}
	return nil
}
