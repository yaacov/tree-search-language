package semantics

import "time"

func toDate(value interface{}) (time.Time, bool) {
	switch v := value.(type) {
	case time.Time:
		return v, true
	case string:
		date, err := time.Parse(time.RFC3339, v)
		if err == nil {
			return date, true
		}
	}
	return time.Time{}, false
}
