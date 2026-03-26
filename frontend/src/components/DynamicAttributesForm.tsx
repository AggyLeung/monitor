import { Form, Input, InputNumber, Select } from "antd";
import { CiType } from "../types";

interface DynamicAttributesFormProps {
  ciType?: CiType;
}

function DynamicAttributesForm({ ciType }: DynamicAttributesFormProps) {
  if (!ciType || ciType.attributes.length === 0) {
    return null;
  }

  return (
    <>
      {ciType.attributes.map((attr) => {
        const fieldName = ["attributes", attr.name];
        if (attr.type === "int") {
          return (
            <Form.Item key={attr.name} label={attr.label} name={fieldName}>
              <InputNumber style={{ width: "100%" }} />
            </Form.Item>
          );
        }
        if (attr.type === "enum") {
          return (
            <Form.Item key={attr.name} label={attr.label} name={fieldName}>
              <Select options={(attr.options ?? []).map((opt) => ({ label: opt, value: opt }))} />
            </Form.Item>
          );
        }
        return (
          <Form.Item key={attr.name} label={attr.label} name={fieldName}>
            <Input />
          </Form.Item>
        );
      })}
    </>
  );
}

export default DynamicAttributesForm;
