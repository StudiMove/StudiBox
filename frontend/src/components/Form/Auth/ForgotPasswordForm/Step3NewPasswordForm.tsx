// Step3NewPasswordForm.tsx
import InputField from '../../../UI/Inputs/InputField/InputField';
import PrimaryButton from '../../../UI/Button/PrimaryButton';
import { updatePassword } from '../../../../services/PasswordService';

interface Step3NewPasswordFormProps {
  email: string;
  newPassword: string;
  confirmPassword: string;
  setNewPassword: (password: string) => void;
  setConfirmPassword: (password: string) => void;
  setErrorMessage: (message: string | null) => void;
  setSuccessMessage: (message: string | null) => void;
  setIsLoading: (loading: boolean) => void;
  isLoading: boolean; // Ajout de isLoading
}

const Step3NewPasswordForm: React.FC<Step3NewPasswordFormProps> = ({
  email,
  newPassword,
  confirmPassword,
  setNewPassword,
  setConfirmPassword,
  setErrorMessage,
  setSuccessMessage,
  setIsLoading,
  isLoading, // Ajout de isLoading
}) => {
  const handleSubmitNewPassword = async () => {
    setIsLoading(true);
    setErrorMessage(null);
    setSuccessMessage(null);

    if (newPassword === confirmPassword) {
      try {
        await updatePassword(email, newPassword);
        setSuccessMessage('Votre mot de passe a été mis à jour avec succès.');
      } catch (error: unknown) {
        setErrorMessage(
          error instanceof Error ? error.message : 'Une erreur est survenue.',
        );
      }
    } else {
      setErrorMessage('Les mots de passe ne correspondent pas.');
    }

    setIsLoading(false);
  };

  return (
    <form
      onSubmit={(e) => {
        e.preventDefault();
        handleSubmitNewPassword();
      }}
      className="space-y-4"
    >
      <InputField
        type="password"
        label="Nouveau Mot de Passe"
        value={newPassword}
        onChange={(value) => setNewPassword(value as string)}
        size="small"
        hasIcon
        isRequired={true}
      />
      <InputField
        type="password"
        label="Confirmer le Mot de Passe"
        value={confirmPassword}
        onChange={(value) => setConfirmPassword(value as string)}
        size="small"
        hasIcon
        isRequired={true}
      />
      <PrimaryButton
        type="submit"
        isLoading={isLoading} // Utilisation de isLoading ici
        text={isLoading ? 'Mise à jour...' : 'Mettre à jour'}
        className="w-full"
      />
    </form>
  );
};

export default Step3NewPasswordForm;
