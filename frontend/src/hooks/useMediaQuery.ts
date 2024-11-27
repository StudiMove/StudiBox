import { useMediaQuery } from 'react-responsive';

const maxWidth = 768;
export const useIsMobile = () => {
  return useMediaQuery({ maxWidth: maxWidth });
};
