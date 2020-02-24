package validator

//验证获取用户信息
type GetUserUri struct {
	Uid uint `uri:"uid" binding:"required,gt=0"`
}

//注册验证
type RegisterForm struct {
	Username string `form:"username" binding:"required,max=100,min=4"`
	Password string `form:"password" binding:"required,max=200,min=6"`
	Mobile   string `form:"mobile" binding:"omitempty,c_mobile"`
	Email    string `form:"email" binding:"required_with=Mobile,email"`
}

//登录验证
type LoginForm struct {
	Username string `form:"username" binding:"required,max=100,min=4"`
	Password string `form:"password" binding:"required,max=200,min=6"`
}
