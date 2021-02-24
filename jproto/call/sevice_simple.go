package call

import (
	"context"
	pb "github.com/ijidan/jproto/jproto/build"
)

//简单模式
type SimpleService struct {
	BasicService
}

//获取服务名称
func (s *SimpleService) GetName() string {
	return "simple_stream"
}

//函数实现
func (s *SimpleService) DoRequest(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	rsp := new(pb.SimpleResponse)
	rsp.RspIdx = "simple_idx"
	rsp.RspData = "simple_data"
	return rsp, nil
}

//获取实例
func NewSimpleService() *SimpleService {
	instance := &SimpleService{}
	return instance
}
