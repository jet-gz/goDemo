package mystruct

type UserInfo struct {
	UserName string `json:"userName" form:"userName" binding:"required"` // 验证非空  好像没起作用
	Password string `json:"pwd" form:"pwd"`
}
