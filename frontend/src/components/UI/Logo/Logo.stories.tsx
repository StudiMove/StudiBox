import Logo from './Logo';
export default {
  title: 'UI/Logo', // Titre de la catégorie dans Storybook
  component: Logo, // Le composant à documenter
};

// Template pour les différentes variantes de Logo
const Template = (args) => <Logo {...args} />;

// Story pour chaque variante
export const IconWithText = Template.bind({});
IconWithText.args = {
  variant: 'iconWithText',
};

export const TextOnly = Template.bind({});
TextOnly.args = {
  variant: 'textOnly',
};

export const IconOnly = Template.bind({});
IconOnly.args = {
  variant: 'iconOnly',
};

export const LargeIconWithText = Template.bind({});
LargeIconWithText.args = {
  variant: 'largeIconWithText',
};
