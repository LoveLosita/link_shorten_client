package response

type Status struct {
	Code    string `thrift:"code,1,required" frugal:"1,required,i32" json:"code"`
	Message string `thrift:"message,2,required" frugal:"2,required,string" json:"message"`
}

func InternalError(err error) Status { //服务器错误
	return Status{
		Code:    "500",
		Message: err.Error(),
	}
}

var (
	Ok = Status{ //正常
		Code:    "10000",
		Message: "success",
	}

	WrongName = Status{ //用户名错误
		Code:    "40001",
		Message: "wrong username",
	}

	WrongPwd = Status{ //密码错误
		Code:    "40002",
		Message: "wrong password",
	}

	InvalidName = Status{ //用户名无效
		Code:    "40003",
		Message: "the username already exists",
	}

	MissingParam = Status{ //缺少参数
		Code:    "40004",
		Message: "missing param",
	}

	WrongParamType = Status{ //参数错误
		Code:    "40005",
		Message: "wrong param type",
	}

	ParamTooLong = Status{ //参数过长
		Code:    "40006",
		Message: "param too long",
	}

	WrongUsernameOrPwd = Status{ //用户名或密码错误
		Code:    "40007",
		Message: "wrong username or password",
	}

	WrongGender = Status{ //性别错误
		Code:    "40008",
		Message: "wrong gender",
	}
)
