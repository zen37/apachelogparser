package apachelogparser

import (
	"fmt"
	"strings"
)

const (
	sizeCommonLog   = 7
	sizeCombinedLog = 9
)

func ParseLogRecord(r string) (*Log, error) {

	s := strings.Fields(r)
	/* 	for _, e := range s {
	   		fmt.Println(e)
	   	}
	   	fmt.Println(s) */
	l := len(s)

	switch l {
	case sizeCommonLog:
		fmt.Println("sizeCommonLog, size:", l)
	case sizeCombinedLog:

	default:
		return nil, fmt.Errorf("Log record is neither Apache common nor combined log")
	}

	return nil, nil

}
