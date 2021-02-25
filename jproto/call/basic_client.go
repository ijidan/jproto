package call

import (
	"context"
	 "github.com/ijidan/jnet/jnet"
	"github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	status2 "google.golang.org/grpc/status"
	"io"
	"time"
)

//基本调用类
type BasicCall struct {
	conn   *grpc.ClientConn
	ctx    context.Context
	tracer opentracing.Tracer
	closer io.Closer
	span   opentracing.Span
}

//构建ctx
func (c *BasicCall) buildCtx() {
	//上下文
	ctx, _ := context.WithTimeout(context.Background(), time.Hour)
	//defer cancel()
	c.ctx = ctx
}


//构建链接
func (c *BasicCall) buildConnection() () {
	address := getAddress()
	conn := jnet.BuildConnection(address)
	//defer conn.Close()
	c.conn = conn
}

//关闭连接
func (c *BasicCall) CloseConn() {
	_ = c.conn.Close()
}

//结束时操作
func (c *BasicCall) endHandler() {
	//c.closeConn()
}

//获取连接
func (c *BasicCall) GetConn() *grpc.ClientConn {
	return c.conn
}

//获取上下文
func (c *BasicCall) GetCtx() context.Context {
	return c.ctx
}

//关闭
func (c *BasicCall) Close() {
	_ = c.conn.Close()
}

//检测超时
func (c *BasicCall) CheckTimeout(err error) {
	c.doCheckErrorCode(err, codes.DeadlineExceeded, "gRpc timeout!")
}

//检测取消
func (c *BasicCall) CheckCancel(err error) {
	c.doCheckErrorCode(err, codes.Canceled, "gRpc canceled")
}

//检测错误
func (c *BasicCall) CheckError(err error) {
	if status, ok := status2.FromError(err); ok {
		message := status.Message()
		log.Error(message)
	}
}

//检测错误
func (c *BasicCall) doCheckErrorCode(err error, code codes.Code, message string) {
	if status, ok := status2.FromError(err); ok {
		statusCode := status.Code()
		if statusCode == code {
			log.Error(message)
		}
	}
}
