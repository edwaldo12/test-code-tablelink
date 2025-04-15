import { Ingredient } from '../domain/Ingredient';
import axios from 'axios';

const BASE_URL = 'http://localhost:3001/api/ingredients'; 

export async function fetchIngredients(
  limit: number,
  offset: number
): Promise<Ingredient[]> {
  const response = await axios.get<Ingredient[]>(
    `${BASE_URL}?limit=${limit}&offset=${offset}`
  );
  return response.data;
}

export async function createIngredient(data: Partial<Ingredient>): Promise<Ingredient> {
  const response = await axios.post<Ingredient>(BASE_URL, data);
  return response.data;
}

export async function updateIngredient(uuid: string, data: Partial<Ingredient>): Promise<Ingredient> {
  const response = await axios.put<Ingredient>(`${BASE_URL}/${uuid}`, data);
  return response.data;
}

export async function softDeleteIngredient(uuid: string): Promise<void> {
  await axios.delete(`${BASE_URL}/${uuid}`);
}
