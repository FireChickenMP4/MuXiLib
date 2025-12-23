package models

// @Summary 用户注册请求
// @Description 注册信息
type RegisterRequst struct {
	Username string `json:"username" binding:"require,min=3,max=20"`
	Password string `json:"password" binding:"require,min=6,max=25"`
}

// @Summary 用户登录请求
// @Description 登录信息
type LoginRequst struct {
	Username string `json:"username" binding:"require"`
	Password string `json:"password" binding:"require"`
}
