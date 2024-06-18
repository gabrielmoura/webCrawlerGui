import { ChangeEvent, Dispatch, SetStateAction } from "react";
import { FormControl, FormLabel, Switch } from "@chakra-ui/react";

interface SwitchConfigProps {
  label: string;
  value: boolean;
  name: string;
  onChange: Dispatch<SetStateAction<boolean>>;
}

export function SwitchConfig({
  value,
  label,
  onChange,
  name,
}: SwitchConfigProps) {
  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    onChange(e.target.checked);
  };

  return (
    <FormControl alignItems="center">
      <FormLabel fontSize="2xl" htmlFor={name} mb="0">
        {label}
      </FormLabel>
      <Switch id={name} isChecked={value} size="lg" onChange={handleChange} />
    </FormControl>
  );
}
