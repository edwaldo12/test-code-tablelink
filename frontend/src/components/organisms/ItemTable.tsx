import React, { useEffect, useState } from 'react';
import { Item } from '../../domain/Item';
// import { fetchItems, softDeleteItem } from '../../grpc/itemService';
import TableHeaderCell from '../atoms/TableHeaderCell';
import TableCell from '../atoms/TableCell';
import TablePaginationControls from '../molecules/TablePaginationControl';

const ItemTable: React.FC = () => {
  const [items, setItems] = useState<Item[]>([]);
  const [limit, setLimit] = useState<number>(10);
  const [page, setPage] = useState<number>(1);

  const loadData = async () => {
    try {
      const offset = (page - 1) * limit;
      // const data = await fetchItems(limit, offset);
      // setItems(data);
    } catch (error) {
      console.error('Failed to fetch items:', error);
    }
  };

  useEffect(() => {
    loadData();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [page, limit]);

  const handleDelete = async (uuid: string) => {
    try {
      // await softDeleteItem(uuid);
      loadData(); // Reload data after deletion
    } catch (error) {
      console.error('Failed to delete item:', error);
    }
  };

  return (
    <div className="overflow-x-auto">
      <table className="min-w-full bg-white border border-gray-200">
        <thead>
          <tr>
            <TableHeaderCell>Name</TableHeaderCell>
            <TableHeaderCell>Price</TableHeaderCell>
            <TableHeaderCell>Status</TableHeaderCell>
            <TableHeaderCell>Actions</TableHeaderCell>
          </tr>
        </thead>
        <tbody>
          {items.map((item) => (
            <tr key={item.uuid} className="hover:bg-gray-50">
              <TableCell>{item.name}</TableCell>
              <TableCell>${item.price.toFixed(2)}</TableCell>
              <TableCell>{item.status === 1 ? 'Active' : 'Inactive'}</TableCell>
              <TableCell>
                <button
                  onClick={() => handleDelete(item.uuid)}
                  className="bg-red-500 hover:bg-red-600 text-white py-1 px-3 rounded"
                >
                  Delete
                </button>
                <button
                  className="bg-blue-500 hover:bg-blue-600 text-white py-1 px-3 rounded ml-2"
                  // You can add an onClick handler for editing functionality here
                >
                  Edit
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

export default ItemTable;
