// import studiBoxIcon from './logoImages/studiBoxIcon.png'; // Ajustez le chemin si nécessaire
// import iconWithText from '../Logo/LogoImages/iconWithText.svg'; // Ajustez le chemin si nécessaire

// type LogoProps = {
//   variant: 'iconWithText' | 'textOnly' | 'iconOnly' | 'largeIconWithText';
//   className?: string; // Propriété pour les classes CSS supplémentaires
// };

// const Logo = function ({ variant, className = '' }: LogoProps) {
//   const renderLogo = () => {
//     switch (variant) {
//       case 'iconWithText':
//         return <img src={iconWithText} alt="Logo avec icon" />;
//       case 'textOnly':
//         return <img src={iconWithText} alt="Logo StudiBox text" />;
//       case 'iconOnly':
//         return <img src={iconWithText} alt="Icon StudiBox" />;
//       case 'largeIconWithText':
//         return <img src={iconWithText} alt="Large logo" />;
//       default:
//         return null;
//     }
//   };

//   return <div className={className}>{renderLogo()}</div>;
// };

// export default Logo;
import iconWithText from '../Logo/LogoImages/iconWithText.svg'; // Ajustez le chemin si nécessaire
import textSB from '../Logo/LogoImages/textSB.svg';
import textOnly from '../Logo/LogoImages/textOnly.svg';
import StudiBoxMobile from '../Logo/LogoImages/StudiBoxMobile.svg';
type LogoProps = {
  variant:
    | 'iconWithText'
    | 'textOnly'
    | 'StudiBoxMobile'
    | 'largeIconWithText'
    | 'largeIconWithTextDesktop'
    | 'textSB';
  className?: string; // Propriété pour les classes CSS supplémentaires
};

const Logo = function ({ variant, className = '' }: LogoProps) {
  const renderLogo = () => {
    const imgClass = 'object-contain h-12 max-h-12'; // Classe Tailwind pour l'image

    switch (variant) {
      case 'iconWithText':
        return <img src={iconWithText} alt="Logo avec icon" />;
      case 'textOnly':
        return <img src={textOnly} alt="Logo StudiBox text" />;
      case 'StudiBoxMobile':
        return <img src={StudiBoxMobile} alt="Icon StudiBox" />;
      case 'textSB':
        return <img src={textSB} alt="Icon StudiBox" />;
      case 'largeIconWithText':
        return <img src={iconWithText} alt="Large logo" />;
      case 'largeIconWithTextDesktop':
        return <img src={iconWithText} alt="Large logo" className={imgClass} />;
      default:
        return null;
    }
  };

  return <div className={className}>{renderLogo()}</div>;
};

export default Logo;
