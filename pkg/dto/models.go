package dto

import (
	"strings"
	"time"
)

type Date time.Time

func (c *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*c = Date(t)
	return nil
}

func (c *Date) String() string {
	x := time.Time(*c)
	return x.Format("02.01.2006")
}
