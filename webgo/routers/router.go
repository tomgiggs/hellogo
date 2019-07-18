package routers

import (
	"hellogo/webgo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/usercenter", &controllers.UserController{})
    beego.Router("/teacher",&controllers.TeacherController{})
    beego.Router("/learner",&controllers.LearnerController{})
    beego.Router("/coupon",&controllers.CouponController{})
	beego.Router("/chat",&controllers.ChatController{})
	beego.Router("/login",&controllers.LoginController{})
	beego.Router("/logout",&controllers.LogoutController{})
	beego.Router("/location",&controllers.LocationController{})
	beego.Router("/order",&controllers.OrderController{})
	beego.Router("/review",&controllers.ReviewController{})
	beego.Router("/rank",&controllers.RankController{})
	beego.Router("/promotion",&controllers.PromotionController{})
	beego.Router("/register",&controllers.RegisterController{})
	beego.Router("/search",&controllers.SearchController{})
	beego.Router("/skill",&controllers.SkillController{})





}
