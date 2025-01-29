import React, { useState, useEffect } from 'react';
import './AttributesGrid.css';
import { BASE_URL } from '../App';
import Swal from 'sweetalert2'
import withReactContent from 'sweetalert2-react-content'

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
  loadNextSuspect: () => void;
  loadSameSuspect: () => void;
};

const MySwal = withReactContent(Swal)

const AttributesGrid: React.FC<AttributesGridProps> = ({ person, loadNextSuspect, loadSameSuspect }) => {
  const [items, setItems] = useState<Item[]>([]);
  const [finished, setFinished] = useState(false)
  const [mistakes, setMistakes] = useState(0)

  // Prevent click of back button
  useEffect(() => {
    window.scrollTo(0, 0);
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
        setItems(
          data.items.map((item) => ({
            attribute: item,
            clicked: false,
            exists: data.correct && data.correct.includes(item)
              ? true 
              : data.wrong && data.wrong.includes(item)
              ? false 
              : null, 
          }))
        );
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

      if (!result.exists) {
        MySwal.fire({
          title: 'Wrong selection!',
          text: '"' + clickedAttribute + '" is not a match.',          
          showCancelButton: false,
          confirmButtonText: 'Try Again',
          footer: (
            <a
              href="#"
              onClick={(e) => {
                e.preventDefault(); 
                loadNextSuspect();
                Swal.close(); 
              }}
            >
              Load another suspect
            </a>
          )
        }).then(() => {
          loadSameSuspect()
        });
      }

      if (result.finished === true) {
        window.scrollTo(0, document.body.scrollHeight + 20);
      }

      setFinished(result.finished);
      setMistakes(result.mistakes);
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div>
      <div className="grid-container">
        {items.map((item, index) => (
          <div
            key={index}
            id={`item-${index}`}
            className={`grid-item ${item.exists === null ? '' : item.exists ? 'correct' : 'wrong'}`}
            {...(item.exists === null && {
              onClick: () => handleClick(index, person.id),
            })}
            style={{ cursor: item.exists != null || item.clicked ? 'default' : 'pointer' }}
          >
            {item.attribute}
          </div>
        ))}

      </div>
      {finished ? (
        <div class="thank-you-section">
          <h2>Thank you for your valuable assistance in identifying the suspect.</h2>
          <p>
            {mistakes > 0
              ? "We have a close match in our database."
              : "We have an exact match in our database!"
            }
            {
              person.noMore ? (
                <div>
                  <div>No more suspects available!</div>
                  <div>
                    <button onClick={() => window.location.href = '/'}>Start over</button>
                  </div>                
                </div>
              ) : (
                <div>
                  <button onClick={() => loadNextSuspect()} >Show next suspect</button>
                </div>
              )
            }
          </p>
        </div>
      ) : (
        <div>You need to select all the attributes that the suspect has before you can continue.</div>
      )}
    </div>
  );
};

export default AttributesGrid;
