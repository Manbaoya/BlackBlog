package controller

type Respond struct {
	Code    int
	Message string
}

var (
	OK     = Respond{200, "成功"}
	Failed = Respond{200, "失败"}
	//客户端

	FailedBind   = Respond{201, "数据绑定失败"}
	EmptyAuth    = Respond{202, "请求头中auth为空"}
	ErrorAuth    = Respond{203, "请求头中auth错误"}
	InvalidToken = Respond{204, "无效的Token"}
	InvalidLogin = Respond{215, "非法登录"}

	InvalidUserName = Respond{205, "用户名不存在"}
	ErrorPassword   = Respond{206, "密码错误"}
	ReUserName      = Respond{207, "该用户名已存在"}
	BadUserName     = Respond{208, "用户名不能包含空格"}
	BadPassword     = Respond{209, "密码只能由数字、大小写字母、‘@’、‘#’、‘$’、‘&’、‘.’组成,并6-12位"}
	DifferentPwd    = Respond{210, "两次密码不一致"}
	RePhone         = Respond{211, "该手机号已被使用"}
	BadPhone        = Respond{212, "手机格式不正确"}
	BadFile         = Respond{213, "文件只能是jpg、jpeg、png格式"}
	FailedUpload    = Respond{214, "上传文件失败"}

	ReTitle = Respond{215, "该标题已存在"}
	BadTime = Respond{216, "截止时间必须大于开始时间！"}

	//服务器
	FailedCreateToken = Respond{501, "生成token失败"}
	FailedGetId       = Respond{502, "获取id失败"}
	FailedGetUsername = Respond{503, "获取用户名失败"}
	FailedEncrypt     = Respond{504, "密码加密失败"}
	FailedDelete      = Respond{505, "删除失败"}
	FailedQuery       = Respond{506, "查询失败"}
	FailedCreate      = Respond{507, "创建失败"}
	FailedAlter       = Respond{508, "修改失败"}
	FailedSaved       = Respond{509, "保存文件失败"}
)
