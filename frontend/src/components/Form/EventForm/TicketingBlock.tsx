import ToggleButton from './ToggleButton';
import InputField from '../../UI/Inputs/InputField/InputField';
import Container from '../../UI/Container/Container';
import IconTitle from '../../UI/Typography/IconTitle';
import cartIcon from '../../Icons/shopping-cart.svg';
import { EventForm } from '../../../types/api/event/EventTypes';

interface TicketingBlockProps {
  use_studibox: boolean;
  ticketPrice: number;
  ticketStock: number;
  externalTicketingUrl: string;
  updateField: (field: keyof EventForm, value: any) => void; // Modifié pour correspondre à EventForm
}

const TicketingBlock = ({
  use_studibox,
  ticketPrice,
  ticketStock,
  externalTicketingUrl,
  updateField,
}: TicketingBlockProps) => {
  const handleInputChange =
    (field: keyof EventForm) => (value: string | number) => {
      updateField(field, value);
    };
  console.log(use_studibox);

  return (
    <>
      <IconTitle title="Billetterie" image={cartIcon} className="mb-6" />
      <div className="mt-6 bg-lightBlue p-8 rounded-lg mb-8">
        <ToggleButton
          label="Utiliser la billetterie de Studibox ?"
          value={use_studibox}
          onChange={(val) => updateField('use_studibox', val)}
        />

        <Container variant="two-input-row" className="mt-4">
          <InputField
            type="text"
            label="URL billetterie externe"
            value={externalTicketingUrl}
            onChange={handleInputChange('externalTicketingUrl')}
            size="small"
            isEditable={!use_studibox}
          />
        </Container>

        <Container variant="two-input-row">
          <InputField
            type="number"
            label="A partir de (€)"
            value={ticketPrice}
            onChange={handleInputChange('ticketPrice')}
            size="small"
            isEditable={!use_studibox}
          />
          <InputField
            type="number"
            label="Stock"
            value={ticketStock}
            onChange={handleInputChange('ticketStock')}
            size="small"
            isEditable={!use_studibox}
          />
        </Container>
      </div>
    </>
  );
};

export default TicketingBlock;
