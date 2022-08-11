package stuff

import (
	"errors"
	"strconv"
	"time"
)

var Name string = "Derek"

// Function name is uppercase so it can be exported
func IntArrToStrArr(intArr []int) []string {
	var strArr []string
	for _, i := range intArr {
		strArr = append(strArr, strconv.Itoa(i))
	}
	return strArr
}

// We capitalize what we want to export and use
// lowercase on what we don't. We don't want the
// user to be able to access day, month, year
// directly
type Date struct {
	day   int
	month int
	year  int
}

// Create a setter function for the values
// Make sure all values are valid or return
// an error message
func (d *Date) SetDay(day int) error {
	if (day < 1) || (day > 31) {
		return errors.New("incorrect day value")
	}
	d.day = day
	return nil
}
func (d *Date) SetMonth(m int) error {
	if (m < 1) || (m > 12) {
		return errors.New("incorrect month value")
	}
	d.month = m
	return nil
}
func (d *Date) SetYear(y int) error {
	if (y < 1875) || (y > time.Now().Year()) {
		return errors.New("incorrect year value")
	}
	d.year = y
	return nil
}

// Getter functions return the values
func (d *Date) Day() int {
	return d.day
}
func (d *Date) Month() int {
	return d.month
}
func (d *Date) Year() int {
	return d.year
}
