import React, { useState, useEffect } from 'react';
import './AttributesGrid.css';
import { BASE_URL } from '../App';

interface Item {
  attribute: string;
  clicked: boolean;
  exists: boolean | null;
}

type Person = {
  id: number;
  noMore: boolean;
};

type AttributesGridProps = {
  person: Person;
  onReload: boolean
};

const AttributesGrid: React.FC<AttributesGridProps> = ({ person, onReload }) => {    
  const [items, setItems] = useState<Item[]>([]);
  const [finished, setFinished] = useState(false)
  const [mistakes, setMistakes] = useState(0)

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
        let personId = person.id;
        const response = await fetch(
          BASE_URL + `/person/${personId}/attributes`, {
            credentials: 'include',
          }          
        );
        const data = await response.json();
        setItems(data.items.map((item: string) => ({ attribute: item, clicked: false, exists: null })));
      } catch (error) {
        console.error('Error fetching items:', error);
      }
    };

    fetchItems();
  }, [person]);

  const handleClick = async (index: number, personId: number) => {
    if (items[index].clicked) return;

    const clickedAttribute = items[index].attribute;
  
    try {
      // Get all items with the class name `grid-item correct` or `grid-item wrong`
      const correctOrWrongItems = items
        .filter((item, idx) => 
          (document.getElementById(`item-${idx}`)?.classList.contains('correct') || 
           document.getElementById(`item-${idx}`)?.classList.contains('wrong')) && 
          idx !== index)
        .map(item => item.attribute);

      // Post these items to the endpoint
      const response = await fetch(BASE_URL + `/person/${personId}/check-attribute`, {
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

      setItems(prevItems => {
        const newItems = [...prevItems];
        newItems[index] = { ...newItems[index], clicked: true, exists: result.exists };
        return newItems;
      });

      setFinished(result.finished);
      setMistakes(result.mistakes);
    } catch (error) {
      console.error(error);
    }    
  };

  const nextSuspect = () => {
    console.log("Show next suspect")
  }

  return (
    <div>
      <div className="grid-container">
        {items.map((item, index) => (
          <div
            key={index}
            id={`item-${index}`}
            className={`grid-item ${item.exists === null ? '' : item.exists ? 'correct' : 'wrong'}`}
            onClick={() => handleClick(index, person.id)}
            style={{ cursor: item.clicked ? 'default' : 'pointer' }}
          >
            {item.attribute}
          </div>
        ))}
        
      </div>
      {finished && (
        <div class="thank-you-section">
        <h2>Thank you for your valuable assistance in identifying the suspect.</h2>
        <p>
        {mistakes > 0 
          ? "We have a close match in our database." 
          : "We have an exact match in our database!"
        }
        {
          person.noMore ? (
            <div>No more suspects available!</div>           
          ) : (
            <div>
              <button onClick={onReload}>Show next suspect</button>
            </div>
          )
        }
        </p>
        </div>
      )}      
    </div>
  );
};

export default AttributesGrid;
