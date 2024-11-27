// src/types/api/event/EventTypes.ts

export interface EventLocation {
  address: string;
  city: string;
  region: string;
  postcode: string;
  country: string;
}

export interface EventOption {
  id?: number; // ID ajouté pour les données récupérées
  title: string;
  description: string;
  price: number;
  stock: number;
}

export interface EventTarif {
  id?: number; // ID ajouté pour les données récupérées
  title: string;
  description: string;
  price: number;
  stock: number;
}

export interface EventDescription {
  title: string;
  description: string;
}

export interface EventForm {
  isOnline: boolean;
  isPublic: boolean;
  title: string;
  subtitle: string;
  categories: string[];
  tags: string[];
  startDate: Date | null;
  startTime: Date | null;
  endDate: Date | null;
  endTime: Date | null;
  description: EventDescription[];
  options: EventOption[];
  tarifs: EventTarif[];
  location: EventLocation;
  images: string[];
  ticketPrice: number;
  ticketStock: number;
  externalTicketingUrl: string;
  video_url: string;
  use_studibox: boolean;
}

export interface EventData {
  id: number;
  title: string;
  subtitle: string;
  start_date: string;
  end_date: string;
  start_time: string;
  end_time: string;
  is_online: boolean;
  is_visible?: boolean; // Ajouté car parfois manquant
  ticket_price: number;
  ticket_stock: number;
  address: string;
  city: string;
  postcode: string;
  region: string;
  country: string;
  categories: string[];
  tags: string[];
  options: EventOption[];
  tarifs: EventTarif[];
  descriptions: EventDescription[];
  external_ticketing_url?: string;
  video_url?: string;
  image?: string[]; // Utilisé lors de la récupération
  use_studibox?: boolean;
  image_urls: string[];
}
// src/types/api/event/EventTypes.ts

export interface EventUpdateData {
  event_id: number;
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
  location: {
    address: string;
    city: string;
    postcode: string;
    region: string;
    country: string;
  };
  options: {
    title: string;
    description: string;
    price: number;
    stock: number;
  }[];
  tarifs: {
    title: string;
    description: string;
    price: number;
    stock: number;
  }[];
  descriptions: {
    title: string;
    description: string;
  }[];
  video_url: string;
  images: string[];
  tags: string[];
  categories: string[];
}
