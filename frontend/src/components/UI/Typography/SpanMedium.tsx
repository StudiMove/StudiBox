// src/components/UI/Typography/SpanMedium.tsx
import { TextBodyProps } from '../../../types/Typography/TextBodyProps';

const SpanMedium = function ({ children, className = '' }: TextBodyProps) {
  return (
    <span
      className={`font-sans text-base  text-darkGray font-medium font-helvetica mb-0 mt-0 ${className}`}
      style={{ fontFamily: 'Helvetica Neue, sans-serif' }}
    >
      {children}
    </span>
  );
};

export default SpanMedium;
