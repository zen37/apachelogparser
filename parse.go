package apachelogparser

import (
	"fmt"
	"net"
	"strings"
)

const (
	sizeCommonLog   = 10
	sizeCombinedLog = 12
)

func ParseLogRecord(r string) (*Log, error) {

	s := strings.Fields(r)
	l := len(s)

	switch l {
	case sizeCommonLog, sizeCombinedLog:

		err := checkFields(s)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("Log record is neither Apache common nor combined log")
	}

	return nil, nil

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
		return fmt.Errorf("IP invalid")
	}
	return nil
}
