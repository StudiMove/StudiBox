// import DataTable from '../../Table/DataTable'; // Chemin d'importation correct
// import { User } from './../../../../types/Table/DataTableTypes'; // Assurez-vous que le chemin est correct

// function UserList() {
//   const users: User[] = [
//     {
//       id: 1,
//       profileImage: 'path/to/image1.jpg',
//       pseudo: 'User1',
//       firstName: 'John',
//       email: 'john@example.com',
//       type: 'Admin',
//     },
//     // Ajoutez d'autres utilisateurs
//   ];

//   // const handleRowClick = (id: number) => {
//   //   console.log('Clicked User ID:', id);
//   //   // Redirection vers la page de profil utilisateur, par exemple :
//   //   // navigate(`/user/${id}`);
//   // };

//   return (
//     <div>
//       <h2>Liste des Utilisateurs</h2>
//       <DataTable
//         data={users}
//         headers={['Image de Profil', 'Pseudo', 'Nom', 'Email', 'Type']}
//         onRowClick={handleRowClick}
//       />
//     </div>
//   );
// }

// export default UserList; // N'oubliez pas d'exporter le composant
import { User } from '../../../../types/Table/DataTableTypes'; // Assurez-vous que le chemin est correct
import DataTable from '../../Table/DataTable';

interface UserListProps {
  data: User[]; // Les données des entreprises passées via props
  headers: string[]; // Les en-têtes du tableau
  onRowClick: (id: number, name: string) => void; // Fonction de callback lors du clic sur une ligne
}

function UserList({ data, headers, onRowClick }: UserListProps) {
  const handleClick = (user: User) => {
    onRowClick(user.id, user.name); // Appeler la fonction onRowClick avec l'ID et le nom de l'entreprise
  };

  return (
    <DataTable
      data={data}
      headers={headers}
      onRowClick={handleClick} // Associer le clic à une ligne du tableau
    />
  );
}

export default UserList;
