import { useEffect, useRef } from "react";
import { Card, Table, Tag } from "antd";
import type { ColumnsType } from "antd/es/table";
import * as echarts from "echarts";
import { SyncTask } from "../types";

const tasks: SyncTask[] = [
  {
    id: "task-1001",
    name: "nightly-cloud-sync",
    status: "success",
    startedAt: "2026-03-26T01:00:00Z",
    durationSec: 142,
    log: "Scanned 124 instances. Upsert success=124."
  },
  {
    id: "task-1002",
    name: "on-demand-network-scan",
    status: "running",
    startedAt: "2026-03-26T15:15:00Z",
    durationSec: 41,
    log: "Scanning site=Shanghai-A..."
  },
  {
    id: "task-1003",
    name: "agent-delta-ingest",
    status: "failed",
    startedAt: "2026-03-26T13:05:00Z",
    durationSec: 21,
    log: "401 from upstream token endpoint."
  }
];

function TasksPage() {
  const chartRef = useRef<HTMLDivElement | null>(null);

  useEffect(() => {
    if (!chartRef.current) {
      return;
    }
    const chart = echarts.init(chartRef.current);
    const stats = {
      success: tasks.filter((t) => t.status === "success").length,
      running: tasks.filter((t) => t.status === "running").length,
      failed: tasks.filter((t) => t.status === "failed").length
    };
    chart.setOption({
      tooltip: { trigger: "item" },
      series: [
        {
          type: "pie",
          radius: ["45%", "70%"],
          data: [
            { name: "success", value: stats.success },
            { name: "running", value: stats.running },
            { name: "failed", value: stats.failed }
          ]
        }
      ]
    });
    const onResize = () => chart.resize();
    window.addEventListener("resize", onResize);
    return () => {
      window.removeEventListener("resize", onResize);
      chart.dispose();
    };
  }, []);

  const columns: ColumnsType<SyncTask> = [
    { title: "任务 ID", dataIndex: "id", width: 140 },
    { title: "任务名", dataIndex: "name", width: 220 },
    {
      title: "状态",
      dataIndex: "status",
      render: (status) => {
        const color = status === "success" ? "green" : status === "running" ? "blue" : status === "failed" ? "red" : "default";
        return <Tag color={color}>{status}</Tag>;
      }
    },
    { title: "开始时间", dataIndex: "startedAt" },
    { title: "耗时(s)", dataIndex: "durationSec", width: 100 },
    { title: "日志", dataIndex: "log" }
  ];

  return (
    <Card title="任务管理">
      <div ref={chartRef} style={{ width: "100%", height: 220, marginBottom: 16 }} />
      <Table rowKey="id" dataSource={tasks} columns={columns} pagination={false} />
    </Card>
  );
}

export default TasksPage;
