// src/types/InputFieldProps.d.ts

import { InputFieldType } from '../types/InputField/InputFieldType';

export interface InputFieldProps {
  label: string;
  value: string | number;
  onChange: (value: string | number) => void;
  size?: 'small' | 'medium' | 'large';
  type: InputFieldType;
}
