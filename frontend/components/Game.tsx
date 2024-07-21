import React, { useState, useEffect } from 'react';

const Game: React.FC = () => {
  const [timeLeft, setTimeLeft] = useState(5);
  const [showImage, setShowImage] = useState(true);

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
          <img src="persons/1.png" alt="Person" />
          <p>Time left: {timeLeft} seconds</p>
        </>
      ) : (
        <p>Which attributes match this person?</p>
      )}
    </div>
  );
};

export default Game;
