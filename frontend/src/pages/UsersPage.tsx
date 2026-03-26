import { useState } from "react";
import { Button, Card, Form, Input, Modal, Select, Space, Switch, Table } from "antd";
import type { ColumnsType } from "antd/es/table";
import { nanoid } from "@reduxjs/toolkit";
import { UserItem } from "../types";

function UsersPage() {
  const [items, setItems] = useState<UserItem[]>([
    { id: "u-001", username: "admin", role: "admin", email: "admin@corp.local", active: true },
    { id: "u-002", username: "alice", role: "operator", email: "alice@corp.local", active: true },
    { id: "u-003", username: "bob", role: "viewer", email: "bob@corp.local", active: false }
  ]);
  const [open, setOpen] = useState(false);
  const [form] = Form.useForm<{ username: string; role: UserItem["role"]; email: string }>();

  const columns: ColumnsType<UserItem> = [
    { title: "用户 ID", dataIndex: "id", width: 120 },
    { title: "用户名", dataIndex: "username", width: 160 },
    {
      title: "角色",
      dataIndex: "role",
      width: 160,
      render: (_, row) => (
        <Select
          value={row.role}
          style={{ width: 130 }}
          options={[
            { label: "admin", value: "admin" },
            { label: "operator", value: "operator" },
            { label: "viewer", value: "viewer" }
          ]}
          onChange={(role) => setItems((prev) => prev.map((it) => (it.id === row.id ? { ...it, role } : it)))}
        />
      )
    },
    { title: "邮箱", dataIndex: "email" },
    {
      title: "启用",
      dataIndex: "active",
      width: 100,
      render: (_, row) => (
        <Switch
          checked={row.active}
          onChange={(active) => setItems((prev) => prev.map((it) => (it.id === row.id ? { ...it, active } : it)))}
        />
      )
    }
  ];

  const submit = async () => {
    const values = await form.validateFields();
    setItems((prev) => [{ id: `u-${nanoid(5)}`, active: true, ...values }, ...prev]);
    setOpen(false);
    form.resetFields();
  };

  return (
    <Card title="用户权限管理">
      <Space style={{ marginBottom: 12 }}>
        <Button type="primary" onClick={() => setOpen(true)}>
          新增用户
        </Button>
      </Space>
      <Table rowKey="id" dataSource={items} columns={columns} pagination={false} />
      <Modal title="新增用户" open={open} onOk={() => void submit()} onCancel={() => setOpen(false)}>
        <Form layout="vertical" form={form}>
          <Form.Item label="用户名" name="username" rules={[{ required: true }]}>
            <Input />
          </Form.Item>
          <Form.Item label="邮箱" name="email" rules={[{ required: true }]}>
            <Input />
          </Form.Item>
          <Form.Item label="角色" name="role" rules={[{ required: true }]}>
            <Select
              options={[
                { label: "admin", value: "admin" },
                { label: "operator", value: "operator" },
                { label: "viewer", value: "viewer" }
              ]}
            />
          </Form.Item>
        </Form>
      </Modal>
    </Card>
  );
}

export default UsersPage;
