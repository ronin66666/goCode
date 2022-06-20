package timer

import "time"

func GetNowTime() time.Time {
	return time.Now()
}

func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	//从字符串中解析出duration，支持格式见方法注释
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	return currentTimer.Add(duration), nil
}



