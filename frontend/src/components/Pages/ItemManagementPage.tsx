import React from 'react';
import ItemTable from '../organisms/ItemTable';

const ItemManagementPage: React.FC = () => {
  return (
    <div className="p-6">
      <h1 className="text-2xl font-bold mb-4">Item Management</h1>
      <ItemTable />
    </div>
  );
};

export default ItemManagementPage;
