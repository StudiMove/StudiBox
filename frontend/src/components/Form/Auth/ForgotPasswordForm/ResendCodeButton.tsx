// ResendCodeButton.tsx
import { useState, useEffect } from 'react';
import PrimaryButtonOutline from '../../../UI/Button/PrimaryButtonOutline';
import { requestPasswordReset } from '../../../../services/PasswordService';

interface ResendCodeButtonProps {
  email: string;
  onResend: () => void;
}

const ResendCodeButton: React.FC<ResendCodeButtonProps> = ({
  email,
  onResend,
}) => {
  const [isCodeResent, setIsCodeResent] = useState(false);
  const [countdown, setCountdown] = useState<number>(0);

  const handleResendCode = async () => {
    try {
      setIsCodeResent(true);
      await requestPasswordReset(email);
      onResend(); // Callback après réenvoi du code
      setCountdown(30); // Lancement du compte à rebours de 30 secondes
    } catch (error) {
      console.error('Erreur lors de l’envoi du code:', error);
    }
  };

  useEffect(() => {
    if (countdown > 0) {
      const timer = setInterval(() => {
        setCountdown((prev) => prev - 1);
      }, 1000);

      return () => clearInterval(timer);
    } else {
      setIsCodeResent(false); // Réinitialisation une fois le compte à rebours terminé
    }
  }, [countdown]);

  return (
    <PrimaryButtonOutline
      type="button"
      label={
        isCodeResent
          ? `Envoyé de nouveau dans ${countdown}s`
          : 'Renvoyer le code'
      }
      onClick={handleResendCode}
      disabled={isCodeResent} // Désactiver le bouton pendant le compte à rebours
      className="w-full"
    />
  );
};

export default ResendCodeButton;
