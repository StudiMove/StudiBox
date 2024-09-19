// src/models/FormStep/FormStepModel.ts
import { FormStep } from '../../types/FormStep/FormStepType';

export class FormStepModel {
  private steps: FormStep[];

  constructor(steps: FormStep[]) {
    this.steps = steps;
  }

  getAllSteps(): FormStep[] {
    return this.steps;
  }
}
