import { Event } from './Event';

export interface EventTag {
  id: number;
  name: string;
  events: Event[];
  createdAt: string;
  updatedAt: string;
}
