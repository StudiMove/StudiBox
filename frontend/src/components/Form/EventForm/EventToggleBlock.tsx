import { useSelector } from 'react-redux';
import { RootState } from '../../../../store';
import Container from '../../UI/Container/Container';
import ToggleButton from './ToggleButton';
import IconTitle from '../../UI/Typography/IconTitle';
import shareIcon from '../../Icons/share.svg';
import { EventForm } from '../../../types/api/event/EventTypes';

interface EventToggleBlockProps {
  formData: {
    isOnline: boolean;
    isPublic: boolean;
  };
  updateField: (field: keyof EventForm, value: boolean) => void;
}

const EventToggleBlock = ({ formData, updateField }: EventToggleBlockProps) => {
  const { isOnline } = formData;
  const userRole = useSelector((state: RootState) => state.auth.role);

  return (
    <>
      <IconTitle title="Share info" image={shareIcon} className="mb-6" />

      <Container variant="two-input-row" className="mb-8">
        <div className="flex flex-col relative flex-1">
          <ToggleButton
            label="En ligne"
            value={isOnline}
            onChange={(newValue) => updateField('isOnline', newValue)} // Nouveau nom de paramètre
          />
        </div>
        {userRole !== 'business' && ( // Afficher le toggle si l'utilisateur n'est pas de rôle business
          <div className="flex flex-col relative flex-1">
            <ToggleButton
              label="Ouvert à tous"
              value={formData.isPublic}
              onChange={(newValue) => updateField('isPublic', newValue)} // Nouveau nom de paramètre
            />
          </div>
        )}
      </Container>
    </>
  );
};

export default EventToggleBlock;
