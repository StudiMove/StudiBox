// src/hooks/useNavigate.ts
import { useNavigate as useReactRouterNavigate } from 'react-router-dom';

const useNavigate = () => {
  const navigate = useReactRouterNavigate();

  const goTo = (path: string) => {
    navigate(path);
  };

  return { goTo };
};

export default useNavigate;
