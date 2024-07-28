import React, { useState, useEffect } from 'react';

const AttributesGrid: React.FC = () => {    
  const [items, setItems] = useState([]);

  useEffect(() => {
    const fetchItems = async () => {
      try {
        const response = await fetch('/api/person/attributes');
        const data = await response.json();
        setItems(data.items);
      } catch (error) {
        console.error('Error fetching items:', error);
      }
    };

    fetchItems();
  }, []);

  const handleClick = (item: any) => {
    console.log('Clicked item:', item);
  };

  return (
    <div className="grid-container">
      {items.map((item, index) => (
        <div key={index} className="grid-item" onClick={() => handleClick(item)}>
          {item}
        </div>
      ))}
    </div>
  );
};

export default AttributesGrid;
