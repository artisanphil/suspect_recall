import React, { useState, useEffect } from 'react';
import AttributesGrid from './AttributesGrid';
import { BASE_URL } from '../App';

type Person = {
  id: number;
};

const Game: React.FC = () => {
  const [timeLeft, setTimeLeft] = useState(5);
  const [showImage, setShowImage] = useState(true);
  const [person, setPerson] = useState<Person | null>(null);

  useEffect(() => {
    const fetchPerson = async () => {
      try {
        const response = await fetch('/api/person');
        const data: Person = await response.json();
        setPerson(data);
      } catch (error) {
        console.error('Error fetching person data:', error);
      }
    };

    fetchPerson();
  }, [])

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
    <div style={{ textAlign: 'center', marginTop: '50px' }}>
      {showImage ? (
        <>
           {person ? (
            <img src={`/persons/${person.id}.png`} alt="Person" />
          ) : (
            <p>Loading...</p>
          )}
          <p>Time left: {timeLeft} seconds</p>
        </>
      ) : (
        <div>
          <p>Which attributes match this person?</p>
          <AttributesGrid person={person} />
        </div>
      )}
    </div>
  );
};

export default Game;
