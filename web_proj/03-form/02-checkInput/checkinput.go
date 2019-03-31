// 表单元素是否为空，如 username
if len(r.Form["username"][0]) == 0 {
	// 为空的处理
}

// 数字,确保表单输入框获取的只能是数字
// 如果判断是正整数，那么先转化为int类型
getint, err := strconv.Atoi(r.Form.Get("age"))
if err != nil {
	// 数字转化出错了，那么可能不是数字
}

// 另一种方式，正则匹配
if m,_ := regexp.MathString("^[0-9]+$", r.Form.Get(age)); !m {
	return false
}

// 中文
// 为保证获取的是正确的中文，有两种方式来验证
// unicode包提供的 func Is(rangTab *RangeTable,r rune) bool 来验证
// 正则方式 
if m,_ := regexp.MatchString("^\\p{Han}+$",r.Form.Get("realname")); !m {
	return false
}


// 英文
if m,_ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("engname")); !m {
	return false
}

// 邮箱地址
if m,_ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4}$`, r.Form.Get("email")); !m {
	return false
}

// 手机号码
if m,_ regexp.MatchString(`^1[3|4|5|8][0-9]\d{4,8}$`,r.Form.Get("mobile")); !m {
	return false
}