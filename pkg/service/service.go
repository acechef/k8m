package service

var localPodService = &podService{}
var localChatService = &chatService{}
var localNodeService = &nodeService{}
var localDeploymentService = &deployService{}
var localClusterService = &clusterService{
	clusterConfigs:        []*ClusterConfig{},
	AggregateDelaySeconds: 61, // 没有秒级支持，所以大于1分钟
}
var localStorageClassService = &storageClassService{}
var localIngressClassService = &ingressClassService{}
var localPVCService = &pvcService{
	CountList: []*pvcCount{},
}
var localPVService = &pvService{
	CountList: []*pvCount{},
}
var localIngressService = &ingressService{
	CountList: []*ingressCount{},
}
var localUserService = &userService{}
var localOperationLogService = &operationLogService{}
var localShellLogService = &shellLogService{}

func ChatService() *chatService {
	return localChatService
}
func DeploymentService() *deployService {
	return localDeploymentService
}
func PodService() *podService {
	return localPodService
}
func NodeService() *nodeService {
	return localNodeService
}
func ClusterService() *clusterService {
	return localClusterService
}
func StorageClassService() *storageClassService {
	return localStorageClassService
}
func IngressClassService() *ingressClassService {
	return localIngressClassService
}
func PVCService() *pvcService {
	return localPVCService
}
func PVService() *pvService {
	return localPVService
}
func IngressService() *ingressService {
	return localIngressService
}

func UserService() *userService {
	return localUserService
}

func OperationLogService() *operationLogService {
	return localOperationLogService
}
func ShellLogService() *shellLogService {
	return localShellLogService
}
