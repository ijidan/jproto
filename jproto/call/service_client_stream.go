package call

import (
	pb "github.com/ijidan/jproto/jproto/build"
	log "github.com/sirupsen/logrus"
	"io"
)

//客户端流
type ClientStreamService struct {
	BasicService
}

//获取服务名称
func (s *ClientStreamService) GetName() string {
	return "client_stream"
}

//函数实现
func (s *ClientStreamService) DoRequest(srv pb.ClientStream_DoRequestServer) error {
	for {
		res, err := srv.Recv()
		//检测客户端主动关闭
		if err == io.EOF {
			log.Error("接收到EOF...")
			//发送结果，并关闭
			return srv.SendAndClose(&pb.ClientStreamResponse{
				RspIdx:  "client_stream_idx_eof",
				RspData: "client_stream_data_eof",
			})
		}
		//读取错误
		if err != nil {
			return err
		}
		reqIdx := res.GetReqIdx()
		reqData := res.GetReqData()
		//输出返回值
		log.Info("接收到数据:" + reqIdx + "," + reqData)
	}
}

//获取实例
func NewClientStreamService() *ClientStreamService {
	instance := &ClientStreamService{}
	return instance
}
