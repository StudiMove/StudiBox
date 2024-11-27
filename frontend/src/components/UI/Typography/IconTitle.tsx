import H2 from './H2';

// Définition des props pour IconTitle
interface IconTitleProps {
  title: string;
  image: string;
  className?: string; // Ajout de la prop className
}

const IconTitle = ({ title, image, className = '' }: IconTitleProps) => {
  return (
    <div className={`flex items-center gap-x-3 ${className}`}>
      {/* Image */}
      <img src={image} alt={title} className="w-8 h-8" />

      {/* Title avec la classe personnalisée */}
      <H2 className="!m-0 !text-black font-medium">{title}</H2>
    </div>
  );
};

export default IconTitle;
