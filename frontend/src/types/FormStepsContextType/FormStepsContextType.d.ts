// src/types/FormStepsContextType/FormStepsContextType.d.ts
import { Dispatch, SetStateAction } from 'react';
import { FormStepModel } from '../../models/FormStep/FormStepModel';

export interface FormStepsContextType<T> {
  currentStep: string; // Assurez-vous que le type correspond à celui utilisé dans votre contexte
  setCurrentStep: Dispatch<SetStateAction<string>>;
  formData: T;
  setFormData: Dispatch<SetStateAction<T>>;
  formStepModel?: FormStepModel;
  nextStep?: () => void; // Ajout des méthodes
  prevStep?: () => void;
}
