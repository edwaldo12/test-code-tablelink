export interface Ingredient {
  uuid: string;
  name: string;
  causeAllergy: boolean;
  type: number;
  status: number;
  createdAt?: Date;
  updatedAt?: Date;
  deletedAt?: Date | null;
}
