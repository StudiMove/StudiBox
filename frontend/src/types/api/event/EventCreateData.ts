// src/types/api/event/EventCreateData.ts

export interface EventLocation {
  address: string;
  city: string;
  postcode: string;
  region: string;
  country: string;
}

export interface EventOption {
  title: string;
  description: string;
  price: number;
  stock: number;
}

export interface EventTarif {
  title: string;
  description: string;
  price: number;
  stock: number;
}

export interface EventDescription {
  title: string;
  description: string;
}

export interface EventCreateData {
  title: string;
  subtitle: string;
  start_date: string;
  end_date: string;
  start_time: string;
  end_time: string;
  is_online: boolean;
  is_visible: boolean;
  use_studibox: boolean;
  ticket_price: number;
  ticket_stock: number;
  location: EventLocation;
  categories: string[];
  tags: string[];
  options: EventOption[];
  tarifs: EventTarif[];
  descriptions: EventDescription[];
  external_ticketing_url?: string;
  video_url?: string;
  images: string;
}
