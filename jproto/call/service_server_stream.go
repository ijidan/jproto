package call

import (
	pb "github.com/ijidan/jproto/jproto/build"
)

//服务端流
type ServerStreamService struct {
	BasicService
}

//获取服务名称
func (s *ServerStreamService) GetName() string {
	return "server_stream"
}

//函数实现
func (s *ServerStreamService) DoRequest(req *pb.ServerStreamRequest, srv pb.ServerStream_DoRequestServer) error {
	for i := 0; i < 100; i++ {
		rsp := new(pb.ServerStreamResponse)
		rsp.RspIdx = "server_stream_" + string(i)
		rsp.RspData = "server_data_" + string(i)
		err := srv.Send(rsp)
		if err != nil {
			return err
		}
	}
	return nil
}

//获取实例
func NewServerStreamService() *ServerStreamService {
	instance := &ServerStreamService{}
	return instance
}
