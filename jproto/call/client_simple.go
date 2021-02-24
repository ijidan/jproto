package call

import (
	pb "github.com/ijidan/jproto/jproto/build"
)

//简单服务
type SimpleCall struct {
	BasicCall
	client pb.SimpleClient
}

//构建客户端
func (c *SimpleCall) buildClient() {
	c.buildConnection()
	c.buildCtx()
	client := pb.NewSimpleClient(c.conn)
	c.client = client
}

//函数实现
func (c *SimpleCall) DoRequest(reqIdx string, reqData string) (string, string, error) {
	c.buildClient()
	defer c.GetConn().Close()
	//请求
	request := &pb.SimpleRequest{
		ReqIdx:  reqIdx,
		ReqData: reqData,
	}
	r, err := c.client.DoRequest(c.ctx, request)
	if err != nil {
		return "", "", err
	}
	rspIdx := r.GetRspIdx()
	rspData := r.GetRspData()
	return rspIdx, rspData, nil
}
