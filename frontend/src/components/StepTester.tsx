import React from 'react';
import {
  FormStepsProvider,
  useFormSteps,
} from '../contexts/FormStepsContext/FormStepsContext';
import { FormStepModel } from '../models/FormStep/FormStepModel';

// Composant TestComponent qui utilise le contexte des étapes du formulaire
const TestComponent = () => {
  const { currentStep, nextStep, prevStep, formData, setFormData } =
    useFormSteps<{ email: string; password: string }>();

  return (
    <div>
      <p>Current Step: {currentStep}</p>
      <button onClick={prevStep}>Previous Step</button>
      <button onClick={nextStep}>Next Step</button>

      {/* Afficher et mettre à jour les données du formulaire en fonction de l'étape */}
      {currentStep === '1' && (
        <div>
          <label>Email:</label>
          <input
            type="email"
            value={formData.email}
            onChange={(e) =>
              setFormData({ ...formData, email: e.target.value })
            }
          />
        </div>
      )}
      {currentStep === '2' && (
        <div>
          <label>Password:</label>
          <input
            type="password"
            value={formData.password}
            onChange={(e) =>
              setFormData({ ...formData, password: e.target.value })
            }
          />
        </div>
      )}
      {/* Ajoutez d'autres champs de formulaire pour les autres étapes si nécessaire */}
    </div>
  );
};

// Composant StepTester qui inclut le FormStepsProvider et TestComponent
const StepTester: React.FC = () => {
  const steps = [
    { id: '1', component: <div>Step 1</div> },
    { id: '2', component: <div>Step 2</div> },
    { id: '3', component: <div>Step 3</div> },
  ];

  return (
    <FormStepsProvider
      initialData={{ email: '', password: '' }}
      formStepModel={new FormStepModel(steps)}
      initialStep="1" // Assurez-vous que vous passez l'étape initiale
    >
      <TestComponent />
    </FormStepsProvider>
  );
};

export default StepTester;
