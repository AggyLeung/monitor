import { useMemo, useState } from "react";
import { Button, Card, Form, Input, Modal, Popconfirm, Select, Space, Table, Tag } from "antd";
import type { ColumnsType } from "antd/es/table";
import { Link } from "react-router-dom";
import DynamicAttributesForm from "../components/DynamicAttributesForm";
import { useAppDispatch, useAppSelector } from "../store/hooks";
import { addResource, softDeleteResource, updateResource } from "../store/slices/resourcesSlice";
import { Resource } from "../types";

interface FormValues {
  name: string;
  type: string;
  owner: string;
  ip: string;
  site: string;
  attributes?: Record<string, string | number>;
}

function ResourcesPage() {
  const dispatch = useAppDispatch();
  const resources = useAppSelector((state) => state.resources.items);
  const ciTypes = useAppSelector((state) => state.ciTypes.items.filter((it) => it.enabled));
  const [search, setSearch] = useState("");
  const [typeFilter, setTypeFilter] = useState<string>("all");
  const [open, setOpen] = useState(false);
  const [editing, setEditing] = useState<Resource | null>(null);
  const [selectedTypeName, setSelectedTypeName] = useState<string>("");
  const [form] = Form.useForm<FormValues>();

  const selectedCiType = useMemo(() => ciTypes.find((it) => it.name === selectedTypeName), [ciTypes, selectedTypeName]);
  const types = useMemo(() => ["all", ...Array.from(new Set(resources.map((r) => r.type)))], [resources]);

  const filtered = useMemo(() => {
    return resources.filter((item) => {
      if (typeFilter !== "all" && item.type !== typeFilter) {
        return false;
      }
      if (!search.trim()) {
        return true;
      }
      const kw = search.toLowerCase();
      return [item.name, item.id, item.owner, item.ip, item.site].join("|").toLowerCase().includes(kw);
    });
  }, [resources, search, typeFilter]);

  const submit = async () => {
    const values = await form.validateFields();
    if (editing) {
      dispatch(updateResource({ id: editing.id, patch: values }));
    } else {
      dispatch(addResource(values));
    }
    setOpen(false);
    setEditing(null);
    setSelectedTypeName("");
    form.resetFields();
  };

  const columns: ColumnsType<Resource> = [
    { title: "CI ID", dataIndex: "id", width: 140 },
    {
      title: "Name",
      dataIndex: "name",
      render: (_, row) => <Link to={`/resources/${row.id}`}>{row.name}</Link>
    },
    { title: "Type", dataIndex: "type" },
    { title: "Owner", dataIndex: "owner" },
    { title: "IP", dataIndex: "ip" },
    { title: "Site", dataIndex: "site" },
    {
      title: "Status",
      dataIndex: "status",
      render: (status) => (status === "active" ? <Tag color="green">active</Tag> : <Tag>deleted</Tag>)
    },
    {
      title: "Action",
      render: (_, row) => (
        <Space>
          <Button
            size="small"
            onClick={() => {
              setEditing(row);
              setSelectedTypeName(row.type);
              form.setFieldsValue({
                name: row.name,
                type: row.type,
                owner: row.owner,
                ip: row.ip,
                site: row.site,
                attributes: row.attributes
              });
              setOpen(true);
            }}
          >
            Edit
          </Button>
          <Popconfirm title="Soft delete this resource?" onConfirm={() => dispatch(softDeleteResource(row.id))}>
            <Button size="small" danger disabled={row.status === "deleted"}>
              Delete
            </Button>
          </Popconfirm>
        </Space>
      )
    }
  ];

  return (
    <Card title="Resource Management">
      <Space style={{ marginBottom: 12 }}>
        <Select value={typeFilter} style={{ width: 180 }} onChange={setTypeFilter} options={types.map((t) => ({ label: t, value: t }))} />
        <Input.Search allowClear placeholder="Search by name/ID/owner/IP/site" value={search} onChange={(e) => setSearch(e.target.value)} style={{ width: 320 }} />
        <Button
          type="primary"
          onClick={() => {
            setEditing(null);
            setSelectedTypeName("");
            form.resetFields();
            setOpen(true);
          }}
        >
          Add Resource
        </Button>
      </Space>

      <Table rowKey="id" columns={columns} dataSource={filtered} pagination={{ pageSize: 8 }} />

      <Modal title={editing ? "Edit Resource" : "Add Resource"} open={open} onOk={() => void submit()} onCancel={() => setOpen(false)}>
        <Form form={form} layout="vertical">
          <Form.Item label="Name" name="name" rules={[{ required: true }]}>
            <Input />
          </Form.Item>
          <Form.Item label="Type" name="type" rules={[{ required: true }]}>
            <Select
              options={ciTypes.map((t) => ({ label: t.name, value: t.name }))}
              onChange={(val) => {
                setSelectedTypeName(val);
                form.setFieldValue("attributes", {});
              }}
              showSearch
            />
          </Form.Item>
          <Form.Item label="Owner" name="owner" rules={[{ required: true }]}>
            <Input />
          </Form.Item>
          <Form.Item label="IP" name="ip" rules={[{ required: true }]}>
            <Input />
          </Form.Item>
          <Form.Item label="Site" name="site" rules={[{ required: true }]}>
            <Input />
          </Form.Item>
          <DynamicAttributesForm ciType={selectedCiType} />
        </Form>
      </Modal>
    </Card>
  );
}

export default ResourcesPage;
