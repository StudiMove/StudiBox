import React from 'react';
import { InputFieldProps } from '../types/InputField/InputFieldProps';

const sizeClasses = {
  small: 'py-1 px-2 text-sm',
  medium: 'py-2 px-3 text-base',
  large: 'py-3 px-4 text-lg',
};

function InputField({
  type,
  label,
  value,
  size = 'medium',
  onChange,
}: InputFieldProps) {
  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    onChange(event.target.value);
  };

  return (
    <div className={`flex flex-col mb-4 ${sizeClasses[size]}`}>
      <label className="mb-1 font-medium">{label}</label>
      <input
        type={type}
        value={value}
        onChange={handleChange}
        className="border border-gray-300 p-2 rounded"
      />
    </div>
  );
}

export default InputField;
