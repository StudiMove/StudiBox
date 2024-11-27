import { useSelector } from 'react-redux';
import Breadcrumb from '../Breadcrumb/Breadcrumb';
import { RootState } from '../../../../../store';
import { useState } from 'react';
import ProfilPopup from '../ProfilPopup/ProfilPopup';
interface TopBarProps {
  isCollapsed: boolean;
}

const TopBar = ({ isCollapsed }: TopBarProps) => {
  // Récupérer l'image de profil et les autres infos depuis le store Redux
  const profileImage = useSelector(
    (state: RootState) => state.profileImage.profileImage,
  );

  const defaultImageUrl = ''; // Image par défaut
  const [isProfilePopupVisible, setProfilePopupVisible] = useState(false);

  const handleProfileClick = () => {
    setProfilePopupVisible((prev) => !prev); // Toggle pour afficher/masquer le ProfilPopup
  };

  // Fonction de gestion d'erreur de chargement de l'image
  const handleImageError = (
    event: React.SyntheticEvent<HTMLImageElement, Event>,
  ) => {
    event.currentTarget.src = defaultImageUrl; // Remplace l'image par défaut en cas d'erreur
  };

  return (
    <div
      className={`fixed top-0 left-0 right-0 z-10 bg-white px-8 py-4 border-b border-lightGray flex justify-between items-center z-30 ${
        isCollapsed ? 'ml-32' : 'ml-72'
      }`}
    >
      {/* Breadcrumb à gauche */}
      <div className="flex-1">
        <Breadcrumb />
      </div>

      {/* Image de profil à droite */}
      <div className="flex items-center">
        <img
          src={profileImage || defaultImageUrl}
          alt="Profile"
          className="w-10 h-10 rounded-full border border-gray-300 cursor-pointer"
          onClick={handleProfileClick} // Affiche le ProfilPopup au clic
          onError={handleImageError}
        />
      </div>

      {isProfilePopupVisible && (
        <div className=" absolute right-8 top-24 z-40">
          <ProfilPopup onClose={() => setProfilePopupVisible(false)} />
        </div>
      )}
    </div>
  );
};

export default TopBar;
