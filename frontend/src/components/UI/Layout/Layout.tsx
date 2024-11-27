// import { Outlet } from 'react-router-dom';

import SideBar from '../Navigation/SideBar/SideBar';
import TopBar from '../Navigation/TopBar/TopBar';
import MobileNavigation from '../Navigation/MobileNavigation/MobileNavigation';
import { useState } from 'react';
import { Outlet } from 'react-router-dom';

import { useIsMobile } from '../../../hooks/useMediaQuery';
const Layout = () => {
  const [isCollapsed, setIsCollapsed] = useState(false);
  const isMobile = useIsMobile();

  const toggleSidebar = () => {
    setIsCollapsed(!isCollapsed);
  };

  if (isMobile) {
    return (
      <div>
        {/* Affichage uniquement sur mobile */}
        <MobileNavigation />
        <div className="mt-16 w-screen overflow-x-hidden">
          <Outlet />
        </div>
      </div>
    );
  }

  return (
    <div className="flex">
      {/* Sidebar fixe avec état contrôlé */}
      <SideBar isCollapsed={isCollapsed} onToggle={toggleSidebar} />

      {/* Conteneur principal qui s'ajuste en fonction de la Sidebar */}
      <div
        className={`flex flex-col transition-all duration-300 w-full ${
          isCollapsed ? 'ml-32' : 'ml-72'
        }`}
      >
        {/* TopBar fixe */}
        <TopBar isCollapsed={isCollapsed} />

        {/* Contenu principal sous la TopBar */}
        <div className="p-8 mt-20">
          <Outlet />
        </div>
      </div>
    </div>
  );
};

export default Layout;
