package call

import (
	"context"
	pb "github.com/ijidan/jproto/jproto/build"
)

//键值对存储服务
type KeyValueStoreService struct {
	BasicService
}

//获取服务名称
func (s *KeyValueStoreService) GetName() string {
	return "key_value"
}

//获取值
func (s *KeyValueStoreService) GetValues(ctx context.Context, req *pb.ValuesRequest) (*pb.ValuesResponse, error) {
	rsp := new(pb.ValuesResponse)
	rsp.Value = "当前值：" + req.Key
	return rsp, nil
}

//获取实例
func NewKeyValueStoreService() *KeyValueStoreService {
	instance := &KeyValueStoreService{}
	return instance
}
