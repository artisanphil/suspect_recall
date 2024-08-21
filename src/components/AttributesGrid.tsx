import React, { useState, useEffect } from 'react';

interface Item {
  attribute: string;
  clicked: boolean;
  exists: boolean | null;
}

const AttributesGrid: React.FC = () => {    
  const [items, setItems] = useState<Item[]>([]);

  // Prevent click of back button
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

    const clickedAttribute = items[index].attribute;
    console.log('Clicked item:', clickedAttribute);
    console.log('person', personId);
  
    try {
      // Get all items with the class name `grid-item correct` or `grid-item wrong`
      const correctOrWrongItems = items
        .filter((item, idx) => 
          (document.getElementById(`item-${idx}`)?.classList.contains('correct') || 
           document.getElementById(`item-${idx}`)?.classList.contains('wrong')) && 
          idx !== index)
        .map(item => item.attribute);

      // Post these items to the endpoint
      const response = await fetch(`/api/person/${personId}/check-attribute`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ clickedAttribute: clickedAttribute, attributes: correctOrWrongItems }),
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
      console.error(error);
    }    
  };

  return (
    <div className="grid-container">
      {items.map((item, index) => (
        <div
          key={index}
          id={`item-${index}`}
          className={`grid-item ${item.exists === null ? '' : item.exists ? 'correct' : 'wrong'}`}
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
