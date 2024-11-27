import React from 'react';
import checkCircle from '../../Icons/check-circle.svg';

interface SaveButtonProps {
  onClick: (e: React.MouseEvent<HTMLButtonElement>) => void;
}

const SaveButton = ({ onClick }: SaveButtonProps) => {
  return (
    <button
      onClick={onClick}
      className="rounded-xl bg-lightGreen text-white w-16 h-16 flex items-center justify-center backdrop-blur-lg bg-opacity-80"
    >
      <img src={checkCircle} alt="Save" />
    </button>
  );
};

export default SaveButton;
