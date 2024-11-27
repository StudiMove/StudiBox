import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import Logo from '../../Logo/Logo';
import dashboardIcon from '../NavItem/dashboardIcon.svg';
import eventsIcon from '../NavItem/eventsIcon.svg';
import organizationIcon from '../../Navigation/NavItem/organizationIcon.svg';
import userIcon from '../NavItem/userIcon.svg';
import NavItem from '../NavItem/NavItem';
import Breadcrumb from '../Breadcrumb/Breadcrumb';
import { hasAllAdminPermission } from '../../../../config/permissions';
import { useSelector } from 'react-redux';
import { RootState } from '../../../../../store';
import hamburgerMenuIcon from '../../../Icons/hamburgerIcon.svg';
import blueCroixIcon from '../../../Icons/blueCroixIcon.svg';
import profilDefaultIcon from '../../../Icons/profilDefaultIcon.svg';
import dashboardIconWhite from '../NavItem/dashboardIconWhite.svg';
import eventsIconWhite from '../NavItem/eventsIconWhite.svg';
import organizationIconWhite from '../../Navigation/NavItem/organizationIconWhite.svg';
import userIconWhite from '../NavItem/userIconWhite.svg';
const MobileNavigation = () => {
  const [isMenuOpen, setIsMenuOpen] = useState(false);

  const profileImage = useSelector(
    (state: RootState) => state.profileImage.profileImage,
  );

  const navigate = useNavigate();

  // Récupérer le rôle pour la gestion des permissions
  const role = useSelector((state: RootState) => state.auth.role);
  const canAccessOrganisation = hasAllAdminPermission(role);

  // Gestion de l'ouverture/fermeture du menu
  const toggleMenu = () => {
    setIsMenuOpen(!isMenuOpen);
  };

  const handleProfileClick = () => {
    navigate('/profil');
  };

  return (
    <div className="relative">
      {/* TopBar mobile */}
      <div className="flex justify-between items-center p-4 bg-white fixed top-0 left-0 right-0 shadow-md z-50">
        {/* Breadcrumb à gauche */}
        <Breadcrumb />

        {/* Conteneur pour les icônes de menu et de profil */}
        <div className="flex items-center space-x-4">
          {/* Icône du menu hamburger */}
          <button onClick={toggleMenu}>
            <div className="flex items-center p-2 bg-lightBlue rounded-lg ">
              <img src={hamburgerMenuIcon} alt="Menu" className="w-6 h-6" />
            </div>
          </button>

          {/* Icône de profil */}
          <img
            src={profileImage || profilDefaultIcon}
            alt="Profile"
            className="w-10 h-10 rounded-full border border-gray-300 cursor-pointer"
            onClick={handleProfileClick}
          />
        </div>
      </div>

      {/* SideMenu mobile */}
      <div
        className={`fixed inset-0 bg-white z-30 transition-transform duration-300 transform ${
          isMenuOpen ? 'translate-x-0' : '-translate-x-full'
        }`}
      >
        <div className="p-4 flex justify-between items-center">
          <div className="h-8 pt-1">
            <Logo className="object-contain" variant="StudiBoxMobile" />
          </div>{' '}
          <button onClick={toggleMenu}>
            <div className="flex items-center p-2 bg-lightBlue rounded-lg ">
              <img src={blueCroixIcon} alt="Menu" className="w-6 h-6" />
            </div>
          </button>
        </div>

        {/* Contenu de la SideBar dans la version mobile */}
        <ul className="mt-10 space-y-4 mx-6">
          <NavItem
            icon={dashboardIcon}
            activeIcon={dashboardIconWhite}
            label="Dashboard"
            to="/dashboard"
            isCollapsed={false}
            onClick={toggleMenu} // Ajout du gestionnaire pour fermer le menu
          />
          <NavItem
            icon={eventsIcon}
            activeIcon={eventsIconWhite}
            label="Evènements"
            to="/events"
            isCollapsed={false}
            onClick={toggleMenu} // Ajout du gestionnaire pour fermer le menu
          />
          {canAccessOrganisation && (
            <NavItem
              icon={organizationIcon}
              activeIcon={organizationIconWhite}
              label="Organisation"
              to="/organisation"
              isCollapsed={false}
              onClick={toggleMenu} // Ajout du gestionnaire pour fermer le menu
            />
          )}
          <NavItem
            icon={userIcon}
            activeIcon={userIconWhite}
            label="Utilisateur"
            to="/*"
            isCollapsed={false}
            onClick={toggleMenu} // Ajout du gestionnaire pour fermer le menu
          />
        </ul>
      </div>
    </div>
  );
};

export default MobileNavigation;
{
  /* <img */
}
//             src={icon}
//             alt={`${label} icon`}
//             className={`w-6 h-6 `}
//             style={{ filter: isActive ? 'invert(1)' : 'none' }}
//           />
//         </div>
