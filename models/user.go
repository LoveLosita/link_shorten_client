package models

type Status struct {
	Code    string `thrift:"code,1,required" frugal:"1,required,i32" json:"code"`
	Message string `thrift:"message,2,required" frugal:"2,required,string" json:"message"`
}
type UserRegisterRequest struct {
	Username    string `thrift:"username,1,required" frugal:"1,required,string" json:"username"`
	Password    string `thrift:"password,2,required" frugal:"2,required,string" json:"password"`
	Gender      string `thrift:"gender,3,required" frugal:"3,required,string" json:"gender"`
	PhoneNumber string `thrift:"phone_number,4,required" frugal:"4,required,string" json:"phone_number"`
	Email       string `thrift:"email,5,required" frugal:"5,required,string" json:"email"`
}
type UserRegisterResponse struct {
	Status *Status `thrift:"status,1,required" frugal:"1,required,Code" json:"status"`
}
type UserLoginRequest struct {
	Username string `thrift:"username,1,required" frugal:"1,required,string" json:"username"`
	Password string `thrift:"password,2,required" frugal:"2,required,string" json:"password"`
}
type UserLoginResponseData struct {
	AccessToken  string `thrift:"access_token,1,required" frugal:"1,required,string" json:"access_token"`
	RefreshToken string `thrift:"refresh_token,2,required" frugal:"2,required,string" json:"refresh_token"`
}
type UserLoginResponse struct {
	Status *Status                `thrift:"status,1,required" frugal:"1,required,Code" json:"status"`
	Data   *UserLoginResponseData `thrift:"data,2,required" frugal:"2,required,UserLoginResponseData" json:"data"`
}
type TokenRefreshRequest struct {
	RefreshToken string `thrift:"refresh_token,1,required" frugal:"1,required,string" json:"refresh_token"`
}
type TokenRefreshResponseData struct {
	AccessToken  string `thrift:"access_token,1,required" frugal:"1,required,string" json:"access_token"`
	RefreshToken string `thrift:"refresh_token,2,required" frugal:"2,required,string" json:"refresh_token"`
}
type TokenRefreshResponse struct {
	Status *Status                   `thrift:"status,1,required" frugal:"1,required,Code" json:"status"`
	Data   *TokenRefreshResponseData `thrift:"data,2,required" frugal:"2,required,TokenRefreshResponseData" json:"data"`
}
