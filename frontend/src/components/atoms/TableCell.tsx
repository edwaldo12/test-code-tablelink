import React from 'react';

interface TableCellProps {
  children: React.ReactNode;
}

const TableCell: React.FC<TableCellProps> = ({ children }) => {
  return <td className="py-2 px-4 border-b border-gray-200">{children}</td>;
};

export default TableCell;
