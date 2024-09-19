// src/models/InputField/InputFieldModel.ts

import {
  InputFieldType,
  InputFieldSize,
} from '../../types/InputField/InputFieldType';

export interface InputFieldModel {
  type: InputFieldType;
  label: string;
  value: string | number;
  size?: InputFieldSize;
}
