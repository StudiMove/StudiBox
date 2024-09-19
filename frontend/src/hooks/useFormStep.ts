import { useContext } from 'react';
import { FormStepsContext } from '../contexts/FormStepsContext/FormStepsContext';
import { FormStepsContextType } from '../types/FormStepsContextType/FormStepsContextType';
import { FormStep } from '../types/FormStep/FormStepType';

const useFormSteps = <T>() => {
  const context = useContext(FormStepsContext) as FormStepsContextType<T>;

  if (!context) {
    throw new Error('useFormSteps must be used within a FormStepsProvider');
  }

  const { currentStep, setCurrentStep, formStepModel } = context;

  const nextStep = () => {
    if (!formStepModel) return;
    const allSteps: FormStep[] = formStepModel.getAllSteps();
    const currentIndex = allSteps.findIndex((step) => step.id === currentStep);
    if (currentIndex < allSteps.length - 1) {
      setCurrentStep(allSteps[currentIndex + 1].id);
    }
  };

  const prevStep = () => {
    if (!formStepModel) return;
    const allSteps: FormStep[] = formStepModel.getAllSteps();
    const currentIndex = allSteps.findIndex((step) => step.id === currentStep);
    if (currentIndex > 0) {
      setCurrentStep(allSteps[currentIndex - 1].id);
    }
  };

  return { currentStep, nextStep, prevStep };
};

export default useFormSteps;
