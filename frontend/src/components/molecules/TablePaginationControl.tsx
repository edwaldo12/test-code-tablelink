import React from 'react';

interface TablePaginationControlsProps {
  currentPage: number;
  onPageChange: (page: number) => void;
  limit: number;
  onLimitChange: (limit: number) => void;
}

const TablePaginationControls: React.FC<TablePaginationControlsProps> = ({
  currentPage,
  onPageChange,
  limit,
  onLimitChange,
}) => {
  return (
    <div className="flex items-center mt-4 gap-4">
      <label className="text-sm">
        Show
        <select
          className="mx-2 border rounded px-2 py-1"
          value={limit}
          onChange={(e) => {
            onLimitChange(parseInt(e.target.value, 10));
            onPageChange(1);
          }}
        >
          <option value={10}>10</option>
          <option value={20}>20</option>
          <option value={50}>50</option>
        </select>
        items
      </label>
      <div className="flex gap-2">
        <button
          className="border rounded px-3 py-1 disabled:opacity-50"
          onClick={() => onPageChange(currentPage - 1)}
          disabled={currentPage <= 1}
        >
          Prev
        </button>
        <span className="px-2">Page {currentPage}</span>
        <button
          className="border rounded px-3 py-1"
          onClick={() => onPageChange(currentPage + 1)}
        >
          Next
        </button>
      </div>
    </div>
  );
};

export default TablePaginationControls;
