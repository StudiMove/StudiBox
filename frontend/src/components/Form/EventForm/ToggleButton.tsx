interface ToggleButtonProps {
  label: string;
  value: boolean;
  onChange: (value: boolean) => void;
}

const ToggleButton = ({ label, value, onChange }: ToggleButtonProps) => {
  return (
    <div className="flex flex-col items-start">
      {/* Label au-dessus du toggle */}
      <label className="text-sm font-medium text-lightBlack mb-2">
        {label}
      </label>
      <label className="inline-flex items-center cursor-pointer">
        <input
          type="checkbox"
          checked={value}
          onChange={(e) => onChange(e.target.checked)}
          className="sr-only peer"
          aria-label={label}
        />
        <div
          className={`relative w-14 h-7 bg-gray-200 rounded-full 
          transition-all duration-300 
          ${value ? 'bg-primary' : 'bg-gray-200'} 
          peer-focus:outline-none 
          peer-focus:ring-4 peer-focus:ring-blue-300`}
        >
          <div
            className={`absolute top-0.5 left-0.5 
            transition-transform duration-300 
            ${value ? 'translate-x-7' : 'translate-x-0'} 
            bg-white border border-gray-300 rounded-full h-6 w-6`}
          />
        </div>
      </label>
    </div>
  );
};

export default ToggleButton;
