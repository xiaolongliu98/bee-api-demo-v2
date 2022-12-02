package response

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

// GetResponseById
// @Description 通过序号id获取Response
// @Author 		liyi
// @Date		2022/12/1 16:50(create);
// @Param		id		int		序号id
// @Return		*response.Response, error
func GetResponseById(id int) (*Response, error) {
	o := orm.NewOrm()
	// -------------------------------------
	r := &Response{Id: id}

	err := o.Read(r)
	if err != nil {
		//fmt.Printf("%v\n", err.Error())
		return nil, err
	}

	return r, err
}

// AddResponse
// @Description 插入一条用户核心数据，成功后在um中会保存uid
// @Author 		liyi
// @Date		2022/12/1 17:00(create);
// @Param		r		*Response		用户核心数据
// @Return		error
func AddResponse(r *Response) error {
	if r == nil {
		return fmt.Errorf("r is nil")
	}

	o := orm.NewOrm()
	id, err := o.Insert(r)
	if err != nil {
		return err
	}
	r.Id = int(id)
	return nil
}
