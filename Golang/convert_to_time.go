package main

import (
	"math"
	"strconv"
)

// This function helps to convert float64 number to hours in string type
func ConvertToTime(text float64) (resp string) {
	var hour, minute, second float64                                                 // There was declared variables of hour, minute and second
	var hourString, minuteString, secondString, hourRes, minuteRes, secondRes string // This variables used to act variables given above

	hour = math.Floor(text)                                      // math.Floor used for taking decimal value of number
	minute = math.Floor((text - float64(hour)) * 60)             // this is the formula how to find minute
	second = math.Floor(((text-float64(hour))*60 - minute) * 60) // this is the formula how to find second
	h := int(hour)                                               // h -> hour which converted to int format
	m := int(minute)                                             // m -> minute which converted to int format
	s := int(second)                                             // s -> second which converted to int format
	hourString = strconv.Itoa(h)                                 // now we converted hour to string from int, because strconv.Itoa converts only int to string
	minuteString = strconv.Itoa(m)                               // now we converted minute to string from int, because strconv.Itoa converts only int to string
	secondString = strconv.Itoa(s)                               // now we converted second to string from int, because strconv.Itos convets only int to string

	if len(hourString) == 1 { // found length of hour, it could be maximux 2. For example 14 or 2 etc.
		hourRes = "0" + strconv.Itoa(int(hour)) // if length of hour is 1, we put 0 before hour. 1 -> 01
	} else {
		hourRes = strconv.Itoa(int(hour)) // otherwise we don't change hour 11 -> 11
	}
	if len(minuteString) == 1 {
		minuteRes = "0" + strconv.Itoa(int(minute)) // if length of minute is 1, we put 0 before minute. 1 -> 01
	} else {
		minuteRes = strconv.Itoa(int(minute)) // otherwise we don't change minute 11 -> 11
	}
	if len(secondString) == 1 {
		secondRes = "0" + strconv.Itoa(int(second)) // if length of second is 2, we put 0 before second. 1 -> 01
	} else {
		secondRes = strconv.Itoa(int(second)) // otherwise we don't change second 11 -> 11
	}

	resp = hourRes + ":" + minuteRes + ":" + secondRes // we put colons between hour, minut, second in order to present them in time format 13:21:12
	return resp
}
