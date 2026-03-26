import { Card } from "antd";
import { useParams, useNavigate } from "react-router-dom";
import TopologyGraph from "../components/TopologyGraph";
import { Relation } from "../types";

const relationMap: Record<string, Relation[]> = {
  "ci-001": [
    { source: "edge-router-01", target: "wan-fw-01", label: "route-through" },
    { source: "wan-fw-01", target: "core-switch-01", label: "uplink" },
    { source: "core-switch-01", target: "k8s-gateway-01", label: "service-link" }
  ]
};

function TopologyPage() {
  const { id = "ci-001" } = useParams();
  const navigate = useNavigate();
  const relations = relationMap[id] ?? relationMap["ci-001"];
  return (
    <Card title={`关系拓扑 - ${id}`} style={{ minHeight: "calc(100vh - 140px)" }}>
      <TopologyGraph rootId={id} relations={relations} onDrill={(targetId) => navigate(`/topology/${targetId}`)} />
    </Card>
  );
}

export default TopologyPage;
