interface SecondaryButtonOutlineProps {
  label: string; // Texte du bouton
  type?: 'button' | 'submit' | 'reset'; // Type de bouton
  onClick: () => void; // Fonction appelée lors du clic sur le bouton
  className?: string; // Optionnel : classes CSS supplémentaires si besoin
  disabled?: boolean; // Optionnel : état désactivé
  isLoading?: boolean; // État de chargement
}

const SecondaryButtonOutline = ({
  label,
  onClick,
  type = 'button',
  className = '',
  disabled = false,
  isLoading = false, // Ajout de l'état de chargement
}: SecondaryButtonOutlineProps) => {
  return (
    <button
      type={type}
      onClick={onClick}
      disabled={disabled || isLoading} // Désactiver le bouton si désactivé ou en chargement
      className={`w-full px-6 py-2 font-helvetica font-bold border-2 border-red text-red transition-colors duration-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary focus:ring-opacity-50
        hover:bg-red hover:text-white ${className}`}
    >
      {isLoading ? 'Chargement...' : label}{' '}
      {/* Affiche "Chargement..." si isLoading est true */}
    </button>
  );
};

export default SecondaryButtonOutline;
