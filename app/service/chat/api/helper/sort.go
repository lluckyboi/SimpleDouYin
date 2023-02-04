package helper

import (
	"SimpleDouYin/app/service/chat/dao/model"
)

func SortMessageByDate(arr1, arr2 []model.Message) []model.Message {
	var (
		l1   = len(arr1)
		l2   = len(arr2)
		idx1 = 0
		idx2 = 0
		res  []model.Message
	)
	for {
		//有一方全部插入，结束
		if idx1 >= l1 || idx2 >= l2 {
			break
		}
		//比较时间,较大的先插入
		if arr1[idx1].CreateTime.After(arr2[idx2].CreateTime) {
			res = append(res, arr1[idx1])
			res = append(res, arr2[idx2])
			idx1++
			idx2++
		} else {
			res = append(res, arr2[idx2])
			res = append(res, arr1[idx1])
			idx1++
			idx2++
		}
	}
	//插入剩余数据
	if idx1 >= l1 {
		res = append(res, arr2[idx2:l2]...)
	} else {
		res = append(res, arr1[idx1:l1]...)
	}

	return res
}
