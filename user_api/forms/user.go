package forms

type PassWordLoginForm struct {
	Mobile    string `form:"mobile" binding:"required,mobile"`
	PassWord  string `form:"password" binding:"required,min=3,max=10"`
	Captcha   string `form:"captcha" binding:"required,min=5,max=5"`
	CaptchaId string `form:"captcha_id" binding:"required"`
}

type RegisterForm struct {
	Mobile   string `form:"mobile" binding:"required,mobile"`
	PassWord string `form:"password" binding:"required,min=3,max=10"`
	Code     string `form:"code" binding:"required,min=6,max=6"`
}
