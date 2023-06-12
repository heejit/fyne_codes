package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"

)

const CONST_DEFAULT_DATE_FORMAT = "02-Jan-2006"

var DAY_POS   = []int{0, 2}
var MONTH_POS = []int{3, 6}
var YEAR_POS  = []int{7, 11}

type MyDateEntry struct {
	widget.Entry
	current_date_value time.Time
}

func NewMyDateEntry() *MyDateEntry {
	entry := new(MyDateEntry)
	entry.ExtendBaseWidget(entry)
	entry.SetPlaceHolder("DD-MMM-YYYY")
	return entry
}

// Get current cursor position to Section (Day Month year)
func (e *MyDateEntry) cursorPosToSection() string {
	if e.CursorColumn >= DAY_POS[0] && e.CursorColumn <= DAY_POS[1] {
		return "d"
	}
	if e.CursorColumn >= MONTH_POS[0] && e.CursorColumn <= MONTH_POS[1] {
		return "m"
	}
	if e.CursorColumn >= YEAR_POS[0] && e.CursorColumn <= YEAR_POS[1] {
		return "y"
	}
	return ""
}

// Get Day Month year to cursor postion
func (e *MyDateEntry) sectionToCursorPos(lsec string) int {
	if lsec == "d" {
		return DAY_POS[0]
	}

	if lsec == "m" {
		return MONTH_POS[0]
	}

	if lsec == "y" {
		return YEAR_POS[0]
	}
	return -1
}

// add year, month, day to current date
func (e *MyDateEntry) addTime(v int, cur_section string) {
	if e.current_date_value.IsZero() == true {
		return
	}

	if cur_section == "d" {
		e.current_date_value = e.current_date_value.AddDate(0, 0, v)
	}

	if cur_section == "m" {
		e.current_date_value = e.current_date_value.AddDate(0, v, 0)
	}

	if cur_section == "y" {
		e.current_date_value = e.current_date_value.AddDate(v, 0, 0)
	}
	e.updateDisplay()
}

// set current date on space key
func (e *MyDateEntry) setCurrentDate() {
	e.current_date_value = time.Now()
	e.updateDisplay()
}

// update current display
func (e *MyDateEntry) updateDisplay() {
	e.SetText(e.current_date_value.Format(CONST_DEFAULT_DATE_FORMAT))
}

// handle key events
func (e *MyDateEntry) TypedKey(key *fyne.KeyEvent) {

	if key.Name == fyne.KeyDelete {
		e.SetText("")
		return
	}

	if key.Name == fyne.KeyUp {
		e.addTime(1, e.cursorPosToSection())
		return
	}

	if key.Name == fyne.KeyDown {
		e.addTime(-1, e.cursorPosToSection())
		return
	}

	if key.Name == fyne.KeySpace {
		e.setCurrentDate()
		return
	}

	if key.Name == fyne.KeyEnter {
		e.parse_and_update_date()
		e.addTime(0, e.cursorPosToSection())
		return
	}

	if key.Name == fyne.KeyReturn {
		e.parse_and_update_date()
		e.addTime(0, e.cursorPosToSection())
		return
	}

	e.Entry.TypedKey(key)
}

// this where we are converting current text to date
func (e *MyDateEntry) FocusLost() {
	e.parse_and_update_date()
	e.Entry.FocusLost()
}

// Date string to time.Time conversion
// we assume 1st part is always Day
// input = 1 -> 1-CurMonth-CurYear
// input = 1.5, 1/5, 1-5 -> 1-5-CurYear
func (e *MyDateEntry) parse_and_update_date()  {
	var date_str = e.Text

	e.TextStyle.Bold = false

	if len(date_str) == 0 {
		e.SetText("")
		return
	}

	var y, m int

	y = time.Now().Year()
	m = int(time.Now().Month())

	date_str = strings.Replace(date_str, ".", "-", -1)
	date_str = strings.Replace(date_str, "/", "-", -1)
	dt := strings.Split(date_str, "-")

	if len(dt) == 1 {
		dt[0] = strings.TrimSpace(dt[0])
		date_str = fmt.Sprintf("%s-%d-%d", dt[0], m, y)
	}

	if len(dt) == 2 {
		dt[0] = strings.TrimSpace(dt[0])
		dt[1] = strings.TrimSpace(dt[1])

		if len(dt[1]) == 0 {
			dt[1] = strconv.Itoa(m)
		}

		date_str = fmt.Sprintf("%s-%s-%d", dt[0], dt[1], y)
	}

	date_str = strings.TrimSpace(date_str)

	var allowed_formats = []string{"02-01-2006", "2-1-2006", "2006-01-02", "2006-1-2", "2-Jan-2006"}
	for _, v := range allowed_formats {
		e.current_date_value, _ = time.Parse(v, date_str)
		if e.current_date_value.IsZero() == false {
			break
		}
	}

	if e.current_date_value.IsZero() == true {
		e.SetText("")
	}  else {
		e.SetText(e.current_date_value.Format(CONST_DEFAULT_DATE_FORMAT))
		e.TextStyle.Bold = true
	}
}

// return current date
func (e *MyDateEntry) ToDate() time.Time {
	return e.current_date_value
}
