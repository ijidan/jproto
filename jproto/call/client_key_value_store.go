package call

import (
	pb "github.com/ijidan/jproto/jproto/build"
)

//服务
type KeyValueStoreCall struct {
	BasicCall
	client pb.KeyValueStoreClient
}

//构建客户端
func (c *KeyValueStoreCall) buildClient() {
	c.buildConnection()
	c.buildCtx()
	client := pb.NewKeyValueStoreClient(c.conn)
	c.client = client
}

//获取值
func (c *KeyValueStoreCall) GetValues() (string, error) {
	c.buildClient()
	defer c.GetConn().Close()
	//请求
	request := &pb.ValuesRequest{Key: "key_1"}
	r, err := c.client.GetValues(c.ctx, request)
	if err != nil {
		return "", nil
	}
	data := r.GetValue()
	return data, nil
}
