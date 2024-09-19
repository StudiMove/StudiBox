// src/types/InputField/InputFieldProps.d.ts

import { InputFieldType, InputFieldSize } from './InputFieldType';

export interface InputFieldProps {
  type: InputFieldType;
  label: string;
  value: string | number;
  size?: InputFieldSize;
  onChange: (value: string | number) => void;
}
