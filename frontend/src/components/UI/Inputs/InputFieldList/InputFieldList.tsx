// import React from 'react';

// interface InputFieldListProps {
//   label: string;
//   value: string | number;
//   options: { value: string | number; label: string }[]; // Liste d'options
//   onChange: (value: string | number) => void;
//   isEditable?: boolean; // Par défaut, le champ est éditable
// }

// const InputFieldList = ({
//   label,
//   value,
//   options,
//   onChange,
//   isEditable = true, // Par défaut, le champ est éditable
// }: InputFieldListProps) => {
//   const handleChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
//     const newValue =
//       typeof value === 'number'
//         ? Number(event.target.value)
//         : event.target.value;
//     onChange(newValue);
//   };

//   return (
//     <div className="flex flex-col relative flex-1">
//       <label className="mb-3 font-medium text-left text-lightBlack">
//         {label}
//       </label>
//       <div className="relative">
//         <select
//           value={value}
//           onChange={handleChange}
//           className="border border-lightGray rounded-xl w-full px-4 h-14 " // Taille par défaut équivalente à 'small'
//           disabled={!isEditable} // Désactive le champ si isEditable est false
//         >
//           {options.map((option) => (
//             <option key={option.value} value={option.value}>
//               {option.label}
//             </option>
//           ))}
//         </select>
//       </div>
//     </div>
//   );
// };

// export default InputFieldList;

import React from 'react';

interface InputFieldListProps {
  label: string;
  value: string | number;
  options: { value: string | number; label: string }[]; // Liste d'options
  onChange: (value: string | number) => void;
  isEditable?: boolean; // Par défaut, le champ est éditable
}

const InputFieldList = ({
  label,
  value,
  options,
  onChange,
  isEditable = true, // Par défaut, le champ est éditable
}: InputFieldListProps) => {
  const handleChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    const newValue =
      typeof value === 'number'
        ? Number(event.target.value)
        : event.target.value;
    onChange(newValue);
  };

  return (
    <div className="flex flex-col relative flex-1">
      <label className="mb-3 font-medium text-left text-lightBlack">
        {label}
      </label>
      <div className="relative">
        <select
          value={value}
          onChange={handleChange}
          className="border border-lightGray rounded-xl w-full px-4 h-14 bg-transparent appearance-none" // Gardez le style et ajoutez bg-transparent
          disabled={!isEditable} // Désactive le champ si isEditable est false
        >
          {options.map((option) => (
            <option key={option.value} value={option.value}>
              {option.label}
            </option>
          ))}
        </select>
        {/* Chevron personnalisé */}
        <span className="absolute right-3 top-1/2 transform -translate-y-1/2 pointer-events-none">
          {/* Utilisez une icône SVG ou une image ici */}
          <svg
            width="10"
            height="6"
            viewBox="0 0 10 6"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path d="M1 1L5 5L9 1" stroke="black" strokeWidth="2" />
          </svg>
        </span>
      </div>
    </div>
  );
};

export default InputFieldList;
