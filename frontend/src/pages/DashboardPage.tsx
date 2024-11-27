// LoginPage.tsx
import Container from '../components/UI/Container/Container';
import { useIsMobile } from '../hooks/useMediaQuery';
import EventStatsComponentDash from '../components/Form/Event/EventStatsComponentDash';
const DashboardPage = () => {
  const isMobile = useIsMobile();
  return (
    <div className={` ${isMobile ? 'px-8' : ''}`}>
      <Container variant="two-input-row">
        <EventStatsComponentDash />
      </Container>
    </div>
  );
};
export default DashboardPage;
