package helper

import "time"

/**/
type TimeEnum struct {
	WeeksMap  map[int]string
	Monday    int
	Tuesday   int
	Wednesday int
	Thursday  int
	Friday    int
	Saturday  int
	Sunday    int

	WorkHoursStartCommon string
	WorkHoursEndCommon   string

	//上午上班时间
	WorkHoursStartUnusual string
	//上午下班时间
	WorkHourAmEndCommon string

	WorkHoursEndUnusual string
	//下午一点
	WorkHourPmsStartCommon string

	//日期转换格式
	DateTimeLayout   string
	DateLayout       string
	DateLayoutAll    string
	DateLayoutNumber string
	TimeLayout       string
	DateLayoutAllEnd string
	StartTime        string
	EndTime          string
	EmptyStartTime   string
	EmptyEndTime     string
	YEARDATE         string
	YEAR             string
	MONTH            string
	//'0000-00-00 00:00:00' 查出的结果是这个 "0001-01-01 00:00:00"
	OUT_ZERO_TIME        string
	NOWTIME_STRING       string
	DateTimeLayoutNumber string
}

var TimeEnumObject TimeEnum

const (
	DateTimeStringLen = 19
)

/**
 *  init
 *  @Description:
 */
func init() {

	TimeEnumObject.Monday = 1
	TimeEnumObject.Tuesday = 2
	TimeEnumObject.Wednesday = 3
	TimeEnumObject.Thursday = 4
	TimeEnumObject.Friday = 5
	TimeEnumObject.Saturday = 6
	TimeEnumObject.Sunday = 0

	TimeEnumObject.WeeksMap = map[int]string{
		TimeEnumObject.Monday:    "周一",
		TimeEnumObject.Tuesday:   "周二",
		TimeEnumObject.Wednesday: "周三",
		TimeEnumObject.Thursday:  "周四",
		TimeEnumObject.Friday:    "周五",
		TimeEnumObject.Saturday:  "周六",
		TimeEnumObject.Sunday:    "周日",
	}

	TimeEnumObject.OUT_ZERO_TIME = "0001-01-01 00:00:00"

	TimeEnumObject.WorkHoursStartCommon = "09:30:00"
	TimeEnumObject.WorkHoursEndCommon = "19:00:00"

	TimeEnumObject.WorkHoursStartUnusual = "10:00:00"
	TimeEnumObject.WorkHoursEndUnusual = "17:30:00"
	TimeEnumObject.YEARDATE = "2006-01"
	TimeEnumObject.YEAR = "2006"
	TimeEnumObject.MONTH = "01"
	//下午1点
	TimeEnumObject.WorkHourPmsStartCommon = "13:00:00"
	TimeEnumObject.WorkHourAmEndCommon = "12:00:00"

	//
	TimeEnumObject.StartTime = "00:00:00"
	TimeEnumObject.EmptyStartTime = " 00:00:00"
	TimeEnumObject.EndTime = "23:59:59"
	TimeEnumObject.EmptyEndTime = " 23:59:59"

	//日期转换格式
	TimeEnumObject.DateTimeLayout = "2006-01-02 15:04:05"
	TimeEnumObject.DateLayout = "2006-01-02"
	TimeEnumObject.DateLayoutAll = "2006-01-02 00:00:00"
	TimeEnumObject.DateLayoutNumber = "20060102"
	TimeEnumObject.TimeLayout = "15:04:05"
	TimeEnumObject.NOWTIME_STRING = time.Now().Format(TimeEnumObject.DateTimeLayout)
	TimeEnumObject.DateTimeLayoutNumber = "20060102150405"
}

func (receiver TimeEnum) name() {

}

type TimeEnumFunc struct {
}

var TimeEnumFuncObject TimeEnumFunc

// 时间转换
func (e *TimeEnumFunc) StringTime(date time.Time) string {
	stringData := date.Format(TimeEnumObject.DateTimeLayout)
	if stringData == "0001-01-01 00:00:00" {
		stringData = ""
	}
	return stringData
}

// 获取年月
func (e *TimeEnumFunc) StringYearAndDate(date time.Time) string {
	return date.Format(TimeEnumObject.YEARDATE)
}

// 获取年月日
func (e *TimeEnumFunc) StringYearAndMouthAndDay(date time.Time) string {

	stringData := date.Format(TimeEnumObject.DateLayout)
	if stringData == "0001-01-01" {
		stringData = ""
	}
	return stringData
}

// 获取年
func (e *TimeEnumFunc) StringYear(date time.Time) string {
	return date.Format(TimeEnumObject.YEAR)
}

// 获取月
func (e *TimeEnumFunc) StringMonth(date time.Time) string {
	return date.Format(TimeEnumObject.MONTH)
}

// 获取当前月的第一天的时间
func (e *TimeEnumFunc) TimeNowMonth(date time.Time) time.Time {
	yearDate := date.Format(TimeEnumObject.YEARDATE)
	monthData, err := time.ParseInLocation(TimeEnumObject.DateLayoutAll, yearDate+"-01 00:00:00", time.Local)
	if err != nil {
		return date
	}
	return monthData
}

// 获取当前时间的前一天时间
func (e *TimeEnumFunc) YesterdayString(date time.Time) string {
	stringData := date.AddDate(0, 0, -1).Format(TimeEnumObject.DateLayout)
	if stringData == "0001-01-01" {
		stringData = ""
	}
	return stringData
}

func (e *TimeEnumFunc) TimeUnic(date string) int64 {
	t, _ := time.ParseInLocation(TimeEnumObject.DateTimeLayout, date, time.Local)
	return t.Unix()
}
