package leetcode

// https://leetcode.com/problems/reformat-date/

// For day, we can scan until facing a non number
// Then prefix it with 0 if needed
// For month and year the length is fixed
func reformatDate(date string) string {
	day := ""
	month := ""
	year := ""
	i := 0
	for ; i < 4 && date[i] >= '0' && date[i] <= '9'; i++ {
		day = day + string(date[i])
	}
	if len(day) == 1 {
		day = "0" + day
	}
	i = i + 3
	month = toMonth(date[i : i+3])
	year = date[len(date)-4:]
	return year + "-" + month + "-" + day
}

func toMonth(month string) string {
	switch month {
	case "Jan":
		return "01"
	case "Feb":
		return "02"
	case "Mar":
		return "03"
	case "Apr":
		return "04"
	case "May":
		return "05"
	case "Jun":
		return "06"
	case "Jul":
		return "07"
	case "Aug":
		return "08"
	case "Sep":
		return "09"
	case "Oct":
		return "10"
	case "Nov":
		return "11"
	case "Dec":
		return "12"
	}
	return ""
}
