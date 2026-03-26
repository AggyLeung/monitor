import { Button, Card, Form, Input, Modal, Select, Space, Switch, Table, Tag } from "antd";
import type { ColumnsType } from "antd/es/table";
import { useState } from "react";
import { useAppDispatch, useAppSelector } from "../store/hooks";
import { addCiType, toggleCiType } from "../store/slices/ciTypesSlice";
import { CiAttributeSchema, CiType } from "../types";

interface FormValues {
  name: string;
  attributes: string;
}

function parseAttributeTemplate(raw: string): CiAttributeSchema[] {
  return raw
    .split(",")
    .map((item) => item.trim())
    .filter(Boolean)
    .map((part) => {
      const [namePart, typePart] = part.split(":").map((v) => v.trim());
      if (!namePart) {
        return null;
      }
      const type = typePart === "int" || typePart === "enum" ? typePart : "string";
      return {
        name: namePart,
        label: namePart,
        type
      } as CiAttributeSchema;
    })
    .filter((v): v is CiAttributeSchema => Boolean(v));
}

function CiTypesPage() {
  const dispatch = useAppDispatch();
  const items = useAppSelector((state) => state.ciTypes.items);
  const [open, setOpen] = useState(false);
  const [form] = Form.useForm<FormValues>();

  const columns: ColumnsType<CiType> = [
    { title: "Type ID", dataIndex: "id", width: 180 },
    { title: "Type Name", dataIndex: "name", width: 180 },
    {
      title: "Attribute Template",
      dataIndex: "attributes",
      render: (attrs: CiAttributeSchema[]) => (
        <Space wrap>
          {attrs.map((attr) => (
            <Tag key={attr.name}>{`${attr.name}:${attr.type}`}</Tag>
          ))}
        </Space>
      )
    },
    {
      title: "Enabled",
      dataIndex: "enabled",
      width: 120,
      render: (_, row) => (
        <Switch checked={row.enabled} onChange={(checked) => dispatch(toggleCiType({ id: row.id, enabled: checked }))} />
      )
    }
  ];

  const submit = async () => {
    const values = await form.validateFields();
    const attributes = parseAttributeTemplate(values.attributes);
    dispatch(addCiType({ name: values.name, attributes }));
    setOpen(false);
    form.resetFields();
  };

  return (
    <Card title="CI Type Management">
      <Button type="primary" onClick={() => setOpen(true)} style={{ marginBottom: 12 }}>
        Add CI Type
      </Button>
      <Table rowKey="id" columns={columns} dataSource={items} pagination={false} />
      <Modal title="Add CI Type" open={open} onOk={() => void submit()} onCancel={() => setOpen(false)}>
        <Form layout="vertical" form={form}>
          <Form.Item label="Type Name" name="name" rules={[{ required: true }]}>
            <Input placeholder="Database" />
          </Form.Item>
          <Form.Item
            label="Attributes"
            name="attributes"
            rules={[{ required: true }]}
            extra="Format: field:type,field:type. Types: string | int | enum"
          >
            <Input placeholder="host:string,port:int,engine:enum" />
          </Form.Item>
          <Form.Item label="Type Examples (optional)">
            <Select
              options={[
                { label: "Server", value: "cpu:int,memory:int,ip:string,os:string" },
                { label: "Firewall", value: "throughput:int,policy_set:string,mode:enum" }
              ]}
              onChange={(value) => form.setFieldValue("attributes", value)}
              allowClear
            />
          </Form.Item>
        </Form>
      </Modal>
    </Card>
  );
}

export default CiTypesPage;
