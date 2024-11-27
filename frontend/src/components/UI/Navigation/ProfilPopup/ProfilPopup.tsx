// ProfilPopup.tsx
import disconnectIcon from '../../../Icons/disconnect-icon.svg';
import updateIcon from '../../../Icons/update-icon.svg';
import ButtonWithIcon from '../../Button/ButtonWithIcon';
import { useSelector } from 'react-redux';
import { RootState } from '../../../../../store';
import { useNavigate } from 'react-router-dom'; // Importer useNavigate
import SpanMedium from '../../Typography/SpanMedium';
import defaultImageUrl from '../../../Icons/profilDefaultIcon.svg';
import { logout } from '../../../../../store/slices/auth.slice';
import { useDispatch } from 'react-redux';
interface ProfilPopupProps {
  onClose: () => void; // Fonction de rappel pour fermer le popup
}

const ProfilPopup = ({ onClose }: ProfilPopupProps) => {
  const profileImage = useSelector(
    (state: RootState) => state.profileImage.profileImage,
  );
  const email = useSelector((state: RootState) => state.auth.email);
  const navigate = useNavigate();
  const dispatch = useDispatch(); // Initialisez useDispatch pour utiliser le dispatch

  return (
    <div className="p-6 rounded-xl inline-flex flex-col items-center justify-center bg-white shadow-md min-w-[300px]">
      <div className="flex flex-col items-center justify-center gap-2 mb-5">
        <img
          src={profileImage || defaultImageUrl}
          alt="Profile"
          className="w-10 h-10 rounded-full border border-gray-300 cursor-pointer"
        />
        <SpanMedium className="text-center">{email}</SpanMedium>
      </div>
      <div className="h-px bg-black w-full mb-5"></div>
      <div className="flex bg-white rounded-lg w-full">
        <div className="flex flex-col gap-3 w-full">
          <ButtonWithIcon
            icon={updateIcon}
            text="Modifier le profil"
            onClick={() => {
              navigate('/profil');
              onClose(); // Fermer le popup
            }}
          />
          <ButtonWithIcon
            icon={disconnectIcon}
            text="Déconnexion"
            onClick={() => {
              dispatch(logout()); // Appel de la fonction logout
              navigate('/'); // Redirection vers la page d'accueil
              onClose(); // Fermer le popup
            }}
          />
        </div>
      </div>
    </div>
  );
};

export default ProfilPopup;
