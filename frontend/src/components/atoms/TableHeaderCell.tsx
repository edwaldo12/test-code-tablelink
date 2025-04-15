import React from 'react';

interface TableHeaderCellProps {
  children: React.ReactNode;
}

const TableHeaderCell: React.FC<TableHeaderCellProps> = ({ children }) => {
  return (
    <th className="py-2 px-4 border-b border-gray-200 bg-gray-100 text-left">
      {children}
    </th>
  );
};

export default TableHeaderCell;
