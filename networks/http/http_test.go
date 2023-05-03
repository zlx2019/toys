/**
  @author: Zero
  @date: 2023/5/3 09:24:56
  @desc: http函数库单元测试

**/

package http

type Course struct {
	TeacherId uint   `json:"teacher_id"`
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	Time      string `json:"time"`
}
