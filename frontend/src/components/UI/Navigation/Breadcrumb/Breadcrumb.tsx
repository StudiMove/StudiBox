// import { Link, useLocation, useParams } from 'react-router-dom';

// interface BreadcrumbItem {
//   name: string;
//   path: string;
// }

// const Breadcrumb = () => {
//   const location = useLocation();
//   const params = useParams();

//   // Diviser l'URL en segments
//   const pathnames = location.pathname.split('/').filter((x) => x);

//   // Générer les éléments du breadcrumb en fonction des segments
//   const breadcrumbItems: BreadcrumbItem[] = pathnames.map((segment, index) => {
//     const path = `/${pathnames.slice(0, index + 1).join('/')}`;
//     let name = segment;

//     // Vérifie les segments pour les routes spécifiques
//     if (segment === 'organisation' && params.companyName) {
//       name = params.companyName; // Remplace par le nom de l'organisation
//     } else if (
//       segment === 'events' &&
//       pathnames[index + 1] === 'Créer un nouvel événement'
//     ) {
//       name = 'Événements';
//     } else if (segment === 'events') {
//       name = 'Événements';
//     }

//     return { name: name.charAt(0).toUpperCase() + name.slice(1), path };
//   });

//   return (
//     <nav aria-label="breadcrumb">
//       <ol className="flex items-center space-x-2">
//         {breadcrumbItems.map((item, index) => (
//           <li
//             key={index}
//             className={`breadcrumb-item flex items-center ${
//               index === breadcrumbItems.length - 1 ? 'font-bold' : ''
//             }`}
//             aria-current={
//               index === breadcrumbItems.length - 1 ? 'page' : undefined
//             }
//           >
//             {index > 0 && <span className="mx-2">{'>'}</span>}
//             {index === breadcrumbItems.length - 1 ? (
//               <span className="text-darkGray">{item.name}</span>
//             ) : (
//               <Link to={item.path} className="text-blue-500">
//                 {item.name}
//               </Link>
//             )}
//           </li>
//         ))}
//       </ol>
//     </nav>
//   );
// };

// export default Breadcrumb;

import { Link, useLocation, useParams } from 'react-router-dom';
import { useIsMobile } from '../../../../hooks/useMediaQuery';

interface BreadcrumbItem {
  name: string;
  path: string;
}

const Breadcrumb = () => {
  const location = useLocation();
  const params = useParams();
  const isMobile = useIsMobile(); // Utilisez directement le hook ici

  // Diviser l'URL en segments et remplacer les %20 par _
  const pathnames = location.pathname
    .split('/')
    .filter((x) => x)
    .map((segment) => segment.replace(/%20/g, '_'));

  // Fonction pour obtenir le premier segment de la navigation
  const getMainBreadcrumbName = () => {
    if (pathnames.length > 0) {
      const firstSegment = pathnames[0];

      // Renommer selon le premier segment
      if (firstSegment === 'organisation') {
        return params.companyName || 'Organisation';
      } else if (firstSegment === 'events') {
        return 'Événements';
      } else if (firstSegment === 'options') {
        return 'Options';
      }
    }
    return 'Accueil'; // Nom par défaut si aucun segment
  };

  // Générer les éléments du breadcrumb en fonction des segments
  const breadcrumbItems: BreadcrumbItem[] = pathnames.map((segment, index) => {
    const path = `/${pathnames.slice(0, index + 1).join('/')}`;
    let name = segment;

    // Vérifie les segments pour les routes spécifiques
    if (segment === 'organisation' && params.companyName) {
      name = params.companyName; // Remplace par le nom de l'organisation
    } else if (segment === 'events') {
      name = 'Événements'; // Renomme pour les événements
    } else if (segment === 'options') {
      name = 'Options'; // Renomme pour les options
    }

    return { name: name.charAt(0).toUpperCase() + name.slice(1), path };
  });

  return (
    <nav aria-label="breadcrumb">
      <ol className="flex items-center space-x-2">
        {isMobile ? (
          <li className="breadcrumb-item font-bold">
            {getMainBreadcrumbName()}{' '}
            {/* Affiche le premier segment en fonction de l'URL */}
          </li>
        ) : (
          breadcrumbItems.map((item, index) => (
            <li
              key={index}
              className={`breadcrumb-item flex items-center ${
                index === breadcrumbItems.length - 1 ? 'font-bold' : ''
              }`}
              aria-current={
                index === breadcrumbItems.length - 1 ? 'page' : undefined
              }
            >
              {index > 0 && <span className="mx-2">{'>'}</span>}
              {index === breadcrumbItems.length - 1 ? (
                <span className="text-darkGray">{item.name}</span>
              ) : (
                <Link to={item.path} className="text-blue-500">
                  {item.name}
                </Link>
              )}
            </li>
          ))
        )}
      </ol>
    </nav>
  );
};

export default Breadcrumb;
