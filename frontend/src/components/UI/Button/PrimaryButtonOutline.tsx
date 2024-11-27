interface PrimaryButtonOutlineProps {
  label: string; // Texte du bouton
  type?: 'button' | 'submit' | 'reset'; // Type de bouton
  onClick: () => void; // Fonction appelée lors du clic sur le bouton
  className?: string; // Optionnel : classes CSS supplémentaires si besoin
  disabled?: boolean; // Optionnel : état désactivé
  isLoading?: boolean; // État de chargement
}

const PrimaryButtonOutline = ({
  label,
  onClick,
  type = 'button',
  className = '',
  disabled = false,
  isLoading = false, // Ajout de l'état de chargement
}: PrimaryButtonOutlineProps) => {
  return (
    <button
      type={type}
      onClick={onClick}
      disabled={disabled || isLoading} // Désactiver le bouton si désactivé ou en chargement
      className={`w-full px-6 py-2 font-helvetica font-bold border-2 transition-colors duration-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary focus:ring-opacity-50
        ${
          disabled
            ? 'border-gray-400 text-gray-400 opacity-50 cursor-not-allowed' // Styles lorsque le bouton est désactivé
            : 'border-primary text-primary hover:bg-primary hover:text-white'
        } // Styles lorsque le bouton est actif
        ${className}`}
    >
      {isLoading ? 'Chargement...' : label}{' '}
      {/* Affiche "Chargement..." si isLoading est true */}
    </button>
  );
};

export default PrimaryButtonOutline;
