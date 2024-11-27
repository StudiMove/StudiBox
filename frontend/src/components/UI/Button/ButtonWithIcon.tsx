import SpanMedium from '../Typography/SpanMedium';

interface ButtonWithIconProps {
  icon: string; // URL de l'icône
  text: string; // Texte à afficher sur le bouton
  onClick?: () => void; // Fonction de rappel pour le clic
  className?: string; // Optionnel : classes CSS supplémentaires si besoin
}

const ButtonWithIcon = ({
  icon,
  text,
  onClick,
  className,
}: ButtonWithIconProps) => {
  const handleClick = () => {
    if (onClick) {
      onClick(); // Appeler la fonction de rappel si elle existe
    }
  };

  return (
    <div
      className={`px-4 py-2 bg-whiteBlack rounded-xl flex items-center cursor-pointer ${className}`}
      onClick={handleClick}
    >
      <img src={icon} alt={`${text} icon`} className="mr-2" />
      <SpanMedium>{text}</SpanMedium>
    </div>
  );
};

export default ButtonWithIcon;
