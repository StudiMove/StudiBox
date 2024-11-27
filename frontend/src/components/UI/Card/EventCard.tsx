import defaultBG from '../../Icons/defaultBG.svg';
import iconBlackClock from '../../Icons/iconBlackClock.svg';
import iconBlueCalendar from '../../Icons/iconBlueCalendar.svg';
import H4 from '../Typography/H4';
import SpanMedium from '../Typography/SpanMedium';

interface EventCardProps {
  event: {
    id: number;
    title: string;
    start_date: string;
    start_time?: string;
    image_urls?: string;
    is_online: boolean;
  };
  onClick: () => void;
}

function EventCard({ event, onClick }: EventCardProps) {
  const formattedStartDate = new Date(event.start_date).toLocaleDateString(
    'fr-FR',
    {
      day: 'numeric',
      month: 'long',
      year: 'numeric',
    },
  );

  const formattedStartTime = event.start_time
    ? new Date(event.start_time).toLocaleTimeString('fr-FR', {
        hour: '2-digit',
        minute: '2-digit',
      })
    : 'Heure non définie';

  const firstImageURL = event.image_urls || defaultBG; // Utiliser une image par défaut si image_urls est vide

  return (
    <div
      className="relative border rounded-lg shadow-lg hover:shadow-xl transition-shadow duration-300 cursor-pointer overflow-hidden"
      onClick={onClick}
    >
      {/* Section pour l'image */}
      <div className="relative">
        <img
          src={firstImageURL}
          alt={event.title || 'Titre non défini'}
          className="h-60 w-full object-cover rounded-t-lg"
        />

        {/* Dégradé noir en bas de la carte */}
        <div className="absolute inset-x-0 bottom-0 bg-gradient-to-t from-black/80 to-transparent p-4 z-10">
          {/* Titre de l'événement */}
          <H4 className="text-lg font-semibold !text-white !mb-2.5">
            {event.title || 'Titre non défini'}
          </H4>

          {/* Date et heure */}
          <div className="flex items-center justify-between !text-white">
            <div className="flex items-center space-x-1">
              <img src={iconBlueCalendar} alt="Calendar icon" />
              <SpanMedium className="!text-white">
                {formattedStartDate}
              </SpanMedium>
            </div>
            <div className="flex items-center space-x-1">
              <img src={iconBlackClock} alt="Clock icon" />
              <SpanMedium className="!text-white !text-[13px]">
                {formattedStartTime}
              </SpanMedium>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default EventCard;
