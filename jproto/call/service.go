package call

import (
	"fmt"
	pb "github.com/ijidan/jproto/jproto/build"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
)

//service
const Host = "127.0.0.1"
const Port = int64(50052)

//consul
const ConsulHost = "192.168.33.10"
const ConsulPort = int64(8500)

//HTTP
const HttpHost = "127.0.0.1"
const HttpPort = int64(8080)

//获取地址
func getAddress() string {
	address := fmt.Sprintf("%s:%d", Host, Port)
	return address
}

//注册服务
func registerServices(server *grpc.Server) {
	//hello服务
	hello := NewHelloService()
	helloName := hello.GetName()
	pb.RegisterHelloServer(server, hello)
	register2ETCD(helloName)
	//register2Consul("id_"+helloName, helloName, nil)

	//KV服务
	keyValue := NewKeyValueStoreService()
	keyValueName := keyValue.GetName()
	pb.RegisterKeyValueStoreServer(server, keyValue)
	register2ETCD(keyValueName)

	//简单模式
	simple := NewSimpleService()
	simpleName := simple.GetName()
	pb.RegisterSimpleServer(server, simple)
	register2ETCD(simpleName)

	//服务端流
	sr := NewServerStreamService()
	srName := sr.GetName()
	pb.RegisterServerStreamServer(server, sr)
	register2ETCD(srName)

	//客户端流
	client := NewClientStreamService()
	clientName := client.GetName()
	pb.RegisterClientStreamServer(server, client)
	register2ETCD(clientName)

	//双向流
	both := NewBothStreamService()
	bothName := both.GetName()
	pb.RegisterBothStreamServer(server, both)
	register2ETCD(bothName)

}

//注册到Consul
//func register2Consul(serviceId string, serviceName string, serviceTags []string) {
//	c := jconsul.NewJConsul(ConsulHost, ConsulPort)
//	_ = c.ServiceRegister(serviceId, serviceName, serviceTags, Host, Port, HttpHost, HttpPort, "/service/consulCheck")
//}

//注册到ETCD
func register2ETCD(serviceName string) {
	//
	//var endpoints = []string{jetcd.EtcdUrl,}
	//var key = "/service/" + serviceName
	//address:=getAddress()
	//sr := jetcd.NewServiceRegister(endpoints, 5, key, address)
	////监听租约
	//go func() {
	//	_ = sr.ListLeaseKeepAliveChan()
	//}()
}

//获取consul 服务列表
//func GetConsulServiceList() {
//	time.Sleep(time.Minute)
//	c := jconsul.NewJConsul(ConsulHost, ConsulPort)
//	serviceList, err := c.ServiceDiscovery()
//	log.Println(serviceList)
//	log.Println(err)
//
//}


//服务开始
func StartCallServer() {
	server := grpc.NewServer(
		////流式拦截器
		//grpc.StreamInterceptor(
		//	grpcMiddleware.ChainStreamServer(
		//	grpcCtxTags.StreamServerInterceptor(),
		//	grpcOpenTracing.StreamServerInterceptor(),
		//	grpcZap.StreamServerInterceptor(GetZapLogger()),
		//	//grpc_auth.StreamServerInterceptor(myAuthFunction),
		//	grpcRecovery.StreamServerInterceptor(),
		//)),
		////普通拦截器
		//grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
		//	grpcCtxTags.UnaryServerInterceptor(),
		//	grpcOpenTracing.UnaryServerInterceptor(),
		//	grpcZap.UnaryServerInterceptor(GetZapLogger()),
		//	//grpc_auth.UnaryServerInterceptor(myAuthFunction),
		//	grpcRecovery.UnaryServerInterceptor(),
		//)),
	)
	//注册服务
	registerServices(server)
	//监听
	address := getAddress()
	grpclog.Infoln("listen on :" + address)
	listen, err := net.Listen("tcp", address)
	if err != nil {
		grpclog.Fatalln("failed to listen：", err)
	}
	_ = server.Serve(listen)
}
