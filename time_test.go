//

package ru

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	var fmts []string
	for _, c := range "dDjlNzWFfMmnYyaAGHghisuvTZ" {
		fmts = append(fmts, "\\"+string(c)+"="+string(c))
	}

	fmt.Println(DayName(6, Nm), MonthName(7, Nm))
	fmt.Println(TimeFormat(strings.Join(fmts, ", "), time.Now()))

}
