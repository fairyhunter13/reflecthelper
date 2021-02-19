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

func parseTime(timeTxt string, option *Option) (result time.Time, err error) {
	layouts := append(option.TimeLayouts, listLayouts...)
	for _, layout := range layouts {
		result, err = time.Parse(layout, timeTxt)
		if err == nil {
			break
		}
	}
	return
}
