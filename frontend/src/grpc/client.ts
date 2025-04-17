import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';
import { IngredientServiceClient, ItemServiceClient } from './service.client';

const transport = new GrpcWebFetchTransport({
  baseUrl: 'http://127.0.0.1:8080',
});

export const IngredientClient = new IngredientServiceClient(transport);
export const ItemClient = new ItemServiceClient(transport);
