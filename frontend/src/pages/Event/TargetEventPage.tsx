import Tabs from '../../components/UI/Tab/Tabs';
import EventFormComponent from '../../components/Form/EventForm/EventFormComponent';
import EventStatsComponent from '../../components/Form/Event/EventStatsComponent';
interface TargetEventPageProps {
  event_id: number; // ID de l'événement passé en props
}

const TargetEventPage = ({ event_id }: TargetEventPageProps) => {
  console.log(event_id);
  console.log(event_id);
  console.log(event_id);
  console.log(event_id);
  const tabs = [
    {
      label: 'Détails',
      component: (
        <EventFormComponent targetEventId={event_id} isUpdate={true} />
      ),
    },
    {
      label: 'Statistiques',
      component: <EventStatsComponent />,
    },
  ];

  return (
    <div className={`w-full`}>
      <Tabs tabs={tabs} />
    </div>
  );
};

export default TargetEventPage;
