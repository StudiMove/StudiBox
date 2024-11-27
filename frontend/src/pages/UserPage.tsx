import { useNavigate } from 'react-router-dom'; // Importer useNavigate
import TabsData from '../components/UI/Tab/TabsData';
import CompanyList from '../components/UI/List/CompanyList/CompanyList';
import { getActiveOrganisations } from '../services/OrganisationService';

const UserPage = () => {
  const headers = ['', 'Pseudo', 'Nom', 'Prenom', 'Mail', 'Type'];
  const navigate = useNavigate();

  // Fonction de mappage des données
  const mapOrganisationData = (data: any[]) => {
    return data.map((company: any) => ({
      id: company.user_id, // Associer `user_id` à `id`
      profileImage: company.profile_image,
      companyName: company.company_name,
      type: company.role_name,
      status: company.status,
      member: '0',
      event: '0',
    }));
  };

  // Fonction pour gérer le clic sur une ligne du tableau
  const handleRowClick = (id: number, companyName: string) => {
    navigate(`/organisation/${companyName}`, { state: { targetId: id } });
  };

  const tabs = [
    {
      label: '',
      dataLoader: getActiveOrganisations,
    },
  ];

  return (
    <div className="px-8">
      <TabsData
        tabs={tabs}
        headers={headers}
        DataTableComponent={(props) => (
          <CompanyList
            data={props.data}
            headers={props.headers}
            onRowClick={handleRowClick} // Passe la fonction onRowClick ici
          />
        )}
        mapData={mapOrganisationData}
        filterKey="companyName"
        // Ajout de la fonction de mappage des données
      />
    </div>
  );
};

export default UserPage;
