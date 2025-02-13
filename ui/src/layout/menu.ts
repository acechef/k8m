export default [
    {
        label: "首页",
        icon: "fa-solid fa-house",
        path: "/",
        key: "home",
    },
    {
        label: "多集群",
        icon: "fa-solid fa-server",
        path: "/cluster/cluster_all",
        key: "cluster_all",
    },
    {
        label: "创建",
        icon: "fa-solid fa-dharmachakra",
        path: "/apply/history",
        key: "apply",
    },
    {
        label: "命名空间",
        icon: "fa-solid fa-border-style",
        path: "/cluster/ns",
        key: "cluster_ns",
    },
    {
        label: "节点",
        icon: "fa-solid fa-computer",
        path: "/cluster/node",
        key: "cluster_node",
    },
    {
        label: "事件",
        icon: "fa-solid fa-bell",
        path: "/ns/event",
        key: "event",
    },
    {
        label: "工作负载",
        icon: "fa-solid fa-cube",
        key: "workload",
        children: [
            {
                label: "部署",
                icon: "fa-solid fa-layer-group",
                path: "/ns/deploy",
                key: "deploy",
            },
            {
                label: "有状态集",
                icon: "fa-solid fa-database",
                path: "/ns/statefulset",
                key: "statefulset",
            },
            {
                label: "守护进程集",
                icon: "fa-solid fa-shield-halved",
                path: "/ns/daemonset",
                key: "daemonset",
            },
            {
                label: "任务",
                icon: "fa-solid fa-list-check",
                path: "/ns/job",
                key: "job",
            },
            {
                label: "定时任务",
                icon: "fa-solid fa-clock",
                path: "/ns/cronjob",
                key: "cronjob",
            },
            {
                label: "容器组",
                icon: "fa-solid fa-cubes",
                path: "/ns/pod",
                key: "pod",
            },
            {
                label: "副本集",
                icon: "fa-solid fa-clone",
                path: "/ns/replicaset",
                key: "replicaset",
            },
        ],
    },
    {
        label: "CRD",
        icon: "fa-solid fa-file-code",
        key: "crd",
        children: [
            {
                label: "自定义资源",
                icon: "fa-solid fa-gears",
                path: "/crd/crd",
                key: "custom_resource",
            }
        ],
    },
    {
        label: "配置",
        icon: "fa-solid fa-sliders-h",
        key: "config",
        children: [
            {
                label: "配置映射",
                icon: "fa-solid fa-map",
                path: "/ns/configmap",
                key: "configmap",
            },
            {
                label: "密钥",
                icon: "fa-solid fa-key",
                path: "/ns/secret",
                key: "secret",
            },
        ],
    },
    {
        label: "网络",
        icon: "fa-solid fa-network-wired",
        key: "network",
        children: [
            {
                label: "SVC服务",
                icon: "fa-solid fa-project-diagram",
                path: "/ns/svc",
                key: "svc",
            },
            {
                label: "Ingress入口",
                icon: "fa-solid  fa-wifi",
                path: "/ns/ing",
                key: "ingress",
            },
        ],
    },
    {
        label: "存储",
        icon: "fa-solid fa-memory",
        key: "storage",
        children: [
            {
                label: "持久卷声明",
                icon: "fa-solid fa-folder",
                path: "/ns/pvc",
                key: "pvc",
            },
            {
                label: "持久卷",
                icon: "fa-solid fa-hdd",
                path: "/cluster/pv",
                key: "pv",
            },
            {
                label: "存储类",
                icon: "fa-solid fa-coins",
                path: "/cluster/storage_class",
                key: "storage_class",
            },
        ],
    },
    {
        label: "访问控制",
        icon: "fa-solid fa-lock",
        key: "access_control",
        children: [
            {
                label: "服务账户",
                icon: "fa-solid fa-user-shield",
                path: "/ns/service_account",
                key: "service_account",
            },
            {
                label: "角色",
                icon: "fa-solid fa-user-tag",
                path: "/ns/role",
                key: "role",
            },
            {
                label: "角色绑定",
                icon: "fa-solid fa-link",
                path: "/ns/role_binding",
                key: "role_binding",
            },
            {
                label: "集群角色",
                icon: "fa-solid fa-users",
                path: "/cluster/cluster_role",
                key: "cluster_role",
            },
            {
                label: "集群角色绑定",
                icon: "fa-solid fa-user-lock",
                path: "/cluster/cluster_role_binding",
                key: "cluster_role_binding",
            },
        ],
    },
    {
        label: "集群配置",
        icon: "fa-solid fa-cog",
        key: "cluster_config",
        children: [
            {
                label: "Ingress入口类",
                icon: "fa-solid fa-sitemap",
                path: "/cluster/ingress_class",
                key: "ingress_class",
            },
            {
                label: "验证钩子",
                icon: "fa-solid fa-check",
                path: "/cluster/validation_webhook",
                key: "validation_webhook",
            },
            {
                label: "变更钩子",
                icon: "fa-solid fa-exchange",
                path: "/cluster/mutating_webhook",
                key: "mutating_webhook",
            },
            {
                label: "端点",
                icon: "fa-solid fa-ethernet",
                path: "/ns/endpoint",
                key: "endpoint",
            },
            {
                label: "端点切片",
                icon: "fa-solid fa-newspaper",
                path: "/ns/endpointslice",
                key: "endpointslice",
            },
            {
                label: "水平自动扩缩",
                icon: "fa-solid fa-arrows-left-right",
                path: "/ns/hpa",
                key: "hpa",
            },
            {
                label: "网络策略",
                icon: "fa-solid fa-project-diagram",
                path: "/ns/network_policy",
                key: "network_policy",
            },
            {
                label: "资源配额",
                icon: "fa-solid fa-chart-pie",
                path: "/ns/resource_quota",
                key: "resource_quota",
            },
            {
                label: "限制范围",
                icon: "fa-solid fa-compress",
                path: "/ns/limit_range",
                key: "limit_range",
            },
            {
                label: "Pod中断配置",
                icon: "fa-solid fa-receipt",
                path: "/ns/pdb",
                key: "pdb",
            },
            {
                label: "租约",
                icon: "fa-solid fa-file-contract",
                path: "/ns/lease",
                key: "lease",
            },
            {
                label: "优先级类",
                icon: "fa-solid fa-sort",
                path: "/cluster/priority_class",
                key: "priority_class",
            },
            {
                label: "运行时类",
                icon: "fa-solid fa-play",
                path: "/cluster/runtime_class",
                key: "runtime_class",
            },
            {
                label: "CSI节点",
                icon: "fa-solid fa-server",
                path: "/cluster/csi_node",
                key: "csi_node",
            },
            {
                label: "API 服务",
                icon: "fa-solid fa-code",
                path: "/cluster/api_service",
                key: "api_service",
            },
            {
                label: "流量规则",
                icon: "fa-solid fa-random",
                path: "/cluster/flow_schema",
                key: "flow_schema",
            },
            {
                label: "优先级配置",
                icon: "fa-solid fa-sliders",
                path: "/cluster/priority_level_config",
                key: "priority_level_config",
            },
            {
                label: "组件状态",
                icon: "fa-solid fa-info-circle",
                path: "/cluster/component_status",
                key: "component_status",
            },
        ],
    },
];
