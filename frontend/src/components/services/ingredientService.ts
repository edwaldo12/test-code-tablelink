import { IngredientClient } from '../../grpc/client';
import {
  DeleteIngredientRequest,
  ListIngredientsRequest,
} from '../../grpc/service';

export const fetchIngredients = async (limit: number, offset: number) => {
  try {
    const request = ListIngredientsRequest.create({ limit, offset });
    const response = await IngredientClient.listIngredients(request);
    console.log(response);
    return response;
  } catch (error) {
    console.error('Error fetching ingredients:', error);
    throw error;
  }
};

export const softDeleteIngredient = async (uuid: string) => {
  try {
    const request = DeleteIngredientRequest.create({ uuid });
    await IngredientClient.deleteIngredient(request);
  } catch (error) {
    console.error('Error deleting ingredient:', error);
    throw error;
  }
};
