package call

import (
	pb "github.com/ijidan/jproto/jproto/build"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"io"
)

//简单服务
type ClientStreamCall struct {
	BasicCall
	client pb.ClientStreamClient
}

//构建客户端
func (c *ClientStreamCall) buildClient() {
	c.buildConnection()
	c.buildCtx()
	client := pb.NewClientStreamClient(c.conn)
	c.client = client
}

//函数实现
func (c *ClientStreamCall) DoRequest(reqIdx string, reqData string) error {
	c.buildClient()
	defer c.GetConn().Close()
	//请求
	opts := grpc.EmptyCallOption{}
	client, err := c.client.DoRequest(c.ctx, opts)
	if err != nil {
		log.Error("client stream init client error:" + err.Error())
		return err
	}
	//循环发送内容
	for n := 0; n <= 100; n++ {
		request := &pb.ClientStreamRequest{
			ReqIdx:  reqIdx,
			ReqData: reqData,
		}
		err := client.Send(request)
		//检测服务端主动关闭
		if err == io.EOF {
			log.Info("client stream:server closed")
			break
		}
		if err != nil {
			log.Error("client stream send error：" + err.Error())
		}
	}
	//关闭流并获取返回消息
	err = client.CloseSend()
	if err != nil {
		log.Error("client stream client close error：" + err.Error())
		return err
	}
	return nil
}
