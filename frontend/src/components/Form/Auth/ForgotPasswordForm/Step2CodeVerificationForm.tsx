// Step2CodeVerificationForm.tsx
import { useState } from 'react';
import InputField from '../../../UI/Inputs/InputField/InputField';
import PrimaryButton from '../../../UI/Button/PrimaryButton';
import { verifyResetCode } from '../../../../services/PasswordService';
import ResendCodeButton from './ResendCodeButton';

interface Step2CodeVerificationFormProps {
  email: string;
  code: number | null;
  setCode: (code: number) => void;
  setStep: (step: number) => void;
  setErrorMessage: (message: string | null) => void;
  setSuccessMessage: (message: string | null) => void;
  setIsLoading: (loading: boolean) => void;
  isLoading: boolean;
}

const Step2CodeVerificationForm = ({
  email,
  code,
  setCode,
  setStep,
  setErrorMessage,
  setSuccessMessage,
  setIsLoading,
  isLoading,
}: Step2CodeVerificationFormProps) => {
  const [isCodeResent, setIsCodeResent] = useState(false);
  const [countdown, setCountdown] = useState(0);

  const handleSubmitCode = async () => {
    setIsLoading(true);
    setErrorMessage(null);
    setSuccessMessage(null);

    try {
      if (code !== null) {
        await verifyResetCode(email, code);
        setSuccessMessage('Code valide !');
        setStep(3);
      }
    } catch (error: unknown) {
      setErrorMessage(
        error instanceof Error ? error.message : 'Une erreur est survenue.',
      );
    } finally {
      setIsLoading(false);
    }
  };

  const handleResendCode = async () => {
    setIsCodeResent(true);
    await handleSubmitCode(); // Réenvoi du code
    setCountdown(30); // Compte à rebours

    const interval = setInterval(() => {
      setCountdown((prev) => {
        if (prev <= 1) {
          clearInterval(interval);
          setIsCodeResent(false);
          return 0;
        }
        return prev - 1;
      });
    }, 1000);
  };

  return (
    <form
      onSubmit={(e) => {
        e.preventDefault();
        handleSubmitCode();
      }}
      className="space-y-4"
    >
      <InputField
        type="number"
        label="Code de réinitialisation"
        value={code ?? ''} // Assure qu'il n'y a pas de valeur null
        onChange={(value) => setCode(parseInt(value as string, 10))}
        size="small"
        isRequired={true}
      />

      <PrimaryButton
        type="submit"
        isLoading={isLoading}
        text={isLoading ? 'Validation en cours...' : 'Valider'}
        className="w-full"
      />
      <ResendCodeButton
        email={email}
        onResend={() => setSuccessMessage('Code renvoyé à votre email.')}
      />
    </form>
  );
};

export default Step2CodeVerificationForm;
