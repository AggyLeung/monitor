import { Button, Card, Form, Input } from "antd";
import { useNavigate } from "react-router-dom";

function LoginPage() {
  const navigate = useNavigate();

  const submit = async (values: { username: string; password: string }) => {
    if (values.username) {
      localStorage.setItem("token", "dev-token");
      navigate("/resources");
    }
  };

  return (
    <div style={{ minHeight: "100vh", display: "grid", placeItems: "center" }}>
      <Card title="Sign In" style={{ width: 360 }}>
        <Form layout="vertical" onFinish={submit}>
          <Form.Item label="Username" name="username" rules={[{ required: true }]}>
            <Input />
          </Form.Item>
          <Form.Item label="Password" name="password" rules={[{ required: true }]}>
            <Input.Password />
          </Form.Item>
          <Button type="primary" htmlType="submit" block>
            Login
          </Button>
        </Form>
      </Card>
    </div>
  );
}

export default LoginPage;
