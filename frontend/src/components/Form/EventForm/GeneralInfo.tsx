import { EventForm } from '../../../types/api/event/EventTypes';
import IconTitle from '../../UI/Typography/IconTitle';
import InputField from '../../UI/Inputs/InputField/InputField';
import InputFieldMultiSelect from '../../UI/Inputs/InputFieldList/InputFieldMultiSelect';
import Container from '../../UI/Container/Container';
import GeneralInfoIcon from '../../Icons/cube-transparent.svg';
import DatePicker, { registerLocale } from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';
import { fr } from 'date-fns/locale';
import calendarIcon from '../../Icons/calendar.svg';
import clockIcon from '../../Icons/clock.svg';
import { useIsMobile } from '../../../hooks/useMediaQuery';

registerLocale('fr', fr);

interface GeneralInfoProps {
  formData: EventForm;
  updateField: (
    field: keyof EventForm,
    value: string | number | boolean | Date | string[] | null,
  ) => void;
  categoryOptions: { value: string; label: string }[];
  onCategoryChange: (
    selectedOptions: { value: string; label: string }[],
  ) => void;
  tagOptions: { value: string; label: string }[];
  onTagChange: (selectedOptions: { value: string; label: string }[]) => void;
}

const GeneralInfo = ({
  formData,
  updateField,
  categoryOptions,
  onCategoryChange,
  tagOptions,
  onTagChange,
}: GeneralInfoProps) => {
  const handleInputChange =
    (key: keyof EventForm) =>
    (value: string | number | Date | string[] | null) => {
      updateField(key, value);
    };

  const isMobile = useIsMobile();

  return (
    <div className="mb-8">
      <IconTitle
        className="mb-6"
        title="Informations Générales"
        image={GeneralInfoIcon}
      />

      <Container variant="two-input-row">
        <InputField
          type="text"
          label="Titre"
          value={formData.title}
          onChange={handleInputChange('title')}
          size="small"
          isRequired={true}
        />
        <InputField
          type="text"
          label="Sous-titre"
          value={formData.subtitle}
          onChange={handleInputChange('subtitle')}
          size="small"
          isRequired={true}
        />
      </Container>

      <Container variant="two-input-row" className="mb-6">
        <InputFieldMultiSelect
          label="Catégorie"
          value={formData.categories.map((cat) => ({
            value: cat,
            label: cat.charAt(0).toUpperCase() + cat.slice(1),
          }))}
          options={categoryOptions.filter(
            (category) => !formData.categories.includes(category.value),
          )}
          onChange={onCategoryChange}
        />
        <InputFieldMultiSelect
          label="Tag"
          value={formData.tags.map((tag) => ({
            value: tag,
            label: tag.charAt(0).toUpperCase() + tag.slice(1),
          }))}
          options={tagOptions.filter(
            (tag) => !formData.tags.includes(tag.value),
          )}
          onChange={onTagChange}
        />
      </Container>

      <div>
        <h3 className="mb-4">Horaires et dates</h3>

        {isMobile ? (
          <>
            <Container variant="two-input-row">
              <label className="mt-2">Date de début</label>

              <div className="flex items-center space-x-2 bg-lightBlue p-2 rounded-lg">
                <img
                  src={calendarIcon}
                  alt="calendar icon"
                  className="h-6 w-6"
                />
                <DatePicker
                  selected={formData.startDate || undefined}
                  onChange={(date) =>
                    handleInputChange('startDate')(date || null)
                  }
                  dateFormat="dd/MM/yyyy"
                  locale="fr"
                  className="bg-transparent border-none focus:outline-none"
                />
              </div>
            </Container>
            <Container variant="two-input-row">
              <label className="mt-2">Heure de début</label>

              <div className="flex items-center space-x-2 bg-lightBlue p-2 rounded-lg">
                <img src={clockIcon} alt="clock icon" className="h-6 w-6" />
                <DatePicker
                  selected={formData.startTime || undefined}
                  onChange={(time) =>
                    handleInputChange('startTime')(time || null)
                  }
                  showTimeSelect
                  showTimeSelectOnly
                  timeIntervals={30}
                  timeCaption="Heure"
                  dateFormat="HH:mm"
                  locale="fr"
                  timeFormat="HH:mm"
                  className="bg-transparent border-none focus:outline-none"
                />
              </div>
            </Container>
            <Container variant="two-input-row">
              <label className="mt-2">Date de fin</label>

              <div
                className={`flex items-center space-x-4 ${
                  !formData.startDate || !formData.startTime ? 'opacity-50' : ''
                }`}
              >
                <div className="flex items-center space-x-2 bg-lightBlue p-2 rounded-lg">
                  <img
                    src={calendarIcon}
                    alt="calendar icon"
                    className="h-6 w-6"
                  />
                  <DatePicker
                    selected={formData.endDate || undefined}
                    onChange={(date) =>
                      handleInputChange('endDate')(date || null)
                    }
                    dateFormat="dd/MM/yyyy"
                    locale="fr"
                    minDate={formData.startDate || undefined}
                    className="bg-transparent border-none focus:outline-none"
                    disabled={!formData.startDate || !formData.startTime}
                  />
                </div>
              </div>
            </Container>
            <Container variant="two-input-row">
              <label className="mt-2">Heure de fin</label>
              <div className="flex items-center space-x-2 bg-lightBlue p-2 rounded-lg">
                <img src={clockIcon} alt="clock icon" className="h-6 w-6" />
                <DatePicker
                  selected={formData.endTime || undefined}
                  onChange={(time) =>
                    handleInputChange('endTime')(time || null)
                  }
                  showTimeSelect
                  showTimeSelectOnly
                  timeIntervals={30}
                  timeCaption="Heure"
                  dateFormat="HH:mm"
                  locale="fr"
                  timeFormat="HH:mm"
                  minTime={
                    formData.startDate === formData.endDate
                      ? formData.startTime || undefined
                      : undefined
                  }
                  className="bg-transparent border-none focus:outline-none"
                  disabled={!formData.startTime || !formData.startDate}
                />
              </div>
            </Container>
          </>
        ) : (
          <Container variant="two-input-row">
            <div className="flex flex-col relative flex-1">
              <label>Date de début</label>

              <div className="flex items-center space-x-4 gap-4">
                <div className="flex items-center space-x-2 bg-lightBlue p-2 rounded-lg">
                  <img
                    src={calendarIcon}
                    alt="calendar icon"
                    className="h-6 w-6"
                  />
                  <DatePicker
                    selected={formData.startDate || undefined}
                    onChange={(date) =>
                      handleInputChange('startDate')(date || null)
                    }
                    dateFormat="dd/MM/yyyy"
                    locale="fr"
                    className="bg-transparent border-none focus:outline-none"
                  />
                </div>

                <div className="flex items-center space-x-2 bg-lightBlue p-2 rounded-lg">
                  <img src={clockIcon} alt="clock icon" className="h-6 w-6" />
                  <DatePicker
                    selected={formData.startTime || undefined}
                    onChange={(time) =>
                      handleInputChange('startTime')(time || null)
                    }
                    showTimeSelect
                    showTimeSelectOnly
                    timeIntervals={30}
                    timeCaption="Heure"
                    dateFormat="HH:mm"
                    locale="fr"
                    timeFormat="HH:mm"
                    className="bg-transparent border-none focus:outline-none"
                  />
                </div>
              </div>
            </div>

            <div className="flex flex-col relative flex-1">
              <label>Date de fin</label>
              <div
                className={`flex items-center space-x-4 ${
                  !formData.startDate || !formData.startTime ? 'opacity-50' : ''
                }`}
              >
                <div className="flex items-center space-x-2 bg-lightBlue p-2 rounded-lg">
                  <img
                    src={calendarIcon}
                    alt="calendar icon"
                    className="h-6 w-6"
                  />
                  <DatePicker
                    selected={formData.endDate || undefined}
                    onChange={(date) =>
                      handleInputChange('endDate')(date || null)
                    }
                    dateFormat="dd/MM/yyyy"
                    locale="fr"
                    minDate={formData.startDate || undefined}
                    className="bg-transparent border-none focus:outline-none"
                    disabled={!formData.startDate || !formData.startTime}
                  />
                </div>
                <div className="flex items-center space-x-2 bg-lightBlue p-2 rounded-lg">
                  <img src={clockIcon} alt="clock icon" className="h-6 w-6" />
                  <DatePicker
                    selected={formData.endTime || undefined}
                    onChange={(time) =>
                      handleInputChange('endTime')(time || null)
                    }
                    showTimeSelect
                    showTimeSelectOnly
                    timeIntervals={30}
                    timeCaption="Heure"
                    dateFormat="HH:mm"
                    locale="fr"
                    timeFormat="HH:mm"
                    minTime={
                      formData.startDate === formData.endDate
                        ? formData.startTime || undefined
                        : undefined
                    }
                    className="bg-transparent border-none focus:outline-none"
                    disabled={!formData.startTime || !formData.startDate}
                  />
                </div>
              </div>
            </div>
          </Container>
        )}
      </div>
    </div>
  );
};

export default GeneralInfo;
