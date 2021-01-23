package format

/*
All letters 'A' to 'Z' and 'a' to 'z' are reserved as pattern letters. The following pattern letters are defined:
    Symbol  Meaning                     Presentation      Examples
    ------  -------                     ------------      -------
     G       era                         text              AD; Anno Domini; A
     u       year                        year              2004; 04
     y       year-of-era                 year              2004; 04
     D       day-of-year                 number            189
     M/L     month-of-year               number/text       7; 07; Jul; July; J
     d       day-of-month                number            10

     Q/q     quarter-of-year             number/text       3; 03; Q3; 3rd quarter
     Y       week-based-year             year              1996; 96
     w       week-of-week-based-year     number            27
     W       week-of-month               number            4
     E       day-of-week                 text              Tue; Tuesday; T
     e/c     localized day-of-week       number/text       2; 02; Tue; Tuesday; T
     F       week-of-month               number            3

     a       am-pm-of-day                text              PM
     h       clock-hour-of-am-pm (1-12)  number            12
     K       hour-of-am-pm (0-11)        number            0
     k       clock-hour-of-am-pm (1-24)  number            0

     H       hour-of-day (0-23)          number            0
     m       minute-of-hour              number            30
     s       second-of-minute            number            55
     S       fraction-of-second          fraction          978
     A       milli-of-day                number            1234
     n       nano-of-second              number            987654321
     N       nano-of-day                 number            1234000000

     V       time-zone ID                zone-id           America/Los_Angeles; Z; -08:30
     z       time-zone name              zone-name         Pacific Standard Time; PST
     O       localized zone-offset       offset-O          GMT+8; GMT+08:00; UTC-08:00;
     X       zone-offset 'Z' for zero    offset-X          Z; -08; -0830; -08:30; -083015; -08:30:15;
     x       zone-offset                 offset-x          +0000; -08; -0830; -08:30; -083015; -08:30:15;
     Z       zone-offset                 offset-Z          +0000; -0800; -08:00;

     p       pad next                    pad modifier      1

     '       escape for text             delimiter
     ''      single quote                literal           '
     [       optional section start
     ]       optional section end
     #       reserved for future use
     {       reserved for future use
     }       reserved for future use

*/
import (
	"bytes"
	"time"
)

// java pattern: yyyy-MM-dd HH:mm:ss
var DateTimeFormat = OfPattern("yyyy-MM-dd HH:mm:ss")

// java pattern: yyyy-MM-dd HH:mm:ss.SSS
var DateTimeMSFormat = OfPattern("yyyy-MM-dd HH:mm:ss.SSS")

// java pattern: yyyy-MM-dd HH:mm
var DateTimeMinuteFormat = OfPattern("yyyy-MM-dd HH:mm")

// java pattern: yyyy-MM-dd
var DateFormat = OfPattern("yyyy-MM-dd")

// java pattern: HH:mm:ss
var TimeFormat = OfPattern("HH:mm:ss")

// java pattern: yyyyMMddHHmmss
var DateTimePureFormat = OfPattern("yyyyMMddHHmmss")

// java pattern: yyyyMMddHHmmssSSS
var DateTimeMSPureFormat = OfPattern("yyyyMMddHHmmssSSS")

// java pattern: yyyyMMddHHmm
var DateTimeMinutePureFormat = OfPattern("yyyyMMddHHmm")

// java pattern: yyyyMMdd
var DatePureFormat = OfPattern("yyyyMMdd")

// java pattern: HHmmss
var TimePureFormat = OfPattern("HHmmss")

type DateTimeFormatter struct {
	parsers     []DateTimePrinterParser
	javaPattern string
	goPattern   string
}
type DateTimePrinterParser struct {
	types  string
	value  string
	length int
}

var fieldMap = map[string]string{
	"y":    "2006",
	"yy":   "06",
	"yyyy": "2006",
	"M":    "1",
	"MM":   "01",
	"MMM":  "Jan",
	"MMMM": "January",
	"d":    "2",
	"dd":   "02",
	"h":    "3",
	"hh":   "03",
	"H":    "15",
	"P":    "PM",
	"p":    "pm",
	"A":    "AM",
	"a":    "am",
	"m":    "4",
	"mm":   "04",
	"s":    "5",
	"ss":   "05",
	"S":    "000",
}

func (format *DateTimeFormatter) Format(t time.Time) string {
	return t.Format(format.goPattern)
}
func OfPattern(javaPattern string) *DateTimeFormatter {
	format := &DateTimeFormatter{
		parsers:     make([]DateTimePrinterParser, 10),
		javaPattern: javaPattern,
	}
	var buffer = new(bytes.Buffer)
	size := len(javaPattern)
	isStr := false
	for i := 0; i < size; i++ {
		value := string(javaPattern[i])
		if value == "'" {
			isStr = !isStr
			continue
		}
		if isStr {
			buffer.WriteString(value)
			continue
		}
		field, ok := fieldMap[value]
		if ok {
			checkStr(buffer, format)
			length := 1
			realPattern := value
			lastRealField := field
			for i < len(javaPattern)-1 && string(javaPattern[i+1]) == value {
				i++
				length++
				realPattern = realPattern + value
				if v, ok := fieldMap[realPattern]; ok {
					lastRealField = v
				}
			}
			format.parsers = append(format.parsers, DateTimePrinterParser{
				types:  "1",
				value:  lastRealField,
				length: length,
			})
		} else {
			buffer.WriteString(value)
		}
	}
	if buffer.Len() > 0 {
		parser := DateTimePrinterParser{
			value:  buffer.String(),
			types:  "2",
			length: buffer.Len(),
		}
		format.parsers = append(format.parsers, parser)
		buffer.Reset()
	}
	buffer.Reset()
	for _, parser := range format.parsers {
		buffer.WriteString(parser.value)
	}
	format.goPattern = buffer.String()
	return format
}

func checkStr(buffer *bytes.Buffer, format *DateTimeFormatter) {
	if buffer.Len() > 0 {
		parser := DateTimePrinterParser{
			value:  buffer.String(),
			types:  "2",
			length: buffer.Len(),
		}
		format.parsers = append(format.parsers, parser)
		buffer.Reset()
	}
}

func (format *DateTimeFormatter) GoPattern() string {
	return format.goPattern
}
