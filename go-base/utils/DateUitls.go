package utils


import (
	"time"
	"fmt"
)

/**
*日期工具类
* autor:zjcjava@163.com
* time:2018-9-10
 */

const (
	DATE         = "2006-01-02"
	SHOR_TDATE   = "06-01-02"
	TIME         = "15:04:05"
	SHORT_TIME   = "15:04"
	DATE_TIME    = "2006-01-02 15:04:05"
	NEW_DATE_TIME = "2006/01/02 15~04~05"
	NEW_TIME      = "15~04~05"
)


/*func main() {
	thisdate := "2014-03-17 14:55:06"

	fmt.Println(Formate(thisdate,DATE_TIME))
	fmt.Println(Formate(thisdate,DATE))
	fmt.Println(Formate(thisdate,DATE))
	fmt.Println(Formate(thisdate,TIME))
	fmt.Println(Formate(thisdate,SHORT_TIME))
	fmt.Println(Formate(thisdate,NEW_DATE_TIME))
	fmt.Println(Formate(thisdate,NEW_TIME))
	fmt.Println(NewDatetime)

}*/

func NewDatetime() string {
	return time.Now().Format(DATE_TIME)
}

func NewDateStamp() int64 {
	return time.Now().Unix()
}

func DateToStamp(datastr string) int64 {
	the_time, err := time.Parse(DATE_TIME, datastr)
	if err == nil {
		unix_time := the_time.Unix()
		fmt.Println(unix_time)
	}
	return the_time.Unix()
}

func Formate(datastr string,layout string) string {
	timeformatdate, _ := time.Parse(DATE_TIME, datastr)
	convdate := timeformatdate.Format(layout)
	return convdate
}