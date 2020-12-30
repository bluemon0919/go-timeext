package timeext_test

import (
	"log"
	"testing"
	"time"

	"github.com/bluemon0919/go-timeext"
)

const TestLayout30 = "2006.01.02 27:04:05"
const TestLayout24 = "2006.01.02 15:04:05"

func TestParse13(t *testing.T) {
	// 13:00:00
	actual, err := timeext.Parse(TestLayout30, "2020.07.12 13:00:00")

	if err != nil {
		log.Fatal(err)
	}

	expect, err := time.Parse(TestLayout24, "2020.07.12 13:00:00")
	if err != nil {
		log.Fatal(err)
	}
	if expect != actual {
		log.Fatal(actual)
	}
}

func TestParse24(t *testing.T) {
	// 24:00:00
	actual, err := timeext.Parse(TestLayout30, "2020.07.12 24:00:00")

	if err != nil {
		log.Fatal(err)
	}

	expect, err := time.Parse(TestLayout24, "2020.07.13 00:00:00")
	if err != nil {
		log.Fatal(err)
	}
	if expect != actual {
		log.Fatal(actual)
	}
}

func TestParse241023(t *testing.T) {
	// 24:10:23
	actual, err := timeext.Parse(TestLayout30, "2020.07.12 24:10:23")

	if err != nil {
		log.Fatal(err)
	}

	expect, err := time.Parse(TestLayout24, "2020.07.13 00:10:23")
	if err != nil {
		log.Fatal(err)
	}
	if expect != actual {
		log.Fatal(actual)
	}
}

func TestParse25(t *testing.T) {
	// 25:00:00
	actual, err := timeext.Parse(TestLayout30, "2020.07.12 25:00:00")

	if err != nil {
		log.Fatal(err)
	}

	expect, err := time.Parse(TestLayout24, "2020.07.13 01:00:00")
	if err != nil {
		log.Fatal(err)
	}
	if expect != actual {
		log.Fatal(actual)
	}
}

func TestParse251023(t *testing.T) {
	// 25:10:23
	actual, err := timeext.Parse(TestLayout30, "2020.07.12 25:10:23")

	if err != nil {
		log.Fatal(err)
	}

	expect, err := time.Parse(TestLayout24, "2020.07.13 01:10:23")
	if err != nil {
		log.Fatal(err)
	}
	if expect != actual {
		log.Fatal(actual)
	}
}

func TestParse295959(t *testing.T) {
	// 29:59:59
	actual, err := timeext.Parse(TestLayout30, "2020.07.12 29:59:59")

	if err != nil {
		log.Fatal(err)
	}

	expect, err := time.Parse(TestLayout24, "2020.07.13 05:59:59")
	if err != nil {
		log.Fatal(err)
	}
	if expect != actual {
		log.Fatal(actual)
	}
}

func TestParse30(t *testing.T) {
	// 30:00:00
	_, err := timeext.Parse(TestLayout30, "2020.07.12 30:00:00")

	if err == nil {
		log.Fatal(err)
	}
}

func TestParseErrLayout(t *testing.T) {
	_, err := timeext.Parse(TestLayout24, "2020.07.12 25:00:00")
	if err == nil {
		log.Fatal(err)
	}
}

func TestParseErrLenMismatch(t *testing.T) {
	// 長さが異なる場合はエラーとする
	_, err := timeext.Parse(TestLayout30, "2020.07.12 24:00:00.00")

	if err == nil {
		log.Fatal(err)
	}
}

func TestErr(t *testing.T) {
	// 変換できない文字列の場合はエラーとする
	_, err := timeext.Parse(TestLayout30, "2020.07.12 AA:00:00")

	if err == nil {
		log.Fatal(err)
	}
}

func TestParse13InUTC(t *testing.T) {
	loc, _ := time.LoadLocation("America/New_York")

	// 13:00:00
	actual, err := timeext.ParseInLocation(TestLayout30, "2020.07.12 13:00:00", loc)

	if err != nil {
		log.Fatal(err)
	}

	expect, err := time.ParseInLocation(TestLayout24, "2020.07.12 13:00:00", loc)
	if err != nil {
		log.Fatal(err)
	}
	if expect != actual {
		log.Fatal(actual)
	}
}

func TestParse24InUTC(t *testing.T) {
	loc, _ := time.LoadLocation("America/New_York")

	// 24:00:00
	actual, err := timeext.ParseInLocation(TestLayout30, "2020.07.12 24:00:00", loc)

	if err != nil {
		log.Fatal(err)
	}

	expect, err := time.ParseInLocation(TestLayout24, "2020.07.13 00:00:00", loc)
	if err != nil {
		log.Fatal(err)
	}
	if expect != actual {
		log.Fatal(actual)
	}
}

func TestParse241023InUTC(t *testing.T) {
	loc, _ := time.LoadLocation("America/New_York")

	// 24:10:23
	actual, err := timeext.ParseInLocation(TestLayout30, "2020.07.12 24:10:23", loc)

	if err != nil {
		log.Fatal(err)
	}

	expect, err := time.ParseInLocation(TestLayout24, "2020.07.13 00:10:23", loc)
	if err != nil {
		log.Fatal(err)
	}
	if expect != actual {
		log.Fatal(actual)
	}
}

func TestParse25InUTC(t *testing.T) {
	loc, _ := time.LoadLocation("America/New_York")

	// 25:00:00
	actual, err := timeext.ParseInLocation(TestLayout30, "2020.07.12 25:00:00", loc)

	if err != nil {
		log.Fatal(err)
	}

	expect, err := time.ParseInLocation(TestLayout24, "2020.07.13 01:00:00", loc)
	if err != nil {
		log.Fatal(err)
	}
	if expect != actual {
		log.Fatal(actual)
	}
}

func TestParse251023InUTC(t *testing.T) {
	loc, _ := time.LoadLocation("America/New_York")

	// 25:10:23
	actual, err := timeext.ParseInLocation(TestLayout30, "2020.07.12 25:10:23", loc)

	if err != nil {
		log.Fatal(err)
	}

	expect, err := time.ParseInLocation(TestLayout24, "2020.07.13 01:10:23", loc)
	if err != nil {
		log.Fatal(err)
	}
	if expect != actual {
		log.Fatal(actual)
	}
}

func TestParse295959InUTC(t *testing.T) {
	loc, _ := time.LoadLocation("America/New_York")

	// 29:59:59
	actual, err := timeext.ParseInLocation(TestLayout30, "2020.07.12 29:59:59", loc)

	if err != nil {
		log.Fatal(err)
	}

	expect, err := time.ParseInLocation(TestLayout24, "2020.07.13 05:59:59", loc)
	if err != nil {
		log.Fatal(err)
	}
	if expect != actual {
		log.Fatal(actual)
	}
}

func TestParse30InUTC(t *testing.T) {
	loc, _ := time.LoadLocation("America/New_York")

	// 30:00:00
	_, err := timeext.ParseInLocation(TestLayout30, "2020.07.12 30:00:00", loc)

	if err == nil {
		log.Fatal(err)
	}
}

func TestParseErrLenMismatchInUTC(t *testing.T) {
	loc, _ := time.LoadLocation("America/New_York")

	// 長さが異なる場合はエラーとする
	_, err := timeext.ParseInLocation(TestLayout30, "2020.07.12 24:00:00.00", loc)

	if err == nil {
		log.Fatal(err)
	}
}

func TestErrInUTC(t *testing.T) {
	loc, _ := time.LoadLocation("America/New_York")

	// 変換できない文字列の場合はエラーとする
	_, err := timeext.ParseInLocation(TestLayout30, "2020.07.12 AA:00:00", loc)

	if err == nil {
		log.Fatal(err)
	}
}

func TestFormat13(t *testing.T) {
	tim, _ := time.Parse(TestLayout24, "2020.07.12 13:00:00")
	actual := timeext.TimeExt(tim).Format(TestLayout30)

	if "2020.07.12 13:00:00" != actual {
		log.Fatal(actual)
	}
}

func TestFormat24(t *testing.T) {
	tim, _ := time.Parse(TestLayout24, "2020.07.12 00:00:00")
	actual := timeext.TimeExt(tim).Format(TestLayout30)
	if "2020.07.11 24:00:00" != actual {
		log.Fatal(actual)
	}

	actual = timeext.TimeExt(tim).Format(TestLayout24)
	if "2020.07.12 00:00:00" != actual {
		log.Fatal(actual)
	}
}

func TestFormat241023(t *testing.T) {
	tim, _ := time.Parse(TestLayout24, "2020.07.12 00:10:23")
	actual := timeext.TimeExt(tim).Format(TestLayout30)
	if "2020.07.11 24:10:23" != actual {
		log.Fatal(actual)
	}

	actual = timeext.TimeExt(tim).Format(TestLayout24)
	if "2020.07.12 00:10:23" != actual {
		log.Fatal(actual)
	}
}

func TestFormat25(t *testing.T) {
	tim, _ := time.Parse(TestLayout24, "2020.07.12 01:00:00")
	actual := timeext.TimeExt(tim).Format(TestLayout30)
	if "2020.07.11 25:00:00" != actual {
		log.Fatal(actual)
	}

	actual = timeext.TimeExt(tim).Format(TestLayout24)
	if "2020.07.12 01:00:00" != actual {
		log.Fatal(actual)
	}
}

func TestFormat251023(t *testing.T) {
	tim, _ := time.Parse(TestLayout24, "2020.07.12 01:10:23")
	actual := timeext.TimeExt(tim).Format(TestLayout30)
	if "2020.07.11 25:10:23" != actual {
		log.Fatal(actual)
	}

	actual = timeext.TimeExt(tim).Format(TestLayout24)
	if "2020.07.12 01:10:23" != actual {
		log.Fatal(actual)
	}
}

func TestFormat295959(t *testing.T) {
	tim, _ := time.Parse(TestLayout24, "2020.07.12 05:59:59")
	actual := timeext.TimeExt(tim).Format(TestLayout30)
	if "2020.07.11 29:59:59" != actual {
		log.Fatal(actual)
	}

	actual = timeext.TimeExt(tim).Format(TestLayout24)
	if "2020.07.12 05:59:59" != actual {
		log.Fatal(actual)
	}
}

func TestFormat30(t *testing.T) {
	tim, _ := time.Parse(TestLayout24, "2020.07.12 06:00:00")
	actual := timeext.TimeExt(tim).Format(TestLayout30)
	if "2020.07.12 06:00:00" != actual {
		log.Fatal(actual)
	}

	actual = timeext.TimeExt(tim).Format(TestLayout24)
	if "2020.07.12 06:00:00" != actual {
		log.Fatal(actual)
	}
}

func TestFormatUnsupportedFormat(t *testing.T) {
	tim, _ := time.Parse(TestLayout24, "2020.07.11 00:00:00")
	actual := timeext.TimeExt(tim).Format("2006.01.02 AA:04:05")
	if "2020.07.11 AA:00:00" != actual {
		log.Fatal(actual)
	}
}

func TestFormat13InUTC(t *testing.T) {
	loc, _ := time.LoadLocation("America/New_York")
	tim, _ := time.ParseInLocation(TestLayout24, "2020.07.12 13:00:00", loc)
	actual := timeext.TimeExt(tim).Format(TestLayout30)
	if "2020.07.12 13:00:00" != actual {
		log.Fatal(actual)
	}

	actual = timeext.TimeExt(tim).Format(TestLayout24)
	if "2020.07.12 13:00:00" != actual {
		log.Fatal(actual)
	}
}

func TestFormat24InUTC(t *testing.T) {
	loc, _ := time.LoadLocation("America/New_York")
	tim, _ := time.ParseInLocation(TestLayout24, "2020.07.12 00:00:00", loc)
	actual := timeext.TimeExt(tim).Format(TestLayout30)
	if "2020.07.11 24:00:00" != actual {
		log.Fatal(actual)
	}

	actual = timeext.TimeExt(tim).Format(TestLayout24)
	if "2020.07.12 00:00:00" != actual {
		log.Fatal(actual)
	}
}
