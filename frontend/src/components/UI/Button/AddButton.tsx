import greenPlusCircle from '../../Icons/green-plus-circle.svg';
import { useNavigate } from 'react-router-dom';

interface AddButtonProps {
  navigateTo: string; // URL vers laquelle rediriger
}

const AddButton = ({ navigateTo }: AddButtonProps) => {
  const navigate = useNavigate();

  const handleClick = () => {
    navigate(navigateTo); // Rediriger vers l'URL passée en props
  };

  return (
    <button
      onClick={handleClick}
      className="rounded-xl bg-lightGreen text-white w-16 h-16 flex items-center justify-center backdrop-blur-lg bg-opacity-80"
    >
      <img src={greenPlusCircle} alt="Save" />
    </button>
  );
};

export default AddButton;
