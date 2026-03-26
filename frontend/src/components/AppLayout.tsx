import { ApartmentOutlined, DatabaseOutlined, DeploymentUnitOutlined, TeamOutlined, UnorderedListOutlined } from "@ant-design/icons";
import { Layout, Menu, Typography } from "antd";
import { Outlet, useLocation, useNavigate } from "react-router-dom";

const { Header, Sider, Content } = Layout;

const menuItems = [
  { key: "/resources", icon: <DatabaseOutlined />, label: "资源管理" },
  { key: "/ci-types", icon: <UnorderedListOutlined />, label: "类型管理" },
  { key: "/tasks", icon: <DeploymentUnitOutlined />, label: "任务管理" },
  { key: "/users", icon: <TeamOutlined />, label: "用户权限" },
  { key: "/topology/ci-001", icon: <ApartmentOutlined />, label: "关系拓扑" }
];

function AppLayout() {
  const navigate = useNavigate();
  const location = useLocation();

  return (
    <Layout style={{ minHeight: "100vh" }}>
      <Sider width={240} theme="light" style={{ borderRight: "1px solid #e8edf2" }}>
        <div className="brand">NetVision OPS</div>
        <Menu
          mode="inline"
          selectedKeys={[location.pathname.startsWith("/resources/") ? "/resources" : location.pathname]}
          items={menuItems}
          onClick={(item) => navigate(item.key)}
        />
      </Sider>
      <Layout>
        <Header className="header">
          <Typography.Title level={4} style={{ margin: 0 }}>
            企业网络监控系统
          </Typography.Title>
        </Header>
        <Content className="content">
          <Outlet />
        </Content>
      </Layout>
    </Layout>
  );
}

export default AppLayout;
