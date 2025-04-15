import axios from 'axios';
import { Item } from '../domain/Item';

const BASE_URL = 'http://localhost:3001/api/items';

export async function fetchItems(
  limit: number,
  offset: number
): Promise<Item[]> {
  const response = await axios.get<Item[]>(
    `${BASE_URL}?limit=${limit}&offset=${offset}`
  );
  return response.data;
}

export async function createItem(data: Partial<Item>): Promise<Item> {
  const response = await axios.post<Item>(BASE_URL, data);
  return response.data;
}

export async function updateItem(
  uuid: string,
  data: Partial<Item>
): Promise<Item> {
  const response = await axios.put<Item>(`${BASE_URL}/${uuid}`, data);
  return response.data;
}

export async function softDeleteItem(uuid: string): Promise<void> {
  await axios.delete(`${BASE_URL}/${uuid}`);
}
