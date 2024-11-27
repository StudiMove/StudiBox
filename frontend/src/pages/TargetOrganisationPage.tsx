import { useLocation } from 'react-router-dom';
import EventPage from './Event/EventPage';
import Tabs from '../components/UI/Tab/Tabs';
import ProfilFormComponent from '../components/Form/ProfilForm/ProfileFormComponent';
const TargetOrganisationPage = () => {
  const location = useLocation();
  const { targetId } = location.state || {}; // Récupérer targetId depuis le state passé via navigate

  // Définir les onglets et leurs composants respectifs
  const tabs = [
    {
      label: 'Details',
      component: <ProfilFormComponent targetId={targetId} />,
    },
    {
      label: 'Events',
      component: <EventPage targetId={targetId} />, // targetId est optionnel
    },
  ];

  return (
    <>
      {/* Utilisation du composant Tabs pour gérer les onglets */}
      <div className="px-8">
        <Tabs tabs={tabs} />
      </div>
    </>
  );
};

export default TargetOrganisationPage;
