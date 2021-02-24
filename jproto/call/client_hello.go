package call

import (
	pb "github.com/ijidan/jproto/jproto/build"
)

//服务
type HelloCall struct {
	BasicCall
	client pb.HelloClient
}

//构建客户端
func (c *HelloCall) buildClient() {
	c.buildConnection()
	c.buildCtx()
	client := pb.NewHelloClient(c.conn)
	c.client = client
}

//函数实现
func (c *HelloCall) SayHello(name string) (string, error) {
	c.buildClient()
	defer c.endHandler()
	//请求
	request := &pb.HelloRequest{Name: name}
	r, err := c.client.SayHello(c.ctx, request)
	if err != nil {
		return "", err
	}
	data := r.GetMessage()
	return data, nil
}

//函数实现
func (c *HelloCall) SayJidan(nameCn string, nameEn string) (map[string]string, error) {
	c.buildClient()
	defer c.endHandler()
	//请求
	request := &pb.JidanRequest{NameCn: nameCn, NameEn: nameEn}
	r, err := c.client.SayJidan(c.ctx, request)
	if err != nil {
		return nil, err
	}
	data := map[string]string{"name_cn": r.GetMessageCn(), "name_en": r.GetMessageEn()}
	return data, nil
}
