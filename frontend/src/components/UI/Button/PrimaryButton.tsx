interface PrimaryButtonProps {
  text: string;
  type?: 'button' | 'submit' | 'reset'; // Type de bouton
  isLoading?: boolean; // État de chargement
  onClick?: () => void; // Gestionnaire de clic
  className?: string; // Classes personnalisées
  disabled?: boolean; // Propriété pour désactiver le bouton
}

const PrimaryButton = function ({
  text,
  type = 'button', // Valeur par défaut
  isLoading = false, // État de chargement par défaut
  onClick,
  className = '', // Classe par défaut
  disabled = false, // Désactivation par défaut
}: PrimaryButtonProps) {
  return (
    <button
      type={type}
      onClick={onClick}
      className={`bg-primary text-white font-bold font-heveltica py-4 px-4 rounded-lg hover:bg-blue-700 transition duration-300 ${className}`}
      disabled={disabled || isLoading} // Désactiver si le bouton est désactivé ou en chargement
    >
      {isLoading ? 'Chargement...' : text}{' '}
      {/* Affiche "Chargement..." si isLoading est true */}
    </button>
  );
};

export default PrimaryButton;
