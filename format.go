package timeext

import (
	"errors"
	"strconv"
	"time"
)

// layoutで使用される時間は特定の時間です。
// timeパッケージではHour=15とすると24時間制、Hour＝27とすると30時間制で表現します。
//  2006.01.02 27:04:05
const (
	stdHour = iota + 1
	std30Hour
)

var errBad = errors.New("bad value for field")
var errLenMismatch = errors.New("Length mismatch")

// Parse はtimeパッケージのParseの拡張ラッパーです。
// timeパッケージが処理できる場合は、その結果をそのまま返します。
// timeパッケージが処理できない30時間制の時間表現を解析し、時間値を返します。
// 30時間制は24:00:00-29:59:59の時間表現を指します。
// layoutは、基準時間をどのように定義するかを示すことでフォーマットを定義します。
func Parse(layout, value string) (TimeExt, error) {
	t, err := time.Parse(layout, value)
	if err != nil {
		t, err = parseIn30HourSystem(layout, value, time.UTC)
		if err != nil {
			return TimeExt(t), err
		}
		return TimeExt(t), nil
	}
	return TimeExt(t), nil
}

// ParseInLocation はtimeパッケージのParseInLocationの拡張ラッパーです。
// Parseは時刻をUTCとして解釈するのに対し、ParseInLocationは指定された場所と同様に時間を解釈します。
func ParseInLocation(layout, value string, loc *time.Location) (TimeExt, error) {
	t, err := time.ParseInLocation(layout, value, loc)
	if err != nil {
		t, err = parseIn30HourSystem(layout, value, loc)
		if err != nil {
			return TimeExt(t), err
		}
		return TimeExt(t), nil
	}
	return TimeExt(t), nil
}

// IsExt は30時間制で拡張処理するかどうかを返します。
// timeパッケージが処理できる場合は、falseを返します。
// timeパッケージが処理できない30時間制の時間表現の場合はtrueを返します。
func IsExt(layout, value string) (bool, error) {
	_, err := time.Parse(layout, value)
	if err != nil {
		_, err = parseIn30HourSystem(layout, value, time.UTC)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

func parseIn30HourSystem(layout, value string, loc *time.Location) (time.Time, error) {
	// 同じ文字列長になるフォーマットしか対応していない
	if len(layout) != len(value) {
		return time.Time{}, errLenMismatch
	}
	prefix, std, suffix := nextStdChunk(layout)

	switch std {
	case 0, stdHour:
		return time.Time{}, errBad
	}
	valueS := value[len(prefix) : len(layout)-len(suffix)]
	hour, err := strconv.Atoi(valueS)
	if err != nil {
		return time.Time{}, errBad
	}
	offsetDay := 0              // 24:00-29:59の場合は1日加算する
	if hour < 6 || hour >= 30 { // ６時前、３０時以降はエラー
		return time.Time{}, errBad
	}
	if hour >= 24 {
		hour -= 24
		offsetDay = 1
	}
	valueE := value[0:len(prefix)] + strconv.Itoa(hour) + value[len(layout)-len(suffix):]

	timeLayout := prefix + "15" + suffix

	var result time.Time
	if loc == time.UTC {
		result, err = time.Parse(timeLayout, valueE)
	} else {
		result, err = time.ParseInLocation(timeLayout, valueE, loc)
	}
	if err != nil {
		return time.Time{}, err
	}
	return result.AddDate(0, 0, offsetDay), nil
}

func nextStdChunk(layout string) (prefix string, std int, suffix string) {
	for i := 0; i < len(layout); i++ {
		switch c := int(layout[i]); c {
		case '1': // 15
			if len(layout) >= i+2 && layout[i+1] == '5' {
				return layout[0:i], stdHour, layout[i+2:]
			}
		case '2': // 27
			if len(layout) >= i+2 && layout[i+1] == '7' {
				return layout[0:i], std30Hour, layout[i+2:]
			}
		}
	}
	return layout, 0, ""
}

// TimeExt はtimeパッケージのTimeを拡張します。
// Formatなどを使って30時間制の時間表現を利用できます。
type TimeExt time.Time

// Format はtimeパッケージのFormatの拡張ラッパーです。
// 30時間制の時間表現でフォーマットした文字列を返します。
// layoutは、基準時間をどのように定義するかを示すことでフォーマットを定義します。
func (t TimeExt) Format(layout string) string {
	prefix, std, suffix := nextStdChunk(layout)
	switch std {
	case 0, stdHour:
		return time.Time(t).Format(layout)
	}

	timeLayout := prefix + "15" + suffix

	origin := time.Time(t).Format(timeLayout)
	value := time.Time(t).AddDate(0, 0, -1).Format(timeLayout)
	// 同じ文字列長になるフォーマットしか対応していない
	if len(layout) != len(value) {
		return origin
	}

	valueS := value[len(prefix) : len(layout)-len(suffix)]
	hour, err := strconv.Atoi(valueS)
	if err != nil {
		return origin
	}
	if 6 <= hour {
		return origin
	}
	hour += 24
	valueE := value[0:len(prefix)] + strconv.Itoa(hour) + value[len(layout)-len(suffix):]
	return valueE
}
