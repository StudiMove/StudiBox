// import React, { useState } from 'react';
// import { InputFieldProps } from '../../../../types/InputField/InputFieldProps';
// import eyeSlash from './eye-slash.svg'; // Import de l'icône "oeil fermé"

// // Taille des marges et des hauteurs
// const sizeClasses = {
//   small: 'mb-5 h-14 md:mb-6 lg:mb-7',
//   medium: 'mb-4 ',
//   large: 'pb-5 flex-1 h-10 md:mb-6 lg:mb-6',
// };

// const InputField = ({
//   type,
//   label,
//   value,
//   placeholder, // Ajout du placeholder ici
//   size = 'medium',
//   hasIcon = false,
//   onChange,
//   isEditable = true, // Par défaut, le champ est éditable
// }: InputFieldProps) => {
//   const [showPassword, setShowPassword] = useState(false); // État pour afficher ou masquer le mot de passe

//   const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
//     const newValue =
//       type === 'number' ? Number(event.target.value) : event.target.value;
//     onChange(newValue);
//   };

//   const togglePasswordVisibility = () => {
//     setShowPassword((prev) => !prev); // Basculer l'affichage du mot de passe
//   };

//   return (
//     <div className="flex flex-col relative flex-1">
//       <label className="mb-3 font-medium text-left text-lightBlack">
//         {label}
//       </label>
//       <div className="relative">
//         <input
//           type={type === 'password' && showPassword ? 'text' : type} // Affiche du texte si le mot de passe est visible
//           value={value}
//           placeholder={placeholder} // Ajout du placeholder ici
//           onChange={handleChange}
//           className={`border border-lightGray rounded-xl w-full px-4 ${
//             hasIcon && type === 'password' ? 'pr-12' : ''
//           } ${sizeClasses[size]}`} // `pr-12` uniquement si hasIcon est vrai
//           disabled={!isEditable} // Désactive le champ si isEditable est false
//         />
//         {hasIcon && type === 'password' && (
//           <div
//             onClick={togglePasswordVisibility}
//             className="absolute inset-y-0 top-[-16px] right-3 flex items-center cursor-pointer text-gray-600"
//           >
//             <img src={eyeSlash} alt="Toggle password visibility" className="" />
//           </div>
//         )}
//       </div>
//     </div>
//   );
// };

// export default InputField;
import React, { useState } from 'react';
import eyeSlash from './eye-slash.svg';

interface InputFieldProps {
  type: 'text' | 'password' | 'email' | 'number';
  label: string;
  value: string | number;
  placeholder?: string;
  size?: 'small' | 'medium' | 'large';
  hasIcon?: boolean;
  onChange: (value: string | number) => void;
  isEditable?: boolean;
  multiline?: boolean;
  rows?: number;
  isRequired?: boolean;
}

const sizeClasses = {
  small: 'mb-5 h-14 md:mb-6 lg:mb-7',
  medium: 'mb-4 ',
  large: 'pb-5 flex-1 h-10 md:mb-6 lg:mb-6',
};

function InputField({
  type,
  label,
  value,
  placeholder,
  size = 'medium',
  hasIcon = false,
  onChange,
  isEditable = true,
  multiline = false,
  rows = 6,
  isRequired = false,
}: InputFieldProps) {
  const [showPassword, setShowPassword] = useState(false);

  const handleChange = (
    event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>,
  ) => {
    const newValue =
      type === 'number' ? Number(event.target.value) : event.target.value;
    onChange(newValue);
  };

  const togglePasswordVisibility = () => {
    setShowPassword((prev) => !prev);
  };

  return (
    <div className="flex flex-col relative flex-1">
      <label className="mb-3 font-medium text-left text-lightBlack">
        {label}
      </label>
      <div className="relative">
        {multiline ? (
          <textarea
            value={value as string}
            placeholder={placeholder}
            onChange={handleChange}
            rows={rows}
            className={`border border-lightGray hover:border-[#3182ce] rounded-xl w-full px-4 ${sizeClasses[size]}`}
            disabled={!isEditable}
          />
        ) : (
          <input
            type={type === 'password' && showPassword ? 'text' : type}
            value={value}
            placeholder={placeholder}
            onChange={handleChange}
            className={`border border-lightGray rounded-xl w-full px-4 ${
              hasIcon && type === 'password' ? 'pr-12' : ''
            } ${sizeClasses[size]}`}
            disabled={!isEditable}
            required={isRequired}
          />
        )}
        {hasIcon && type === 'password' && (
          <div
            onClick={togglePasswordVisibility}
            className="absolute inset-y-0 top-[-16px] right-3 flex items-center cursor-pointer text-gray-600"
          >
            <img src={eyeSlash} alt="Toggle password visibility" />
          </div>
        )}
      </div>
    </div>
  );
}

export default InputField;
