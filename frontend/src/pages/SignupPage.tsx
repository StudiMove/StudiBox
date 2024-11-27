import Logo from '../components/UI/Logo/Logo';
import Container from '../components/UI/Container/Container';
import SignupFormComponent from '../components/Form/SignupForm/SignupFormComponent';

const SignupPage = () => {
  return (
    <>
      {/* Wrapper pour le logo */}
      <div className="relative h-20">
        <Logo
          variant="largeIconWithTextDesktop"
          className="hidden md:block absolute top-8 left-8"
        />
      </div>

      {/* Conteneur pour le formulaire, en utilisant flex-grow pour occuper l'espace restant */}
      <div className="flex flex-1 items-center justify-center min-h-[calc(100vh-5rem)]">
        <Container variant="auth-container">
          <div className="mb-12 md:hidden">
            <Logo variant="iconWithText" />
          </div>

          <SignupFormComponent />
        </Container>

        <div className="flex text-blue-500 hidden md:flex"></div>
      </div>
    </>
  );
};

export default SignupPage;
