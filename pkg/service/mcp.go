package service

import (
	"context"
	"time"

	"github.com/weibaohui/k8m/internal/dao"
	"github.com/weibaohui/k8m/pkg/mcp"
	"github.com/weibaohui/k8m/pkg/models"
	"k8s.io/klog/v2"
)

type mcpService struct {
	host *mcp.MCPHost
}

func (m *mcpService) Init() {
	if m.host == nil {
		m.host = mcp.NewMCPHost()
	}
	m.Start()
}
func (m *mcpService) Host() *mcp.MCPHost {
	return m.host
}
func (m *mcpService) AddServer(server models.MCPServerConfig) {
	// 将server转换为mcp.ServerConfig
	serverConfig := mcp.ServerConfig{
		ID:      server.ID,
		Name:    server.Name,
		URL:     server.URL,
		Enabled: server.Enabled,
	}
	err := m.host.AddServer(serverConfig)
	if err != nil {
		klog.V(6).Infof("Failed to add server %s: %v", server.Name, err)
		return
	}

	if server.Enabled {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		err := m.host.ConnectServer(ctx, server.Name)
		if err != nil {
			klog.V(6).Infof("Failed to connect to server %s: %v", server.Name, err)
			return
		}
		klog.V(6).Infof("Successfully connected to server: %s", server.Name)
	}

}
func (m *mcpService) AddServers(servers []models.MCPServerConfig) {
	for _, server := range servers {
		// 将server转换为mcp.ServerConfig
		serverConfig := mcp.ServerConfig{
			ID:      server.ID,
			Name:    server.Name,
			URL:     server.URL,
			Enabled: server.Enabled,
		}
		err := m.host.AddServer(serverConfig)
		if err != nil {
			klog.V(6).Infof("Failed to add server %s: %v", server.Name, err)
			continue
		}

		if server.Enabled {
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			err := m.host.ConnectServer(ctx, server.Name)
			cancel()
			if err != nil {
				klog.V(6).Infof("Failed to connect to server %s: %v", server.Name, err)
				continue
			}
			klog.V(6).Infof("Successfully connected to server: %s", server.Name)
		}
	}

}
func (m *mcpService) RemoveServer(server models.MCPServerConfig) {
	// 将server转换为mcp.ServerConfig
	serverConfig := mcp.ServerConfig{
		Name:    server.Name,
		URL:     server.URL,
		Enabled: server.Enabled,
	}
	m.host.RemoveServer(serverConfig)
}
func (m *mcpService) Start() {

	var mcpServers []models.MCPServerConfig
	err := dao.DB().Model(&models.MCPServerConfig{}).Find(&mcpServers).Error
	if err != nil {
		return
	}
	m.AddServers(mcpServers)
}

func (m *mcpService) RemoveServerById(server models.MCPServerConfig) {
	m.host.RemoveServerById(server.ID)
}

func (m *mcpService) UpdateServer(entity models.MCPServerConfig) {
	m.RemoveServerById(entity)
	m.AddServer(entity)
}
