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
	Ok = Status{ //正常,和服务端统一
		Code:    "10000",
		Message: "ok",
	}

	WrongParamType = Status{ //参数错误
		Code:    "41001", //此处错误码若是41xxx，则是客户端的错误码；若是40xxx，则是服务端的错误码
		Message: "wrong param type",
	}

	ParamTooLong = Status{ //参数过长
		Code:    "41002",
		Message: "param too long",
	}

	RpcServerConnectTimeOut = Status{ //rpc服务端连接超时
		Code:    "41003",
		Message: "rpc server connect time out",
	}
)
