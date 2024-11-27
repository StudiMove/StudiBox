// H3.tsx
import { HeadingProps } from '../../../types/Typography/HeadingProps';

const H3 = ({ children, className = '' }: HeadingProps) => {
  return (
    <h3
      className={`text-xl md:text-xl lg:text-xl font-bold text-darkGray  font-helvetica mb-6 mt-0 ${className}`}
    >
      {children}
    </h3>
  );
};

export default H3;
