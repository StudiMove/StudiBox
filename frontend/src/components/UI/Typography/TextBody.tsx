// src/components/UI/Typography/TextBody.tsx
import { TextBodyProps } from '../../../types/Typography/TextBodyProps';

const TextBody = function ({ children, className = '' }: TextBodyProps) {
  return (
    <p
      className={`font-sans text-sm leading-normal text-darkGray font-helvetica mb-0 mt-0 ${className}`}
      style={{ fontFamily: 'Helvetica Neue, sans-serif' }}
    >
      {children}
    </p>
  );
};

export default TextBody;

// // H2.tsx
// import { HeadingProps } from '../../../types/Typography/HeadingProps';

// const H2 = ({ children, className = '' }: HeadingProps) => {
//   return (
//     <h2
//       className={`text-2xl md:text-2xl lg:text-2xl font-bold text-darkGray  font-helvetica mb-6 mt-0 ${className}`}
//     >
//       {children}
//     </h2>
//   );
// };
// T
// export default H2;
