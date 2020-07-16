package model

//私有
type person struct {
	Name string
	age  int     //私有
	sal  float64 //私有
}

//工厂模式	即Java有参构造方法
func NewPerson(name string) *person {
	return &person{Name: name}
}

//为了访问私有字段		即Java get()set()
func (p *person) SetAge(age int) {
	p.age = age
}
func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSal(sal float64) {
	p.sal = sal
}
func (p *person) GetSal() float64 {
	return p.sal
}
