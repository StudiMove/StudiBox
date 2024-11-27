// import Container from '../../Container/Container'; // Container général pour le style

// import EventCard from '../../Card/EventCard';
// interface EventCardListProps {
//   data: any[]; // Les données des événements passées via props
//   onCardClick: (id: number) => void; // Fonction de callback lors du clic sur une carte
// }

// function EventCardList({ data, onCardClick }: EventCardListProps) {
//   const handleClick = (event: any) => {
//     onCardClick(event.id); // Appeler la fonction onCardClick avec l'ID de l'événement
//   };

//   return (
//     <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-5">
//       {/* Utilisation d'une grille pour l'affichage des cartes */}
//       {data.map((event) => (
//         <EventCard
//           key={event.id}
//           event={event}
//           onClick={() => handleClick(event)} // Clic sur une carte
//         />
//       ))}
//     </div>
//   );
// }

// export default EventCardList;

// import Container from '../../Container/Container'; // Container général pour le style
// import EventCard from '../../Card/EventCard';

// interface EventCardListProps {
//   data: any[]; // Les données des événements passées via props
//   onCardClick: (id: number) => void; // Fonction de callback lors du clic sur une carte
// }

// function EventCardList({ data, onCardClick }: EventCardListProps) {
//   const handleClick = (event: any) => {
//     onCardClick(event.id); // Appeler la fonction onCardClick avec l'ID de l'événement
//   };

//   return (
//     <Container
//       variant="desktop-layout-container"
//       className="pt-8 relative mb-16"
//     >
//       <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-5">
//         {/* Utilisation d'une grille pour l'affichage des cartes */}
//         {data.map((event) => (
//           <EventCard
//             key={event.id}
//             event={event}
//             onClick={() => handleClick(event)} // Clic sur une carte
//           />
//         ))}
//       </div>
//     </Container>
//   );
// }

// export default EventCardList;

import EventCard from '../../Card/EventCard';

interface EventCardListProps {
  data: any[];
  onCardClick: (id: number) => void;
}

function EventCardList({ data, onCardClick }: EventCardListProps) {
  const handleClick = (event: any) => {
    onCardClick(event.id);
  };

  return (
    <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-5 pt-8 relative mb-16">
      {data.map((event) => (
        <EventCard
          key={event.id}
          event={event}
          onClick={() => handleClick(event)}
        />
      ))}
    </div>
  );
}

export default EventCardList;
