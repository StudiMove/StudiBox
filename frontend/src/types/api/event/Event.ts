// src/types/Event.ts
export interface Event {
  id: number;
  title: string;
  start_date: string;
  end_date: string;
  is_online: boolean;
  is_visible?: boolean;
}
