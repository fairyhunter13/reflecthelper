package reflecthelper

import "time"

var (
	listLayouts = []string{
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	}
)

// ParseTime parses the timeTxt string to the time.Time using various formats.
func ParseTime(timeTxt string) (result time.Time, err error) {
	for _, layout := range listLayouts {
		result, err = time.Parse(layout, timeTxt)
		if err == nil {
			break
		}
	}
	return
}
