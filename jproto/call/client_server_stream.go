package call

import (
	pb "github.com/ijidan/jproto/jproto/build"
	log "github.com/sirupsen/logrus"
	"io"
)

//服务流客户端
type ServerStreamCall struct {
	BasicCall
	client pb.ServerStreamClient
}

//构建客户端
func (c *ServerStreamCall) buildClient() {
	c.buildConnection()
	c.buildCtx()
	client := pb.NewServerStreamClient(c.conn)
	c.client = client
}

//函数实现
func (c *ServerStreamCall) DoRequest(reqIdx string, reqData string) (string, error) {
	c.buildClient()
	defer c.GetConn().Close()
	//请求
	request := &pb.ServerStreamRequest{
		ReqIdx:  reqIdx,
		ReqData: reqData,
	}
	stream, err := c.client.DoRequest(c.ctx, request)
	if err != nil {
		return "", err
	}
	//循环读取数据
	for {
		res, err := stream.Recv()
		//读取结束
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Error("server stream get stream error：" + err.Error())
			return "", err
		}
		if res == nil {
			log.Error("server stream get stream error：res is nil")
			continue
		}
		rspIdx := res.GetRspIdx()
		rspData := res.GetRspData()
		//输出返回值
		log.Info("接收到数据:" + rspIdx + "," + rspData)
	}
	return "", nil
}
