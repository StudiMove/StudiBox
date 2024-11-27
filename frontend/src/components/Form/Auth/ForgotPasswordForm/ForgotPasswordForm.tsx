import { useState } from 'react';
import H1 from '../../../UI/Typography/H1';
import H2 from '../../../UI/Typography/H2';
import TextBody from '../../../UI/Typography/TextBody';
import passwordIllustration from '../../../../assets/illustration/passwordIllustration.svg';
import Step1EmailForm from './Step1EmailForm';
import Step2CodeVerificationForm from './Step2CodeVerificationForm';
import Step3NewPasswordForm from './Step3NewPasswordForm';
import useNavigate from '../../../../hooks/useNavigate';

const ForgotPasswordForm = () => {
  const [step, setStep] = useState<number>(1);
  const [email, setEmail] = useState<string>('');
  const [code, setCode] = useState<number | null>(null);
  const [newPassword, setNewPassword] = useState<string>('');
  const [confirmPassword, setConfirmPassword] = useState<string>('');
  const [errorMessage, setErrorMessage] = useState<string | null>(null);
  const [successMessage, setSuccessMessage] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const { goTo } = useNavigate();

  const handleRedirectToLogin = () => goTo('/login');

  const renderStepForm = () => {
    switch (step) {
      case 1:
        return (
          <Step1EmailForm
            email={email}
            setEmail={setEmail}
            setStep={setStep}
            setErrorMessage={setErrorMessage}
            setSuccessMessage={setSuccessMessage}
            setIsLoading={setIsLoading}
            isLoading={isLoading}
          />
        );
      case 2:
        return (
          <Step2CodeVerificationForm
            email={email}
            code={code}
            setCode={(value) => setCode(value)}
            setStep={setStep}
            setErrorMessage={setErrorMessage}
            setSuccessMessage={setSuccessMessage}
            setIsLoading={setIsLoading}
            isLoading={isLoading}
          />
        );
      case 3:
        return (
          <Step3NewPasswordForm
            email={email}
            newPassword={newPassword}
            confirmPassword={confirmPassword}
            setNewPassword={setNewPassword}
            setConfirmPassword={setConfirmPassword}
            setErrorMessage={setErrorMessage}
            setSuccessMessage={setSuccessMessage}
            setIsLoading={setIsLoading}
            isLoading={isLoading}
          />
        );
      default:
        return null;
    }
  };

  return (
    <div className="flex flex-col md:flex-row gap-20">
      <div className="w-full md:w-1/2 bg-white flex flex-col">
        <H1 className="w-96">Mot de passe oublié ?</H1>
        <H2>Récupérer votre compte</H2>
        <TextBody className="font-medium mb-6">
          Vous connaissez votre mot de passe ?{' '}
          <span
            className="text-primary cursor-pointer"
            role="button"
            onClick={handleRedirectToLogin}
          >
            Connectez-vous
          </span>
        </TextBody>
        {renderStepForm()}
      </div>
      <div className="hidden md:flex w-1/2 items-center justify-center">
        <img
          src={passwordIllustration}
          alt="Illustration de réinitialisation"
        />
      </div>
    </div>
  );
};

export default ForgotPasswordForm;
