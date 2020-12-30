# timeext

<img src="https://img.shields.io/github/workflow/status/bluemon0919/go-timeext/Go">

An extension wrapper for the time package that handles 30-hour time representation.

## Install

```bash
$ go get -u github.com/bluemon0919/go-timeext
```

## Example

```go
const Layout string = "2006.01.02 15:04:05"
parsed, _ := timeext.Parse(Layout, "2020.07.12 25:00:00")

s24 := time.Time(parsed).Format(Layout)
fmt.Println(s24) // 2020.07.13 01:00:00

s30 := parsed.Format(Layout)
fmt.Println(s30) // 2020.07.12 25:00:00
```
