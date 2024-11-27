import { Event } from './Event';

export interface EventCategory {
  id: number;
  name: string;
  events: Event[];
  createdAt: string;
  updatedAt: string;
}
