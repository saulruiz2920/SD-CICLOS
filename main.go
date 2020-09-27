package main

import "fmt"

type zodiac struct {
    name string
    start_month int
    start_day int
}

func initZodiacs() []zodiac {
  zodiacs := []zodiac {
    {
      name: "aries",
      start_month: 3,
      start_day: 21,
    },
    {
      name: "tauro",
      start_month: 4,
      start_day: 21,
    },
    {
      name: "geminis",
      start_month: 5,
      start_day: 21,
    },
    {
      name: "cancer",
      start_month: 6,
      start_day: 22,
    },
    {
      name: "leo",
      start_month: 7,
      start_day: 23,
    },
    {
      name: "virgo",
      start_month: 8,
      start_day: 23,
    },
    {
      name: "libra",
      start_month: 9,
      start_day: 23,
    },
    {
      name: "escorpio",
      start_month: 10,
      start_day: 23,
    },
    {
      name: "sagitario",
      start_month: 11,
      start_day: 23,
    },
    {
      name: "capricornio",
      start_month: 12,
      start_day: 22,
    },
    {
      name: "acuario",
      start_month: 1,
      start_day: 21,
    },
    {
      name: "piscis",
      start_month: 2,
      start_day: 19,
    },
  }
  return zodiacs
}
func isDayAndMonthValid(day int, month int) bool {
  if day < 1 || day > 31  || month < 1 || month > 12{
    return false
  }
  return true
}

func getZodiac(day int, month int) string {
  zodiacs := initZodiacs()
  for idx := range zodiacs {
    if zodiacs[idx].start_month == month {
      if day >= zodiacs[idx].start_day {
        return zodiacs[idx].name
      } else {
        if idx-1 < 0 {
          idx = len(zodiacs)
        }
        return zodiacs[idx-1].name
      }
    }
  }
  return ""
}

func main() {
  var day int
  var month int
  fmt.Scan(&day)
  fmt.Scan(&month)
  if isDayAndMonthValid(day, month) {
    fmt.Println(getZodiac(day, month))
  }

}