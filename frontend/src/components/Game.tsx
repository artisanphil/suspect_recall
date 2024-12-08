import React, { useState, useEffect } from 'react';
import AttributesGrid from './AttributesGrid';
import './Game.css';

type Person = {
  id: number;
  noMore: boolean;
};

const Game: React.FC = () => {
  const [showImage, setShowImage] = useState(true);
  const [timeLeft, setTimeLeft] = useState(5);
  const [reloadTrigger, setReloadTrigger] = useState(false);
  const [person, setPerson] = useState<Person | null>(null);

  const reloadParent = () => {
    // Update to force a re-render of the component
    if (person) {
      setShowImage(true);
    }
    setTimeLeft(5);
    setReloadTrigger(prev => !prev);
  };


  useEffect(() => {
    const fetchPerson = async () => {
      try {
        const response = await fetch('/api/person', { credentials: 'include' });
        const data: Person = await response.json();
        setPerson(data);
      } catch (error) {
        console.error('Error fetching person data:', error);
      }
    };

    fetchPerson();
  }, [reloadTrigger])

  useEffect(() => {
    if (timeLeft > 0) {
      const timer = setInterval(() => {
        setTimeLeft(timeLeft - 1);
      }, 1000);
      return () => clearInterval(timer);
    } else {
      setShowImage(false);
    }
  }, [timeLeft]);

  return (
    <div className="App">
      <header className="App-header min-h-screen flex flex-col">

        <div className="text-center mt-12 max-w-full max-h-full">
          {showImage ? (
            <>
              {person ? (
                <div>
                  <h1>Take a good look at the suspect!</h1>
                    <div className="flex justify-center">
                      <img className="suspect" src={`/persons/${person.id}.png`} alt="Person" />
                    </div>                    
                </div>
              ) : (
                <p>Loading...</p>
              )}
              <p>Time left: {timeLeft} seconds</p>
            </>
          ) : (
            <div>
              <h1>Which attributes match this person?</h1>
              <AttributesGrid person={person} onReload={reloadParent} />
            </div>
          )}
        </div>
      </header>
    </div>
  );
};

export default Game;
