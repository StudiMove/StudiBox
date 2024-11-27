// // src/components/UI/Container/Container.tsx
// import React from 'react';

// interface ContainerProps {
//   children: React.ReactNode;
//   className?: string;
//   variant?:
//     | 'auth-container'
//     | 'desktop-layout-container'
//     | 'two-input-row'
//     | 'two-input-row-1'; // Ajoutez une prop pour gérer les variantes
// //
// }

// // Utilisation d'une fonction classique
// const Container = function ({
//   children,
//   className = '',
//   variant = 'auth-container', // Valeur par défaut
// }: ContainerProps) {
//   // Définissez les classes en fonction de la variante
//   const variantClasses =
//     variant === 'auth-container'
//       ? 'flex flex-col min-h-screen mx-8 justify-center'
//       : variant === 'two-input-row'
//       ? 'flex justify-between gap-8' // Style pour two-input-row
//       : variant === 'two-input-row-1' // Correction de la syntaxe ici
//       ? 'flex w-1/2 gap-8 pr-4' // Style pour two-input-row-1
//       : variant === 'desktop-layout-container' // Correction de la syntaxe ici
//       ? 'flex flex-col min-h-screen mx-8 '
//       : 'px-16 py-8'; // Style par défaut

//   return <div className={`${variantClasses} ${className}`}>{children}</div>;
// };

// export default Container;
// src/components/UI/Container/Container.tsx

import React from 'react';
import { useIsMobile } from '../../../hooks/useMediaQuery';

interface ContainerProps {
  children: React.ReactNode;
  className?: string;
  variant?:
    | 'auth-container'
    | 'desktop-layout-container'
    | 'two-input-row'
    | 'two-input-row-1'
    | 'two-input-row-75-25'; // Ajoutez la nouvelle variante ici
}

// Utilisation d'une fonction classique
const Container = function ({
  children,
  className = '',
  variant = 'auth-container', // Valeur par défaut
}: ContainerProps) {
  // Définissez les classes en fonction de la variante
  const isMobile = useIsMobile();

  const variantClasses =
    variant === 'auth-container'
      ? 'flex flex-col h-full mx-8 justify-center'
      : variant === 'two-input-row'
      ? isMobile
        ? 'flex flex-col gap-4' // Si isMobile est true, appliquer juste 'flex'
        : 'flex justify-between gap-8'
      : variant === 'two-input-row-1'
      ? isMobile
        ? 'flex flex-col gap-4' // Si isMobile est true, appliquer juste 'flex'
        : 'flex w-1/2 gap-8 pr-4' // Style pour two-input-row-1
      : variant === 'two-input-row-75-25' // Ajout de la nouvelle variante
      ? 'flex gap-8' // Flex container avec écartement
      : variant === 'desktop-layout-container' // Correction de la syntaxe ici
      ? 'flex flex-col h-full mx-8 '
      : 'px-16 py-8'; // Style par défaut

  return <div className={`${variantClasses} ${className}`}>{children}</div>;
};

export default Container;
