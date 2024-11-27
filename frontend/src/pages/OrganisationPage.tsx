import { useNavigate } from 'react-router-dom';
import TabsData from '../components/UI/Tab/TabsData';
import CompanyList from '../components/UI/List/CompanyList/CompanyList';
import {
  getAllOrganisations,
  getActiveOrganisations,
  getInactiveOrganisations,
  getPendingOrganisations,
  getSuspendedOrganisations,
} from '../services/OrganisationService';
import { Company } from '../types/Table/DataTableTypes';

function OrganisationPage() {
  const headers = ['', 'Nom', 'Type', 'Statut', 'Évènements', 'Membres'];
  const navigate = useNavigate();

  // Fonction de mappage des données de l'organisation
  const mapOrganisationData = (data: any[]): Company[] => {
    return data.map((organisation) => ({
      user_id: organisation.user_id,
      profile_image: organisation.profile_image,
      name: organisation.name,
      role_name: organisation.role_name,
      status: organisation.status,
      is_validated: organisation.is_validated,
      is_activated: organisation.is_activated,
      is_pending: organisation.is_pending,
    }));
  };

  // Gérer le clic sur une ligne
  const handleRowClick = (id: number, companyName: string) => {
    navigate(`/organisation/${companyName}`, { state: { targetId: id } });
  };

  const tabs = [
    { label: 'Tout', dataLoader: getAllOrganisations },
    { label: 'Validés', dataLoader: getActiveOrganisations },
    { label: 'En Attente', dataLoader: getPendingOrganisations },
    { label: 'Suspendu', dataLoader: getSuspendedOrganisations },
    { label: 'Invalidés', dataLoader: getInactiveOrganisations },
  ];
  console.log(mapOrganisationData);
  return (
    <div className="px-8">
      <TabsData
        tabs={tabs}
        headers={headers}
        DataTableComponent={(props) => (
          <CompanyList
            data={props.data}
            headers={props.headers}
            onRowClick={handleRowClick}
          />
        )}
        mapData={mapOrganisationData}
        filterKey="name"
      />
    </div>
  );
}

export default OrganisationPage;
