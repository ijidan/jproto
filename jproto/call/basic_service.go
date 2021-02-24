package call

//基本服务类
type BasicService struct {
	Name string
}

//获取服务名称
func (s *BasicService) GetName() string {
	return s.Name
}

//func (s *BasicService)  {
//
//}
