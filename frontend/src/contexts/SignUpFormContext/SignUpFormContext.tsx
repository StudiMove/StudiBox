// import React, { createContext, useState, ReactNode } from 'react';
// import { SignUpFormContextType } from '../../types/SignUpForm/SignUpFormContextType';

// export const SignUpFormContext = createContext<
//   SignUpFormContextType | undefined
// >(undefined);

// export const SignUpFormProvider: React.FC<{ children: ReactNode }> = ({
//   children,
// }) => {
//   const [currentStep, setCurrentStep] = useState(0);
//   const [formData, setFormData] = useState({
//     email: '',
//     password: '',
//     address: '',
//     postalCode: '',
//     country: '',
//   });

//   return (
//     <SignUpFormContext.Provider
//       value={{ currentStep, setCurrentStep, formData, setFormData }}
//     >
//       {children}
//     </SignUpFormContext.Provider>
//   );
// };
