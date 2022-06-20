package cmd

import (
	"flag/cmd_utils/internal/timer"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
	"time"
)

func init() {
	//time 添加 now 和 calc 子命令
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)

	//calc 添加命令行参数
	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculateTime", "c", "", "需要计算的时间，有效时间为时间搓或者已格式化的额时间")
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", `持续时间，有效单位为"ns", "us" (or "µs"), "ms", "s", "m", "h".`)

}

/**
time 子命令
 */
var calculateTime string
var duration string

/*
time子命令
用户获取当前格式化时间和时间搓
 */
var timeCmd = &cobra.Command{
	Use:"time", //子命令
	Short:"时间格式处理",
	Long: "时间格式处理",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

/**
time now 子命令
运行方式 go run main.go time now
 */
var nowTimeCmd = &cobra.Command{
	Use: "now",
	Short:"获取当前时间",
	Long:"获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		//nowTime.Format("2006-01-02 15:04:05"): 按给定时间格式输出
		//nowTime.Unix(): 发挥Unix时间搓
		log.Fatalf("输出结果： %s,  %d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
	},
}

/**
time calc子命令
验证方法
go run main.go time calc
 */
var calculateTimeCmd = &cobra.Command{
	Use:"calc",
	Short:"计算所需时间",
	Long: "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2006-01-02 15:04:05"
		if calculateTime == "" {
			currentTimer = timer.GetNowTime()
		}else  {
			var err error
			space := strings.Count(calculateTime, " ")
			if space == 0 {
				 layout = "2006-01-02"
			}
			if space == 1 {
				 layout = "2006-01-02 15:04"
			}
			currentTimer, err = time.Parse(layout, calculateTime)
			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t), 0)
			}
		}
		t, err := timer.GetCalculateTime(currentTimer, duration)
		if err != nil {
			log.Fatalf("timer.GetCalculateTime err: %v", err)
		}
		log.Printf("输出结果：%s, %d", t.Format(layout), t.Unix())
	},
}

