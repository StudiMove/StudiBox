import React from 'react';
import Select from 'react-select';
import makeAnimated from 'react-select/animated';

const animatedComponents = makeAnimated();

interface InputFieldMultiSelectProps {
  label: string;
  value: { value: string; label: string }[]; // Valeurs sélectionnées
  options: { value: string; label: string }[]; // Liste d'options
  onChange: (selectedValues: { value: string; label: string }[]) => void; // Gestionnaire de changement
  isEditable?: boolean; // Champ éditable par défaut
}

const customStyles = {
  control: (provided: any) => ({
    ...provided,
    height: '3.5rem', // h-14 en Tailwind équivaut à 3.5rem
    borderColor: '#a0aec0',
    boxShadow: 'none',
    borderRadius: '0.75rem',
    '&:hover': {
      borderColor: '#3182ce',
    },
    '@media (max-width: 768px)': {
      minHeight: 'auto', // On smaller screens, let the height adjust dynamically
    },
  }),
  multiValue: (provided: any) => ({
    ...provided,
    backgroundColor: '#E7EEF8',
    display: 'flex',
    flexWrap: 'wrap', // Allow wrapping on mobile
  }),
  multiValueLabel: (provided: any) => ({
    ...provided,
    color: '#0F52BA',
    fontFamily: 'Helvetica Neue, Arial, sans-serif',
    '@media (max-width: 768px)': {
      fontSize: '12px', // Adjust font size on mobile
    },
  }),
  multiValueRemove: (provided: any) => ({
    ...provided,
    color: '#7BA0D9',
    ':hover': {
      backgroundColor: '#edf2f7',
      color: '#e53e3e',
    },
  }),
};

const InputFieldMultiSelect: React.FC<InputFieldMultiSelectProps> = ({
  label,
  value,
  options,
  onChange,
  isEditable = true,
}) => {
  const handleChange = (
    selectedOptions: { value: string; label: string }[],
  ) => {
    onChange(selectedOptions);
  };

  return (
    <div className="flex flex-col relative flex-1">
      <label className="mb-3 font-medium text-left text-lightBlack">
        {label}
      </label>
      <Select
        closeMenuOnSelect={false}
        components={animatedComponents}
        isDisabled={!isEditable}
        isMulti
        options={options}
        value={value}
        onChange={handleChange}
        placeholder="Veuillez sélectionner..."
        classNamePrefix="react-select"
        styles={customStyles}
      />
    </div>
  );
};

export default InputFieldMultiSelect;
