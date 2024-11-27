import { Meta, StoryFn } from '@storybook/react'; // Utilisez StoryFn si Story n'est pas disponible
import InputField from './InputField';
import { InputFieldProps } from '../../../../types/InputField/InputFieldProps';
export default {
  title: 'Components/InputField',
  component: InputField,
  parameters: {
    docs: {
      description: {
        component:
          "Un champ de saisie réutilisable avec des tailles ajustables. Il peut être utilisé pour divers types d'inputs tels que texte, email et mot de passe.",
      },
    },
  },
  argTypes: {
    size: {
      control: { type: 'select', options: ['small', 'medium', 'large'] },
      description: 'Taille du champ de saisie.',
    },
    label: {
      control: 'text',
      description: "Texte affiché dans l'étiquette.",
    },
    type: {
      control: 'text',
      description: "Type d'input (ex: text, email, password).",
    },
    value: {
      control: 'text',
      description: 'Valeur actuelle du champ.',
    },
    onChange: {
      action: 'changed',
      description: 'Fonction appelée lors du changement de valeur.',
    },
  },
} as Meta<typeof InputField>;

const Template: StoryFn<InputFieldProps> = (args) => <InputField {...args} />;

export const Default = Template.bind({});
Default.args = {
  label: 'Email',
  type: 'email',
  value: '',
  size: 'medium',
};

/**
 * Exemple d'utilisation :
 *
 * <YourForm>
 *   <InputField label="Email" type="email" value={email} size="medium" onChange={setEmail} />
 * </YourForm>
 */

export const Small = Template.bind({});
Small.args = {
  label: 'Small Input',
  type: 'text',
  value: '',
  size: 'small',
};

export const Large = Template.bind({});
Large.args = {
  label: 'Large Input',
  type: 'text',
  value: '',
  size: 'large',
};

export const Password = Template.bind({});
Password.args = {
  label: 'Password',
  type: 'password',
  value: '',
  size: 'medium',
};
