import { useEffect } from 'react';
import { EventOption } from '../../../types/EventForm/EventForm';
import InputField from '../../UI/Inputs/InputField/InputField';
import Container from '../../UI/Container/Container';
import IconTitle from '../../UI/Typography/IconTitle';
import optionIcon from '../../Icons/plus-circle.svg';
import redTrash from '../../Icons/trash/red-trash.svg';
import blackTrash from '../../Icons/trash/black-trash.svg';
import H3 from '../../UI/Typography/H3';
import PrimaryButtonOutline from '../../UI/Button/PrimaryButtonOutline';

interface OptionsBlockProps {
  options: EventOption[];
  updateField: (options: EventOption[]) => void;
}

const OptionsBlock = ({ options, updateField }: OptionsBlockProps) => {
  useEffect(() => {
    // S'assurer qu'il y a toujours au moins une option par défaut
    if (options.length === 0) {
      updateField([{ title: '', description: '', price: 0, stock: 0 }]);
    }
  }, [options.length, updateField]);

  const addOption = () => {
    updateField([
      ...options,
      { title: '', description: '', price: 0, stock: 0 },
    ]);
  };

  const updateOption = (
    index: number,
    field: keyof EventOption,
    value: string | number,
  ) => {
    const newOptions = [...options];
    newOptions[index] = { ...newOptions[index], [field]: value };
    updateField(newOptions);
  };

  const removeOption = (index: number) => {
    // Ne pas permettre de supprimer si c'est la seule option
    if (options.length > 1) {
      const newOptions = options.filter((_, i) => i !== index);
      updateField(newOptions);
    }
  };

  return (
    <div className="mb-8">
      <IconTitle title="Options" image={optionIcon} className="mb-6" />
      {options.map((option, index) => (
        <div key={index} className="mb-4 bg-lightBlue p-8 rounded-lg">
          <div className="flex justify-between">
            <H3 className="text-primary">Option {index + 1}</H3>
            {options.length > 1 && ( // Afficher l'icône de suppression seulement s'il y a plus d'une option
              <img
                src={blackTrash}
                alt="trash-icon"
                className="cursor-pointer"
                onMouseEnter={(e) => (e.currentTarget.src = redTrash)}
                onMouseLeave={(e) => (e.currentTarget.src = blackTrash)}
                onClick={() => removeOption(index)}
              />
            )}
          </div>
          <InputField
            type="text"
            label="Titre"
            value={option.title}
            onChange={(value) => updateOption(index, 'title', value)}
            size="small"
          />
          <Container variant="two-input-row">
            <InputField
              type="number"
              label="Prix"
              value={option.price}
              onChange={(value) => updateOption(index, 'price', Number(value))}
              size="small"
            />
            <InputField
              type="number"
              label="Stock"
              value={option.stock}
              onChange={(value) => updateOption(index, 'stock', Number(value))}
              size="small"
            />
          </Container>
          <InputField
            type="text"
            label="Description"
            value={option.description}
            onChange={(value) => updateOption(index, 'description', value)}
            size="small"
          />
        </div>
      ))}
      <PrimaryButtonOutline
        label="Ajouter une option"
        onClick={addOption}
        className="mt-4"
      />
    </div>
  );
};

export default OptionsBlock;
