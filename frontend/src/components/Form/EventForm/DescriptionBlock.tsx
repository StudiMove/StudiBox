import { useEffect } from 'react';
import InputField from '../../UI/Inputs/InputField/InputField';
import IconTitle from '../../UI/Typography/IconTitle';
import descriptionIcon from '../../Icons/plus-circle.svg';
import redTrash from '../../Icons/trash/red-trash.svg';
import blackTrash from '../../Icons/trash/black-trash.svg';
import H3 from '../../UI/Typography/H3';
import PrimaryButtonOutline from '../../UI/Button/PrimaryButtonOutline';
import { EventDescription } from '../../../types/api/event/EventTypes';

interface DescriptionBlockProps {
  descriptions: EventDescription[];
  updateField: (descriptions: EventDescription[]) => void;
}

const DescriptionBlock = ({
  descriptions,
  updateField,
}: DescriptionBlockProps) => {
  useEffect(() => {
    if (descriptions.length === 0) {
      updateField([{ title: '', description: '' }]);
    }
  }, [descriptions.length, updateField]);

  const addDescription = () => {
    updateField([...descriptions, { title: '', description: '' }]);
  };

  const updateDescription = (
    index: number,
    field: keyof EventDescription,
    value: string,
  ) => {
    const newDescriptions = [...descriptions];
    newDescriptions[index] = { ...newDescriptions[index], [field]: value };
    updateField(newDescriptions);
  };

  const removeDescription = (index: number) => {
    if (descriptions.length > 1) {
      const newDescriptions = descriptions.filter((_, i) => i !== index);
      updateField(newDescriptions);
    }
  };

  return (
    <div className="mb-8">
      <IconTitle
        title="Descriptions"
        image={descriptionIcon}
        className="mb-6"
      />
      {descriptions.map((description, index) => (
        <div key={index} className="mb-4 bg-lightBlue p-8 rounded-lg">
          <div className="flex justify-between">
            <H3 className="text-primary">Description {index + 1}</H3>
            {descriptions.length > 1 && (
              <img
                src={blackTrash}
                alt="trash-icon"
                className="cursor-pointer"
                onMouseEnter={(e) => (e.currentTarget.src = redTrash)}
                onMouseLeave={(e) => (e.currentTarget.src = blackTrash)}
                onClick={() => removeDescription(index)}
              />
            )}
          </div>
          <InputField
            type="text"
            label="Titre"
            value={description.title}
            onChange={(value) =>
              updateDescription(index, 'title', value as string)
            }
            size="small"
          />
          <InputField
            type="text"
            label="Description"
            value={description.description}
            onChange={(value) =>
              updateDescription(index, 'description', value as string)
            }
            size="small"
          />
        </div>
      ))}
      <PrimaryButtonOutline
        label="Ajouter une description"
        onClick={addDescription}
        className="mt-4"
      />
    </div>
  );
};

export default DescriptionBlock;
