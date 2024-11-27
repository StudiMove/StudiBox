// LocationBlock.tsx
import { EventLocation } from '../../../types/api/event/EventTypes';
import InputField from '../../UI/Inputs/InputField/InputField';
import Container from '../../UI/Container/Container';
import IconTitle from '../../UI/Typography/IconTitle';
import mapIcon from '../../Icons/mapIcon.svg'; // Icone pour la localisation

interface LocationBlockProps {
  location: EventLocation;
  updateField: (location: EventLocation) => void;
}

const LocationBlock = ({ location, updateField }: LocationBlockProps) => {
  const handleInputChange =
    (field: keyof EventLocation) => (value: string | number) => {
      updateField({ ...location, [field]: value });
    };

  return (
    <div className="mb-8">
      {/* Titre avec icône */}
      <IconTitle className="mb-6" title="Localisation" image={mapIcon} />

      {/* Conteneur avec les champs pour la localisation */}
      <Container variant="two-input-row">
        <InputField
          type="text"
          label="Ville"
          value={location.city}
          onChange={handleInputChange('city')}
          size="small"
          isRequired={true}
        />
        <InputField
          type="text"
          label="Région"
          value={location.region}
          onChange={handleInputChange('region')}
          size="small"
          isRequired={true}
        />
      </Container>

      <Container variant="two-input-row">
        <InputField
          type="text"
          label="Code postal"
          value={location.postcode}
          onChange={handleInputChange('postcode')}
          size="small"
          isRequired={true}
        />
        <InputField
          type="text"
          label="Pays"
          value={location.country}
          onChange={handleInputChange('country')}
          size="small"
          isRequired={true}
        />
      </Container>

      <Container variant="two-input-row-1">
        <InputField
          type="text"
          label="Adresse"
          value={location.address}
          onChange={handleInputChange('address')}
          size="small"
          isRequired={true}
        />
      </Container>
    </div>
  );
};

export default LocationBlock;
