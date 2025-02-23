package cb

import (
	"fmt"

	"github.com/weibaohui/k8m/pkg/constants"
	"github.com/weibaohui/k8m/pkg/models"
	"github.com/weibaohui/k8m/pkg/service"
	"github.com/weibaohui/kom/kom"
	"k8s.io/klog/v2"
)

func RegisterCallback() {
	clusters := service.ClusterService().ConnectedClusters()
	for _, cluster := range clusters {
		selectedCluster := service.ClusterService().ClusterID(cluster)

		deleteCallback := kom.Cluster(selectedCluster).Callback().Delete()
		_ = deleteCallback.Before("kom:delete").Register("k8m:delete", handleDelete)
		updateCallback := kom.Cluster(selectedCluster).Callback().Update()
		_ = updateCallback.Before("kom:update").Register("k8m:update", handleUpdate)
		patchCallback := kom.Cluster(selectedCluster).Callback().Patch()
		_ = patchCallback.Before("kom:patch").Register("k8m:patch", handlePatch)
		createCallback := kom.Cluster(selectedCluster).Callback().Create()
		_ = createCallback.Before("kom:create").Register("k8m:create", handleCreate)
	}
}

func handleCommonLogic(k8s *kom.Kubectl, action string) (string, string, error) {
	stmt := k8s.Statement
	cluster := k8s.ID

	username := fmt.Sprintf("%s", stmt.Context.Value(constants.JwtUserName))
	role := fmt.Sprintf("%s", stmt.Context.Value(constants.JwtUserRole))
	klog.V(6).Infof("cb: cluster= %s,user= %s, role= %s, operation=%s, gck=[%s], resource=[%s/%s] ",
		cluster, username, role, action, stmt.GVK.String(), stmt.Namespace, stmt.Name)

	log := models.OperationLog{
		Action:    action,
		Cluster:   cluster,
		Kind:      stmt.GVK.Kind,
		Name:      stmt.Name,
		Namespace: stmt.Namespace,
		UserName:  username,
		Group:     stmt.GVK.Group,
		Role:      role,
	}
	go func() {
		service.OperationLogService().Add(&log)
	}()

	if role == models.RoleClusterReadonly {
		return username, role, fmt.Errorf("非管理员不能%s资源", action)
	}
	return username, role, nil
}

func handleDelete(k8s *kom.Kubectl) error {
	_, _, err := handleCommonLogic(k8s, "delete")
	return err
}

func handleUpdate(k8s *kom.Kubectl) error {
	_, _, err := handleCommonLogic(k8s, "update")
	return err
}

func handlePatch(k8s *kom.Kubectl) error {
	_, _, err := handleCommonLogic(k8s, "patch")
	return err
}

func handleCreate(k8s *kom.Kubectl) error {
	_, _, err := handleCommonLogic(k8s, "create")
	return err
}
