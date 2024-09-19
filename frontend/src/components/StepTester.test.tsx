import { render, screen, fireEvent } from '@testing-library/react';
import { FormStepsProvider } from '../contexts/FormStepsContext/FormStepsContext';
import { FormStepModel } from '../models/FormStep/FormStepModel';
import StepTester from './StepTester';

// Mocking the FormStepModel
const mockFormStepModel = {
  getAllSteps: () => [
    { id: '1', component: <div>Step 1 Content</div> },
    { id: '2', component: <div>Step 2 Content</div> },
  ],
};

const initialData = { email: '', password: '' };

test('renders StepTester and navigates steps', () => {
  render(
    <FormStepsProvider
      initialData={initialData}
      formStepModel={mockFormStepModel as unknown as FormStepModel}
      initialStep="1" // Ajout de la propriété initialStep
    >
      <StepTester />
    </FormStepsProvider>,
  );

  // Vérifier que l'étape actuelle est correcte
  expect(screen.getByText('Current Step: 1')).toBeInTheDocument(); // Ajuster pour correspondre au texte réel

  // Vérifier la présence du bouton "Next Step"
  expect(screen.getByText('Next Step')).toBeInTheDocument();

  // Cliquez sur le bouton "Next Step" pour avancer
  fireEvent.click(screen.getByText('Next Step'));

  // Vérifier que l'étape a changé
  expect(screen.getByText('Current Step: 2')).toBeInTheDocument();

  // Vérifier la présence du bouton "Previous Step"
  expect(screen.getByText('Previous Step')).toBeInTheDocument();

  // Cliquez sur le bouton "Previous Step" pour revenir
  fireEvent.click(screen.getByText('Previous Step'));

  // Vérifier que l'étape est revenue à l'étape précédente
  expect(screen.getByText('Current Step: 1')).toBeInTheDocument();
});
