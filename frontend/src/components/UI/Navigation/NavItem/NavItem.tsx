import { useLocation, Link } from 'react-router-dom';
import SpanMedium from '../../Typography/SpanMedium'; // Assurez-vous que le chemin est correct

interface NavItemProps {
  icon: string;
  label: string;
  to: string;
  isCollapsed: boolean;
  activeIcon?: string; // Icone active envoyée en tant que prop
  onClick?: () => void; // Ajoutez cette prop
}

const NavItem = ({ icon, activeIcon, label, to, onClick }: NavItemProps) => {
  const location = useLocation();
  const isActive = location.pathname === to;

  return (
    <Link to={to} onClick={onClick}>
      <div className="flex items-center  mb-6">
        {/* Appliquer le fond bleu et l'icône blanche si actif */}
        <div className={`p-2.5 rounded-xl ${isActive ? 'bg-primary' : ''}`}>
          <img
            src={isActive ? activeIcon : icon}
            alt={`${label} icon`}
            className={`w-6 h-6 `}
          />
        </div>
        {/* Utiliser SpanMedium pour le label */}
        <SpanMedium
          className={`${isActive ? 'font-bold' : 'text-darkGray'} ml-4`}
        >
          {label}
        </SpanMedium>
      </div>
    </Link>
  );
};

export default NavItem;
// import { useLocation, Link } from 'react-router-dom';
// import SpanMedium from '../../Typography/SpanMedium'; // Assurez-vous que le chemin est correct

// interface NavItemProps {
//   icon: string; // Icone inactive envoyée en tant que prop
//   activeIcon: string; // Icone active envoyée en tant que prop
//   label: string;
//   to: string;
//   isCollapsed: boolean;
//   onClick?: () => void; // Ajoutez cette prop
// }

// const NavItem = ({ icon, activeIcon, label, to, onClick }: NavItemProps) => {
//   const location = useLocation();
//   const isActive = location.pathname === to;

//   return (
//     <Link to={to} onClick={onClick}>
//       <div className="flex items-center mb-6">
//         {/* Appliquer l'icône active ou inactive selon le statut */}
//         <div
//           className={`p-2.5 rounded-xl transition-colors duration-300 ${
//             isActive ? 'bg-primary text-white' : ''
//           }`}
//         >
//           {isActive ? activeIcon : icon}{' '}
//           {/* Afficher l'icône active si actif, sinon l'icône normale */}
//         </div>

//         {/* Utiliser SpanMedium pour le label */}
//         <SpanMedium
//           className={`ml-4 transition-colors duration-300 ${
//             isActive ? 'font-bold text-primary' : 'text-darkGray'
//           }`}
//         >
//           {label}
//         </SpanMedium>
//       </div>
//     </Link>
//   );
// };

// export default NavItem;
