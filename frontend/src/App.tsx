// src/App.tsx
import React, { useState } from 'react';
import IngredientManagementPage from './components/Pages/IngredientManagementPage';
import ItemManagementPage from './components/Pages/ItemManagementPage';

function App() {
  const [activeTab, setActiveTab] = useState<'ingredients' | 'items'>(
    'ingredients'
  );

  return (
    <div>
      <nav className="flex space-x-4 bg-gray-100 p-4">
        <button
          onClick={() => setActiveTab('ingredients')}
          className={`py-1 px-3 ${
            activeTab === 'ingredients'
              ? 'bg-blue-500 text-white'
              : 'bg-gray-300'
          } rounded`}
        >
          Ingredients
        </button>
        <button
          onClick={() => setActiveTab('items')}
          className={`py-1 px-3 ${
            activeTab === 'items' ? 'bg-blue-500 text-white' : 'bg-gray-300'
          } rounded`}
        >
          Items
        </button>
      </nav>

      {activeTab === 'ingredients' && <IngredientManagementPage />}
      {activeTab === 'items' && <ItemManagementPage />}
    </div>
  );
}

export default App;
