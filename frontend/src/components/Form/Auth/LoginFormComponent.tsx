import { useState } from 'react';
import { useDispatch } from 'react-redux';
import { LoginForm } from '../../../types/api/auth/AuthForm';
import { LoginResponse } from '../../../types/api/auth/AuthResponse';
import InputField from '../../UI/Inputs/InputField/InputField';
import { loginUser } from '../../../services/AuthService';
import useNavigate from '../../../hooks/useNavigate';
import { login } from '../../../../store/slices/auth.slice';
import { setProfileImage } from '../../../../store/slices/profileImage.slice';
import PrimaryButton from '../../UI/Button/PrimaryButton';
import H1 from '../../UI/Typography/H1';
import H2 from '../../UI/Typography/H2';
import TextBody from '../../UI/Typography/TextBody';
import SpanMedium from '../../UI/Typography/SpanMedium';
import authIllustration from '../../../assets/illustration/authIllustration.svg';

const LoginFormComponent = () => {
  // Gestion des états du formulaire de connexion
  const [formData, setFormData] = useState<LoginForm>({
    email: '',
    password: '',
  });
  const [errorMessage, setErrorMessage] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState<boolean>(false);

  // Gestion de la navigation et du dispatch pour Redux
  const { goTo } = useNavigate();
  const dispatch = useDispatch();

  // Fonction pour mettre à jour les champs du formulaire
  const handleChange = (key: keyof LoginForm) => (value: string | number) => {
    setFormData((prevData) => ({ ...prevData, [key]: value }));
  };

  // Fonction pour gérer la soumission du formulaire
  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    setErrorMessage(null);

    try {
      const data: LoginResponse = await loginUser(formData);
      if (data.isAuthenticated && data.token) {
        // Dispatch l'action de connexion et définit l'image de profil
        dispatch(
          login({
            token: data.token,
            isAuthenticated: data.isAuthenticated,
            role: data.role,
            email: formData.email,
          }),
        );
        dispatch(setProfileImage(data.profileImage.url));

        // Redirection vers le tableau de bord après connexion
        goTo('/dashboard');
      } else {
        setErrorMessage(
          'Échec de la connexion. Veuillez vérifier vos identifiants.',
        );
      }
    } catch (error) {
      setErrorMessage(
        'Échec de la connexion. Veuillez vérifier vos identifiants.',
      );
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="flex flex-col md:flex-row gap-20">
      {/* Section de gauche avec le formulaire de connexion */}
      <div className="w-full md:w-1/2 bg-white flex flex-col">
        <H1>Bienvenue sur StudiBox</H1>
        <H2>Connexion</H2>
        <TextBody className="font-medium mb-6">
          Pas encore de compte ?{' '}
          <span
            className="text-primary cursor-pointer"
            role="button"
            onClick={() => goTo('/')}
          >
            S’inscrire
          </span>
        </TextBody>

        {/* Formulaire de connexion */}
        <form onSubmit={handleSubmit} className="space-y-4 pb-5">
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

          {errorMessage && <p className="text-red-500">{errorMessage}</p>}

          <PrimaryButton
            type="submit"
            isLoading={isLoading}
            text={isLoading ? 'Connexion en cours...' : 'Se Connecter'}
            className="w-full"
          />

          <div onClick={() => goTo('/forgotpassword')}>
            <SpanMedium className="!font-bold flex justify-center cursor-pointer">
              Mot de passe oublié ?
            </SpanMedium>
          </div>
        </form>
      </div>

      {/* Section de droite avec l'illustration */}
      <div className="hidden md:flex w-1/2 items-center justify-center">
        <img src={authIllustration} alt="Illustration d'inscription" />
      </div>
    </div>
  );
};

export default LoginFormComponent;
