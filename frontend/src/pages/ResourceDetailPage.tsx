import { useMemo } from "react";
import { Button, Card, Col, Descriptions, Row, Table, Tag } from "antd";
import type { ColumnsType } from "antd/es/table";
import { Link, useNavigate, useParams } from "react-router-dom";
import TopologyGraph from "../components/TopologyGraph";
import { useAppSelector } from "../store/hooks";
import { ChangeRecord, Relation } from "../types";

const sampleChanges: ChangeRecord[] = [
  {
    id: "chg-1",
    field: "owner",
    before: "NOC-B",
    after: "NOC-A",
    operator: "admin",
    at: "2026-03-25T10:11:00Z"
  },
  {
    id: "chg-2",
    field: "policySet",
    before: "legacy-main",
    after: "corp-main",
    operator: "secops-admin",
    at: "2026-03-24T08:21:00Z"
  }
];

const sampleRelations: Relation[] = [
  { source: "edge-router-01", target: "wan-fw-01", label: "route-through" },
  { source: "wan-fw-01", target: "core-switch-01", label: "uplink" }
];

function ResourceDetailPage() {
  const navigate = useNavigate();
  const { id } = useParams();
  const resource = useAppSelector((state) => state.resources.items.find((item) => item.id === id));

  const attrRows = useMemo(
    () => Object.entries(resource?.attributes ?? {}).map(([k, v]) => ({ key: k, value: v })),
    [resource?.attributes]
  );

  const changeColumns: ColumnsType<ChangeRecord> = [
    { title: "字段", dataIndex: "field" },
    { title: "变更前", dataIndex: "before" },
    { title: "变更后", dataIndex: "after" },
    { title: "操作人", dataIndex: "operator" },
    { title: "时间", dataIndex: "at" }
  ];

  if (!resource) {
    return (
      <Card title="资源详情">
        资源不存在，返回 <Link to="/resources">资源列表</Link>
      </Card>
    );
  }

  return (
    <Row gutter={[16, 16]}>
      <Col span={24}>
        <Card
          title={`资源详情 - ${resource.name}`}
          extra={<Button onClick={() => navigate(`/topology/${resource.id}`)}>全屏查看拓扑</Button>}
        >
          <Descriptions bordered column={3}>
            <Descriptions.Item label="CI ID">{resource.id}</Descriptions.Item>
            <Descriptions.Item label="类型">{resource.type}</Descriptions.Item>
            <Descriptions.Item label="状态">
              {resource.status === "active" ? <Tag color="green">active</Tag> : <Tag>deleted</Tag>}
            </Descriptions.Item>
            <Descriptions.Item label="负责人">{resource.owner}</Descriptions.Item>
            <Descriptions.Item label="IP">{resource.ip}</Descriptions.Item>
            <Descriptions.Item label="站点">{resource.site}</Descriptions.Item>
            <Descriptions.Item label="最近更新时间" span={3}>
              {resource.updatedAt}
            </Descriptions.Item>
          </Descriptions>
        </Card>
      </Col>
      <Col span={12}>
        <Card title="属性详情">
          <Table
            pagination={false}
            dataSource={attrRows}
            columns={[
              { title: "属性", dataIndex: "key", key: "key" },
              { title: "值", dataIndex: "value", key: "value" }
            ]}
          />
        </Card>
      </Col>
      <Col span={12}>
        <Card title="关联关系拓扑">
          <TopologyGraph rootId={resource.name} relations={sampleRelations} onDrill={(target) => navigate(`/topology/${target}`)} />
        </Card>
      </Col>
      <Col span={24}>
        <Card title="变更历史">
          <Table rowKey="id" dataSource={sampleChanges} columns={changeColumns} pagination={false} />
        </Card>
      </Col>
    </Row>
  );
}

export default ResourceDetailPage;
