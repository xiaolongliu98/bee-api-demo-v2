package response

// Response结构体
type Response struct {
	Id int `orm:"pk;"`

	APIId  int    `orm:"column(api_id)" json:"apiId"` // api_id int UNSIGNED
	Msg    string // 消息 varchar(512)
	Code   int    // 代码 int
	Result string // 返回结果 varchar(8192)
}

// 指定Response结构体默认绑定的表名

func (um *Response) TableName() string {
	return "response"
}
