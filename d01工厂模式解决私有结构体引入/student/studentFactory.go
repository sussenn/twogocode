package student

//首字母小写,无法被其他包访问	可使用"工厂模式"
type student struct {
	Name  string
	score float64 //私有
}

//工厂模式
//返回指针类型的 结构体对象
func NewStudent(n string, s float64) *student {
	return &student{n, s}
}

//使私有字段score 能被其他包获取		方法名首字目必须大写!
func (stu *student) GetScore() float64 {
	//(*stu).score	的简写
	return stu.score
}
