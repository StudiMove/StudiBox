// H1.tsx
import { HeadingProps } from '../../../types/Typography/HeadingProps';

const H1 = ({ children, className = '' }: HeadingProps) => {
  return (
    <h1
      className={`text-4xl md:text-4xl lg:text-4xl font-helvetica font-bold text-darkGray mb-8 mt-0 ${className}`}
    >
      {children}
    </h1>
  );
};

export default H1;
