import { useState } from 'react';
import { SignupForm } from '../../../types/api/auth/AuthForm';
import InputField from '../../UI/Inputs/InputField/InputField';
import InputFieldList from '../../UI/Inputs/InputFieldList/InputFieldList';
import { signupUser } from '../../../services/AuthService';
import useNavigate from '../../../hooks/useNavigate';
import signupIllustration from '../../../assets/illustration/signupIllustration.svg';
import Container from '../../UI/Container/Container';
import PrimaryButton from '../../UI/Button/PrimaryButton';
import H1 from '../../UI/Typography/H1';
import H2 from '../../UI/Typography/H2';
import TextBody from '../../UI/Typography/TextBody';

const SignupFormComponent = () => {
  const [step, setStep] = useState(1);
  const [formData, setFormData] = useState<SignupForm>({
    email: '',
    password: '',
    organisationName: '',
    address: '',
    postalCode: '',
    city: '',
    country: '',
    phone: '',
    description: '',
    organisationType: 'business',
  });

  const [errorMessage, setErrorMessage] = useState<string | null>(null);
  const [passwordError, setPasswordError] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const { goTo } = useNavigate();

  const passwordRegex =
    /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d!@#$%^&*()_+={}\[\]:;"'<>,.?/~`\\-]{8,}$/;

  const handleChange = (key: keyof SignupForm) => (value: string | number) => {
    setFormData((prevData) => ({ ...prevData, [key]: value }));

    if (key === 'password') {
      if (!passwordRegex.test(value as string)) {
        setPasswordError(
          'Le mot de passe doit contenir au moins 8 caractères, dont une majuscule et un chiffre.',
        );
      } else {
        setPasswordError(null);
      }
    }
  };

  const handleNextStep = (e: React.FormEvent) => {
    e.preventDefault();

    if (!formData.email || !formData.password) {
      setErrorMessage('Veuillez renseigner votre email et mot de passe.');
    } else if (passwordError) {
      setErrorMessage(passwordError);
    } else {
      setErrorMessage(null);
      setStep(2);
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    setErrorMessage(null);

    try {
      const data = await signupUser(formData);
      if (data.success) {
        goTo('/login');
      } else {
        setErrorMessage(
          "Échec de l'inscription. Veuillez vérifier vos informations.",
        );
      }
    } catch (error) {
      setErrorMessage("Échec de l'inscription. Veuillez réessayer.");
    } finally {
      setIsLoading(false);
    }
  };

  const redirectToLogin = () => {
    goTo('/login');
  };

  return (
    <div className="flex flex-col md:flex-row gap-20 ">
      <div className="w-full md:w-1/2 bg-white flex flex-col overflow-y-auto max-h-[60vh] p-4">
        {step === 1 ? (
          <div>
            <H1>Bienvenue sur StudiBox</H1>
            <H2>Inscription</H2>
            <TextBody className="font-medium mb-6">
              Vous avez déjà un compte ?{' '}
              <span
                className="text-primary cursor-pointer z-10"
                role="button"
                onClick={redirectToLogin}
              >
                Connectez-vous
              </span>{' '}
            </TextBody>

            <form onSubmit={handleNextStep} className="space-y-4">
              <InputField
                type="email"
                label="Email"
                value={formData.email}
                onChange={handleChange('email')}
                size="small"
                isRequired={true}
              />
              <InputField
                type="password"
                label="Mot de Passe"
                value={formData.password}
                onChange={handleChange('password')}
                hasIcon={true}
                size="small"
                isRequired={true}
              />

              <TextBody className="font-medium">
                Minimum 8 caractères, contenant au moins 1 lettre minuscule,
              </TextBody>
              <div className="min-h-[24px]">
                {passwordError && (
                  <p className="text-red-500">{passwordError}</p>
                )}
              </div>

              <PrimaryButton
                type="submit"
                isLoading={isLoading}
                text="Continuer"
                className="w-full"
              />
            </form>
          </div>
        ) : (
          <div className="my-16">
            <H2>Compléter vos informations</H2>
            <form onSubmit={handleSubmit} className="space-y-4">
              <InputFieldList
                label="Type d'organisation"
                value={formData.organisationType}
                options={[
                  { label: 'Entreprise', value: 'business' },
                  { label: 'Association', value: 'association' },
                  { label: 'École', value: 'school' },
                ]}
                onChange={handleChange('organisationType')}
              />
              <InputField
                type="text"
                label="Nom de votre organisation"
                value={formData.organisationName}
                onChange={handleChange('organisationName')}
                size="small"
                isRequired={true}
              />
              <Container variant="two-input-row-75-25">
                <div className="w-3/4">
                  <InputField
                    type="text"
                    label="Adresse"
                    value={formData.address}
                    onChange={handleChange('address')}
                    size="small"
                    isRequired={true}
                  />
                </div>
                <div className="w-1/4">
                  <InputField
                    type="text"
                    label="Code Postal"
                    value={formData.postalCode}
                    onChange={handleChange('postalCode')}
                    size="small"
                    isRequired={true}
                  />
                </div>
              </Container>

              <Container variant="two-input-row-75-25">
                <div className="w-3/4">
                  <InputField
                    type="text"
                    label="Ville"
                    value={formData.city}
                    onChange={handleChange('city')}
                    size="small"
                    isRequired={true}
                  />
                </div>
                <div className="w-1/4">
                  <InputField
                    type="text"
                    label="Pays"
                    value={formData.country}
                    onChange={handleChange('country')}
                    size="small"
                    isRequired={true}
                  />
                </div>
              </Container>

              <InputField
                type="text"
                label="Téléphone"
                value={formData.phone}
                onChange={handleChange('phone')}
                size="small"
                isRequired={true}
              />
              <InputField
                type="text"
                label="Description de la demande et des évènements proposés"
                value={formData.description}
                onChange={handleChange('description')}
                size="small"
                multiline={true}
                rows={6}
              />

              <PrimaryButton
                type="submit"
                isLoading={isLoading}
                text="Inscription"
                className="w-full"
              />
            </form>
          </div>
        )}
      </div>

      <div className="hidden md:flex w-1/2 items-center justify-center">
        <img src={signupIllustration} alt="Illustration d'inscription" />
      </div>
    </div>
  );
};

export default SignupFormComponent;
