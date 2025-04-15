import React, { useEffect, useState } from 'react';
import { Ingredient } from '../../domain/Ingredient';
import {
  fetchIngredients,
  softDeleteIngredient,
} from '../../services/ingredientService';

import TableHeaderCell from '../atoms/TableHeaderCell';
import TableCell from '../atoms/TableCell';
import TablePaginationControls from '../molecules/TablePaginationControl';

const IngredientTable: React.FC = () => {
  const [ingredients, setIngredients] = useState<Ingredient[]>([]);
  const [limit, setLimit] = useState(10);
  const [page, setPage] = useState(1);

  const loadData = async () => {
    try {
      const offset = (page - 1) * limit;
      const data = await fetchIngredients(limit, offset);
      setIngredients(data);
    } catch (err) {
      console.error('Failed to fetch ingredients', err);
    }
  };

  useEffect(() => {
    loadData();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [page, limit]);

  const handleDelete = async (uuid: string) => {
    try {
      await softDeleteIngredient(uuid);
      loadData(); // reload after deleting
    } catch (error) {
      console.error('Failed to delete ingredient', error);
    }
  };

  return (
    <div className="overflow-x-auto">
      <table className="min-w-full bg-white border border-gray-200">
        <thead>
          <tr>
            <TableHeaderCell>Name</TableHeaderCell>
            <TableHeaderCell>Allergy</TableHeaderCell>
            <TableHeaderCell>Type</TableHeaderCell>
            <TableHeaderCell>Status</TableHeaderCell>
            <TableHeaderCell>Actions</TableHeaderCell>
          </tr>
        </thead>
        <tbody>
          {ingredients.map((ingredient) => (
            <tr key={ingredient.uuid} className="hover:bg-gray-50">
              <TableCell>{ingredient.name}</TableCell>
              <TableCell>{ingredient.causeAllergy ? 'Yes' : 'No'}</TableCell>
              <TableCell>
                {ingredient.type === 0
                  ? 'None'
                  : ingredient.type === 1
                  ? 'Veggie'
                  : 'Vegan'}
              </TableCell>
              <TableCell>
                {ingredient.status === 1 ? 'Active' : 'Inactive'}
              </TableCell>
              <TableCell>
                <button
                  onClick={() => handleDelete(ingredient.uuid)}
                  className="bg-red-500 hover:bg-red-600 text-white py-1 px-3 rounded"
                >
                  Delete
                </button>
              </TableCell>
            </tr>
          ))}
        </tbody>
      </table>

      <TablePaginationControls
        currentPage={page}
        onPageChange={setPage}
        limit={limit}
        onLimitChange={setLimit}
      />
    </div>
  );
};

export default IngredientTable;
