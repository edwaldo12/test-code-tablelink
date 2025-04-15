import React from 'react';
import IngredientTable from '../organisms/IngredientTable';

const IngredientManagementPage: React.FC = () => {
  return (
    <div className="p-6">
      <h1 className="text-2xl font-bold mb-4">Ingredient Management</h1>
      <IngredientTable />
    </div>
  );
};

export default IngredientManagementPage;
