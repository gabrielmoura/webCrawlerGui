import { ChangeEvent, Dispatch, SetStateAction } from "react";
import { FormControl, Input, Text } from "@chakra-ui/react";

interface InputConfigProps {
  label: string;
  value: string | number;
  type?: string;
  onChange: Dispatch<SetStateAction<any>>;
}

export function InputConfig({
  value,
  label,
  onChange,
  type,
}: InputConfigProps) {
  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    const newValue =
      type === "number" ? parseInt(e.target.value) : e.target.value;
    onChange(newValue);
  };

  return (
    <FormControl alignItems="center">
      <Text fontSize="2xl">{label}</Text>
      <Input
        defaultValue={value}
        type={type ?? "text"}
        onChange={handleChange}
      />
    </FormControl>
  );
}
