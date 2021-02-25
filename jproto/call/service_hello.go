package call

import (
	"context"
	pb "github.com/ijidan/jproto/jproto/build"
	"github.com/opentracing/opentracing-go"
)

//Hello服务
type HelloService struct {
	BasicService
}

//获取服务名称


func (s *HelloService) GetName() string {
	return "hello"
}

//函数实现
func (s *HelloService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	rsp := new(pb.HelloResponse)
	rsp.Message = "Hello " + req.Name + "."

	span, _ := opentracing.StartSpanFromContext(ctx, s.GetName())
	defer func() {
		span.SetTag("request", req)
		span.SetTag("reply", rsp)
		span.Finish()
	}()

	return rsp, nil

}

//函数实现
func (s *HelloService) SayJidan(ctx context.Context, req *pb.JidanRequest) (*pb.JidanResponse, error) {
	rsp := new(pb.JidanResponse)
	rsp.MessageCn = "鸡蛋：" + req.NameCn
	rsp.MessageEn = "jidan:" + req.NameEn

	span, _ := opentracing.StartSpanFromContext(ctx, s.GetName())
	defer func() {
		// 4. 接口调用完，在tag中设置request和reply
		span.SetTag("request", req)
		span.SetTag("reply", rsp)
		span.Finish()
	}()

	return rsp, nil
}

//获取实例
func NewHelloService() *HelloService {
	instance := &HelloService{}
	return instance
}
