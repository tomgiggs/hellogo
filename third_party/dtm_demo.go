package third_party

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yedf/dtm/common"
	"github.com/yedf/dtm/dtmcli"
)

// 启动命令：go run app/main.go qs

// 事务参与者的服务地址
const qsBusiAPI = "/api/busi_start"
const qsBusiPort = 8082

// DtmServer dtm service address
const DtmServer = "http://localhost:8080/api/dtmsvr"

var qsBusi = fmt.Sprintf("http://localhost:%d%s", qsBusiPort, qsBusiAPI)
var config = common.DtmConfig

func sdbGet() *sql.DB {
	db, err := dtmcli.SdbGet(config.DB)
	dtmcli.FatalIfError(err)
	return db
}

// QsStartSvr 1
func QsStartSvr() {
	app := common.GetGinApp()
	qsAddRoute(app)
	dtmcli.Logf("quick qs examples listening at %d", qsBusiPort)
	go app.Run(fmt.Sprintf(":%d", qsBusiPort))
	time.Sleep(100 * time.Millisecond)
}

// QsFireRequest 1
func QsFireRequest() string {
	req := &gin.H{"amount": 30} // 微服务的载荷
	// DtmServer为DTM服务的地址
	saga := dtmcli.NewSaga(DtmServer, dtmcli.MustGenGid(DtmServer)).
		// 添加一个TransOut的子事务，正向操作为url: qsBusi+"/TransOut"， 逆向操作为url: qsBusi+"/TransOutCompensate"
		Add(qsBusi+"/TransOut", qsBusi+"/TransOutCompensate", req).
		// 添加一个TransIn的子事务，正向操作为url: qsBusi+"/TransOut"， 逆向操作为url: qsBusi+"/TransInCompensate"
		Add(qsBusi+"/TransIn", qsBusi+"/TransInCompensate", req)
	// 提交saga事务，dtm会完成所有的子事务/回滚所有的子事务
	err := saga.Submit()
	dtmcli.FatalIfError(err)
	return saga.Gid
}

func qsAdjustBalance(uid int, amount int) (interface{}, error) {
	_, err := dtmcli.DBExec(sdbGet(), "update dtm_busi.user_account set balance = balance + ? where user_id = ?", amount, uid)
	return dtmcli.ResultSuccess, err
}

func qsAddRoute(app *gin.Engine) {

	qsapi := app.Group(qsBusiAPI)
	{
		qsapi.POST("/TransIn", common.WrapHandler(func(c *gin.Context) (interface{}, error) {
			return qsAdjustBalance(2, 30)
		}))
		qsapi.POST("/TransInCompensate", common.WrapHandler(func(c *gin.Context) (interface{}, error) {
			return qsAdjustBalance(2, -30)
		}))
		qsapi.POST("/TransOut", common.WrapHandler(func(c *gin.Context) (interface{}, error) {
			return qsAdjustBalance(1, -30)
		}))
		qsapi.POST("/TransOutCompensate", common.WrapHandler(func(c *gin.Context) (interface{}, error) {
			return qsAdjustBalance(1, 30)
		}))
	}

}
