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

  const handleClick = async (attribute: string, personId: number) => {
    console.log('Clicked item:', attribute);
    console.log('person', personId)
  
    try {
      const response = await fetch(`/api/person/${personId}/check-attribute`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ attribute }),
      });
  
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
  
      const result = await response.json();
      console.log(`Attribute exists: ${result.exists}`);
    } catch (error) {
      console.error('Error checking attribute:', error);
    }
  };

  return (
    <div className="grid-container">
      {items.map((item, index) => (
        <div key={index} className="grid-item" onClick={() => handleClick(item, 1)}>
          {item}
        </div>
      ))}
    </div>
  );
};

export default AttributesGrid;
