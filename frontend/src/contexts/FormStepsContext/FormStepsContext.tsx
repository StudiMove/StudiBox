import { createContext, useState, ReactNode, useContext } from 'react';
import { FormStepsContextType } from '../../types/FormStepsContextType/FormStepsContextType';
import { FormStepModel } from '../../models/FormStep/FormStepModel';

type FormStepsProviderProps<T> = {
  children: ReactNode;
  initialData: T;
  formStepModel: FormStepModel; // Assurez-vous que formStepModel est requis
  initialStep: string; // Ajout de la propriété initialStep
};

export const FormStepsContext = createContext<
  FormStepsContextType<any> | undefined
>(undefined);

export const FormStepsProvider = <T,>({
  children,
  initialData,
  formStepModel,
  initialStep, // Ajout de la propriété initialStep
}: FormStepsProviderProps<T>) => {
  const [currentStep, setCurrentStep] = useState<string>(initialStep); // Utiliser initialStep
  const [formData, setFormData] = useState<T>(initialData);

  const nextStep = () => {
    const allSteps = formStepModel.getAllSteps();
    const currentIndex = allSteps.findIndex((step) => step.id === currentStep);
    if (currentIndex < allSteps.length - 1) {
      setCurrentStep(allSteps[currentIndex + 1].id);
    }
  };

  const prevStep = () => {
    const allSteps = formStepModel.getAllSteps();
    const currentIndex = allSteps.findIndex((step) => step.id === currentStep);
    if (currentIndex > 0) {
      setCurrentStep(allSteps[currentIndex - 1].id);
    }
  };

  return (
    <FormStepsContext.Provider
      value={{
        currentStep,
        setCurrentStep,
        formData,
        setFormData,
        formStepModel,
        nextStep,
        prevStep,
      }}
    >
      {children}
    </FormStepsContext.Provider>
  );
};

export const useFormSteps = <T,>() => {
  const context = useContext(FormStepsContext) as FormStepsContextType<T>;

  if (context === undefined) {
    throw new Error('useFormSteps must be used within a FormStepsProvider');
  }

  return context;
};
