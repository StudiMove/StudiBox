import { useSelector } from 'react-redux';
import { RootState } from '../../../../../store';
import { hasAllAdminPermission } from '../../../../config/permissions';
import Logo from '../../Logo/Logo';

import dashboardIcon from '../NavItem/dashboardIcon.svg';
import eventsIcon from '../NavItem/eventsIcon.svg';
import organizationIcon from '../../Navigation/NavItem/organizationIcon.svg';
import userIcon from '../NavItem/userIcon.svg';

import dashboardIconWhite from '../NavItem/dashboardIconWhite.svg';
import eventsIconWhite from '../NavItem/eventsIconWhite.svg';
import organizationIconWhite from '../../Navigation/NavItem/organizationIconWhite.svg';
import userIconWhite from '../NavItem/userIconWhite.svg';

import arrowNavigation from '../SideBar/SideBarIcon/arrow-navigation.svg';
import NavItem from '../NavItem/NavItem';
import SpanMedium from '../../Typography/SpanMedium';

interface SideBarProps {
  isCollapsed: boolean;
  onToggle: () => void;
}

const SideBar = ({ isCollapsed, onToggle }: SideBarProps) => {
  // Récupérer le rôle de l'utilisateur depuis le store Redux
  const role = useSelector((state: RootState) => state.auth.role);

  // Vérifier si l'utilisateur a les permissions nécessaires
  const canAccessOrganisation = hasAllAdminPermission(role);

  return (
    <div
      className={`fixed top-0 left-0 h-full px-6 pt-8 flex flex-col border-r border-lightGray shadow-md bg-white z-50 ${
        isCollapsed ? 'w-32 items-center' : 'w-72'
      } `}
    >
      {/* Logo */}
      <div className="h-8 w-full flex items-center justify-center mb-4">
        {!isCollapsed ? (
          <Logo className="h-full object-contain" variant="textOnly" />
        ) : (
          <Logo className="h-full object-contain" variant="textSB" />
        )}
      </div>

      {/* Icône de flèche pour réduire/agrandir la sidebar */}
      <div
        className={`p-4 cursor-pointer text-white border-b border-lightGray flex items-center justify-${
          isCollapsed ? 'center' : 'between'
        }`}
        onClick={onToggle}
      >
        {!isCollapsed && (
          <SpanMedium className="mr-2">Réduire la navigation</SpanMedium>
        )}
        <img
          src={arrowNavigation}
          alt="arrow-navigation"
          className={`w-6 h-6 transform ${isCollapsed ? 'rotate-180' : ''}`}
        />
      </div>

      {/* Navigation Items */}
      <ul
        className={`mt-10 space-y-4 flex-1 ${
          isCollapsed ? 'justify-center' : ''
        }`}
      >
        <NavItem
          icon={dashboardIcon}
          activeIcon={dashboardIconWhite}
          label={!isCollapsed ? 'Dashboard' : ''}
          to="/dashboard"
          isCollapsed={isCollapsed}
        />
        <NavItem
          icon={eventsIcon}
          activeIcon={eventsIconWhite}
          label={!isCollapsed ? 'Evènements' : ''}
          to="/events"
          isCollapsed={isCollapsed}
        />

        {canAccessOrganisation && ( // Conditionner l'affichage de l'élément Organisation
          <NavItem
            icon={organizationIcon}
            activeIcon={organizationIconWhite}
            label={!isCollapsed ? 'Organisation' : ''}
            to="/organisation"
            isCollapsed={isCollapsed}
          />
        )}
        <NavItem
          icon={userIcon}
          activeIcon={userIconWhite}
          label={!isCollapsed ? 'Utilisateur' : ''}
          to="/*"
          isCollapsed={isCollapsed}
        />
      </ul>
    </div>
  );
};

export default SideBar;
