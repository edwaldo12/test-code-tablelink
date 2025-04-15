export interface Item {
  uuid: string;
  name: string;
  price: number;
  status: number;
  createdAt?: Date;
  updatedAt?: Date;
  deletedAt?: Date | null;
}
