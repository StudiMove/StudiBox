export interface Ticket {
  id: number;
  userId: number;
  eventId: number;
  issueDate: string;
  ticketCode: string;
  status: 'valid' | 'cancelled' | 'used';
  createdAt: string;
  updatedAt: string;
}
