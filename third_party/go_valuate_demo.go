package third_party

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/gogf/gf/util/gconv"
	"github.com/panjf2000/ants/v2"
	"runtime"
	"strings"
	"time"
)

type ExecAlarmTaskCondition struct {
	Indicator      string  `json:"indicator"`       // 数据指标
	OperatorSymbol string  `json:"operator_symbol"` // 操作符
	Threshold      float64 `json:"threshold"`       // 数据指标时有值
	TimeThreshold  []int   `json:"time_threshold"`  // 时间指标时有值
}

func GenerateExpression(conditions []ExecAlarmTaskCondition) string {
	exp := make([]string,0)
	for _,cc := range conditions{
		exp =append(exp, fmt.Sprintf("%s %s %f",cc.Indicator,cc.OperatorSymbol,cc.Threshold))
	}
	return strings.Join(exp," && ")
}

func GenerateDateExpression(conditions []ExecAlarmTaskCondition) string {
	exp := make([]string,0)
	for _,cc := range conditions{
		num2str := make([]string,0)
		for _, d := range cc.TimeThreshold{
			num2str = append(num2str,fmt.Sprintf("%d",d))
		}
		exp =append(exp, fmt.Sprintf("%s %s %s",cc.Indicator,cc.OperatorSymbol,("("+strings.Join(num2str,",")+")")))
	}
	return strings.Join(exp," && ")
}

func ValuateDemo()  {
	rules := []ExecAlarmTaskCondition{
		//{
		//	Indicator:      "age",
		//	OperatorSymbol: "==",
		//	Threshold:      20,
		//	TimeThreshold:  []int{2,3,4},
		//},
		//{
		//	Indicator:      "num",
		//	OperatorSymbol: "<=",
		//	Threshold:      500,
		//	TimeThreshold:  []int{8,9},
		//},

		//-----
		{
			Indicator:      "age",
			OperatorSymbol: "in",
			Threshold:      20,
			TimeThreshold:  []int{2,3,4},
		},
		{
			Indicator:      "num",
			OperatorSymbol: "in",
			Threshold:      500,
			TimeThreshold:  []int{8,9},
		},
	}

	exp := GenerateDateExpression(rules)
	fmt.Println("rule is: ",exp)
	opt := ants.Options{
		ExpiryDuration:  time.Duration(300)*time.Second,
		Nonblocking:      false,
	}
	pool,err := ants.NewPool(runtime.NumCPU(),ants.WithOptions(opt))
	if err != nil {
		return
	}
	err = pool.Submit(func() {
		expression, err := govaluate.NewEvaluableExpression(exp)
		if err != nil {
			fmt.Errorf("run valuate error:%v",err)
			return
		}

		parameters := make(map[string]interface{}, 8)
		parameters["age"] = 20
		parameters["num"] = 2

		result, err := expression.Evaluate(parameters)
		fmt.Println("result is: ",result)
	})

	fmt.Println(gconv.Int64s("30"),err)
}

