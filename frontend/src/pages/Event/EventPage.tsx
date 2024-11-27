import { useState } from 'react';
import { useSelector } from 'react-redux';
import TabsData from '../../components/UI/Tab/TabsData';
import EventCardList from '../../components/UI/List/EventCardList/EventCardList';
import TargetEventPage from './TargetEventPage';
import AddButton from '../../components/UI/Button/AddButton';
import {
  getAllEvents,
  getUserEvents,
  getUserEventsByOrganisation,
} from '../../services/EventService';

interface SimpleEvent {
  id: number;
  title: string;
  start_date: string;
  start_time: string;
  is_online: boolean;
  is_activated: boolean;
}

interface EventPageProps {
  targetId?: number;
}

const EventPage = ({ targetId }: EventPageProps) => {
  const [selectedEventId, setSelectedEventId] = useState<number | null>(null);
  const userRole = useSelector((state: any) => state.auth.role);

  const handleCardClick = (id: number) => setSelectedEventId(id);

  // Adaptation des données dans le format SimpleEvent
  const loadAllEvents = async (): Promise<SimpleEvent[]> => {
    const events = await (targetId
      ? getUserEventsByOrganisation(targetId)
      : userRole === 'admin'
      ? getAllEvents()
      : getUserEvents());

    // Map des événements pour correspondre à SimpleEvent
    return events.map((event: any) => {
      let imageUrl = ''; // Initialiser imageUrl
      if (event.image_urls) {
        try {
          // Transforme la chaîne en un tableau en utilisant split
          const imageUrlsArray = event.image_urls.split(', ');
          imageUrl = imageUrlsArray.length > 0 ? imageUrlsArray[0] : ''; // Prendre le premier élément ou une chaîne vide
        } catch (error) {
          console.error(
            "Erreur lors de la transformation de l'URL d'image:",
            error,
          );
        }
      }

      return {
        id: event.id,
        title: event.title,
        start_date: event.start_date,
        start_time: event.start_time,
        is_online: event.is_online,
        image_urls: imageUrl,
        is_activated: event.is_activated,
      };
    });
  };

  // Fonction pour charger uniquement les événements en ligne
  const loadOnlineEvents = async (): Promise<SimpleEvent[]> => {
    const allEvents = await loadAllEvents();
    return allEvents.filter((event) => event.is_online);
  };

  // Fonction pour charger uniquement les événements hors ligne
  const loadOfflineEvents = async (): Promise<SimpleEvent[]> => {
    const allEvents = await loadAllEvents();
    return allEvents.filter((event) => !event.is_online);
  };
  // Fonction pour charger uniquement les événements hors ligne
  const loadPassedEvents = async (): Promise<SimpleEvent[]> => {
    const allEvents = await loadAllEvents();
    return allEvents.filter((event) => !event.is_activated && !event.is_online);
  };
  return (
    <div className={`${targetId ? '' : 'px-8'}`}>
      {selectedEventId ? (
        <TargetEventPage event_id={selectedEventId} />
      ) : (
        <>
          <div>
            <TabsData
              tabs={[
                { label: 'Tout', dataLoader: loadAllEvents },
                { label: 'En ligne', dataLoader: loadOnlineEvents },
                { label: 'Hors ligne', dataLoader: loadOfflineEvents },
                { label: 'Passé / Annulé', dataLoader: loadPassedEvents },
              ]}
              headers={[]}
              DataTableComponent={(props) => (
                <EventCardList
                  data={props.data}
                  onCardClick={handleCardClick}
                />
              )}
              mapData={(data: SimpleEvent[]) => data}
              filterKey="title"
            />
          </div>
          <div className="fixed right-4 bottom-4">
            <AddButton navigateTo="/events/createNewEvent" />
          </div>
        </>
      )}
    </div>
  );
};

export default EventPage;
