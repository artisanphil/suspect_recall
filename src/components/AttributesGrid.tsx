import React, { useState, useEffect } from 'react';

interface Item {
  attribute: string;
  clicked: boolean;
  exists: boolean | null;
}

const AttributesGrid: React.FC = () => {    
  const [items, setItems] = useState<Item[]>([]);

  //prevent click of back button
  useEffect(() => {
    window.history.pushState(null, '', window.location.href);

    const handlePopState = () => {
      window.history.pushState(null, '', window.location.href);
    };

    window.addEventListener('popstate', handlePopState);

    return () => {
      window.removeEventListener('popstate', handlePopState);
    };
  }, []);

  useEffect(() => {
    const fetchItems = async () => {
      try {
        const response = await fetch('/api/person/attributes');
        const data = await response.json();
        setItems(data.items.map((item: string) => ({ attribute: item, clicked: false, exists: null })));
      } catch (error) {
        console.error('Error fetching items:', error);
      }
    };

    fetchItems();
  }, []);

  const handleClick = async (index: number, personId: number) => {
    if (items[index].clicked) return;

    const attribute = items[index].attribute;
    console.log('Clicked item:', attribute);
    console.log('person', personId);
  
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

      setItems(prevItems => {
        const newItems = [...prevItems];
        newItems[index] = { ...newItems[index], clicked: true, exists: result.exists };
        return newItems;
      });
    } catch (error) {
      console.error('Error checking attribute:', error);
    }
  };

  return (
    <div className="grid-container">
      {items.map((item, index) => (
        <div 
          key={index} 
          className={`grid-item ${item.clicked ? (item.exists ? 'green' : 'red') : ''}`} 
          onClick={() => handleClick(index, 1)}
          style={{ cursor: item.clicked ? 'default' : 'pointer' }}
        >
          {item.attribute}
        </div>
      ))}
    </div>
  );
};

export default AttributesGrid;