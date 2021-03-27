package main

// NameInterface 接口
type NameInterface interface {
	GetName() string
}

// StringInterface 接口
type StringInterface interface {
	GetString() string
}

// RefundService 接口
// 这个接口里面的方法不要自调用，不然会引起无限循环调用
type RefundService interface {
	GetRefundInfo() string
}

// BaseAbstruct 抽象接口
type BaseAbstruct struct {
	NameInterface
	StringInterface

	Name string
}

// GetRefundInfo 实现了RefundService接口
func (ba *BaseAbstruct) GetRefundInfo() string {
	return ba.GetName() + ba.GetString()
}

// StringAbstruct 抽象, 子类不要调用父类已经实现乐了的方法，除非子类重写了父类的方法
type StringAbstruct struct {
	BaseAbstruct
}

// GetString 实现StringInterface
func (str *StringAbstruct) GetString() string {
	return str.GetName() + "_what"
}

// HotelRefundService 真正的结构体
type HotelRefundService struct {
	StringAbstruct
}

// GetName 实现NameInterface
func (hotel *HotelRefundService) GetName() string {
	return "what"
}

// PrintName 打印信息出来
func PrintName(service RefundService) {
	println(service.GetRefundInfo())
}

func main() {
	service := &HotelRefundService{}
	service.Name = "test"
	service.NameInterface = service
	service.StringInterface = service
	PrintName(service)

	// abService := &BaseAbstruct{}
	// abService.Name = "abstruct"
	// PrintName(abService)

}
