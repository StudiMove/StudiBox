// Step1EmailForm.tsx

import InputField from '../../../UI/Inputs/InputField/InputField';
import PrimaryButton from '../../../UI/Button/PrimaryButton';
import { requestPasswordReset } from '../../../../services/PasswordService';

interface Step1EmailFormProps {
  email: string;
  setEmail: (email: string) => void;
  setStep: (step: number) => void;
  setErrorMessage: (message: string | null) => void;
  setSuccessMessage: (message: string | null) => void;
  setIsLoading: (loading: boolean) => void;
  isLoading: boolean;
}

function Step1EmailForm({
  email,
  setEmail,
  setStep,
  setErrorMessage,
  setSuccessMessage,
  setIsLoading,
  isLoading,
}: Step1EmailFormProps) {
  const handleSubmitEmail = async () => {
    setIsLoading(true);
    setErrorMessage(null);
    setSuccessMessage(null);

    try {
      await requestPasswordReset(email);
      setSuccessMessage('Un code a été envoyé à votre adresse e-mail.');
      setStep(2);
    } catch (error: unknown) {
      setErrorMessage(
        error instanceof Error ? error.message : 'Une erreur est survenue.',
      );
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <form
      onSubmit={(e) => {
        e.preventDefault();
        handleSubmitEmail();
      }}
      className="space-y-4"
    >
      <InputField
        type="email"
        label="Email"
        value={email}
        onChange={(value) => setEmail(value as string)}
        size="small"
        isRequired={true}
      />
      <PrimaryButton
        type="submit"
        isLoading={isLoading}
        text={isLoading ? 'Envoi en cours...' : 'Réinitialiser'}
        className="w-full"
      />
    </form>
  );
}

export default Step1EmailForm;
