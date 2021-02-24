package call

import (
	pb "github.com/ijidan/jproto/jproto/build"
	log "github.com/sirupsen/logrus"
	"io"
	"strconv"
)

//双向流
type BothStreamService struct {
	BasicService
	//队列使用数组吧。不然要阻塞
	readCh  chan *pb.BothStreamRequest  //读取消息
	writeCh chan *pb.BothStreamResponse //关闭消息
}

//获取服务名称
func (s *BothStreamService) GetName() string {
	return "both_stream"
}

//函数实现：主函数，不能退出。不然那就断了。
func (s *BothStreamService) DoRequest(srv pb.BothStream_DoRequestServer) error {
	srv.Context()
	//读取信息
	go func() {
		_ = s.readMessage(srv)
	}()
	//发送信息
	go func() {
		_ = s.sendMessage()
	}()
	//处理关闭信号
	for {
		select {
		case req := <-s.readCh:
			reqIdx := req.GetReqIdx()
			reqData := req.GetReqData()
			log.Info("~~~~~~~~~~~~~~~~~~~~~get :%s %s", reqIdx, reqData)
		case rsp := <-s.writeCh:
			idx := rsp.GetRspIdx()
			log.Info("---------------------send %s", idx)
			err := srv.Send(rsp)
			//发送错误
			if err != nil {
				return err
			}
		default:
		}
	}
}

//读取信息
func (s *BothStreamService) readMessage(srv pb.BothStream_DoRequestServer) error {
	for {
		//读取
		req, err := srv.Recv()
		//客户端主动关闭
		if err == io.EOF {
			return nil
		}
		//读取错误：读取错误没处理样
		if err != nil {
			return err
		}
		s.readCh <- req
	}
}

//写信息
func (s *BothStreamService) sendMessage() error {
	//发送信息
	n := 1
	for {
		//发送数据
		rsp := new(pb.BothStreamResponse)
		rsp.RspIdx = "both_stream_idx " + strconv.Itoa(n)
		rsp.RspData = "both_stream_data " + strconv.Itoa(n)
		s.writeCh <- rsp
		n++
	}
}

//获取实例
func NewBothStreamService() *BothStreamService {
	service := &BothStreamService{
		readCh:  make(chan *pb.BothStreamRequest),
		writeCh: make(chan *pb.BothStreamResponse),
	}
	return service
}
