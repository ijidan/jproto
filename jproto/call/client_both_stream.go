package call

import (
	"errors"
	pb "github.com/ijidan/jproto/jproto/build"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"io"
	"strconv"
)

//简单服务
type BothStreamCall struct {
	BasicCall
	client  pb.BothStreamClient
	readCh  chan *pb.BothStreamResponse //读取消息
	writeCh chan *pb.BothStreamRequest  //发送消息
}

//构建客户端
func (c *BothStreamCall) buildClient() {
	c.buildConnection()
	c.buildCtx()
	client := pb.NewBothStreamClient(c.conn)
	c.client = client
}

//函数实现：主函数
func (c *BothStreamCall) DoRequest(reqIdx string, reqData string) error {
	c.buildClient()
	defer c.GetConn().Close()
	//请求
	opts := grpc.EmptyCallOption{}
	client, err := c.client.DoRequest(c.ctx, opts)
	if err != nil {
		log.Error("both stream init client error:" + err.Error())
		return err
	}
	//读取信息
	go func() {
		_ = c.readMessage(client)
	}()
	//发送信息
	go func() {
		_ = c.sendMessage(client, reqIdx, reqData)
	}()
	//处理关闭信号
	for {
		select {
		case rsp := <-c.readCh:
			rspIdx := rsp.GetRspIdx()
			rspData := rsp.GetRspData()
			log.Info("~~~~~~~~~~~~~~~~~~~~~get %s %s", rspIdx, rspData)
		case req := <-c.writeCh:
			idx := req.GetReqIdx()
			log.Info("---------------------send %s", idx)
			err = client.Send(req)
			//发送错误
			if err != nil {
				return err
			}
		default:
		}
		if c.readCh == nil && c.writeCh == nil {
			return nil
		}
	}
}

//读取消息
func (c *BothStreamCall) readMessage(client pb.BothStream_DoRequestClient) error {
	for {
		//循环读取
		rsp, err := client.Recv()
		//服务端主动关闭
		if err == io.EOF {
			close(c.readCh)
			return err
		}
		//读取错误
		if err != nil {
			close(c.readCh)
			return err
		}
		//响应错误
		if rsp == nil {
			close(c.readCh)
			return errors.New("rsp nil")
		}
		c.readCh <- rsp
	}
}

//发送信息
func (c *BothStreamCall) sendMessage(client pb.BothStream_DoRequestClient, reqIdx string, reqData string) error {
	//发送信息
	n := 1
	closeSignal := 100
	var err error
	for {
		//关闭流并获取返回消息
		if n > closeSignal {
			err = client.CloseSend()
			close(c.writeCh)
			if err != nil {
				return err
			}
			return nil
		}
		//发送内容
		nStr := strconv.Itoa(n)
		request := &pb.BothStreamRequest{
			ReqIdx:  reqIdx + " " + nStr,
			ReqData: reqData + " " + nStr,
		}
		c.writeCh <- request
		n++
	}
}

//获取实例
func NewBothStreamCall() *BothStreamCall {
	instance := &BothStreamCall{
		readCh:  make(chan *pb.BothStreamResponse),
		writeCh: make(chan *pb.BothStreamRequest),
	}
	return instance
}
