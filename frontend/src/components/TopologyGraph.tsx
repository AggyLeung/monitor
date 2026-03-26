import { useEffect, useMemo, useRef } from "react";
import { Button, Space, message } from "antd";
import { Relation } from "../types";

interface TopologyGraphProps {
  rootId: string;
  relations: Relation[];
  onDrill: (targetId: string) => void;
}

function TopologyGraph({ rootId, relations, onDrill }: TopologyGraphProps) {
  const containerRef = useRef<HTMLDivElement | null>(null);
  const graphRef = useRef<any>(null);

  const nodes = useMemo(() => {
    const nodeSet = new Set<string>();
    nodeSet.add(rootId);
    relations.forEach((r) => {
      nodeSet.add(r.source);
      nodeSet.add(r.target);
    });
    return Array.from(nodeSet).map((id) => ({
      id,
      data: { label: id }
    }));
  }, [relations, rootId]);

  useEffect(() => {
    let disposed = false;

    const render = async () => {
      if (!containerRef.current) {
        return;
      }
      try {
        const g6 = await import("@antv/g6");
        if (disposed) {
          return;
        }
        graphRef.current?.destroy?.();
        const graph = new g6.Graph({
          container: containerRef.current,
          autoFit: "view",
          data: {
            nodes,
            edges: relations.map((edge) => ({
              source: edge.source,
              target: edge.target,
              data: { label: edge.label ?? "" }
            }))
          },
          node: {
            type: "circle",
            style: {
              size: 40,
              fill: "#e8f3f6",
              stroke: "#1f7a8c",
              labelText: (d: any) => d.id
            }
          },
          edge: {
            type: "line",
            style: {
              stroke: "#8aa6b0",
              endArrow: true
            }
          },
          layout: {
            type: "force",
            linkDistance: 150
          },
          behaviors: ["drag-canvas", "zoom-canvas", "drag-element"]
        });

        graph.on("node:click", (evt: any) => {
          const id = evt?.target?.id;
          if (id) {
            onDrill(id);
          }
        });

        await graph.render();
        graphRef.current = graph;
      } catch (error) {
        message.error("拓扑渲染失败，请检查 G6 依赖");
      }
    };

    void render();

    return () => {
      disposed = true;
      graphRef.current?.destroy?.();
      graphRef.current = null;
    };
  }, [nodes, onDrill, relations]);

  const exportAsJson = () => {
    const payload = {
      rootId,
      relations
    };
    const blob = new Blob([JSON.stringify(payload, null, 2)], { type: "application/json" });
    const url = URL.createObjectURL(blob);
    const a = document.createElement("a");
    a.href = url;
    a.download = `topology-${rootId}.json`;
    a.click();
    URL.revokeObjectURL(url);
  };

  return (
    <div>
      <Space style={{ marginBottom: 12 }}>
        <Button onClick={exportAsJson}>导出</Button>
      </Space>
      <div ref={containerRef} className="topology-container" />
    </div>
  );
}

export default TopologyGraph;
