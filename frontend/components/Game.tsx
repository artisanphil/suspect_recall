import React, { useState, useEffect } from 'react';

const Game = () => {
  const [imageSrc, setImageSrc] = useState('');

  useEffect(() => {
    const fetchImage = async () => {
      try {
        const response = await fetch('/persons/1.png');
        const imageBlob = await response.blob();
        const imageObjectURL = URL.createObjectURL(imageBlob);
        setImageSrc(imageObjectURL);
      } catch (error) {
        console.error('Error fetching the image:', error);
      }
    };

    fetchImage();
  }, []);

  return (
    <div>
      {imageSrc ? <img src={imageSrc} alt="Person" /> : <p>Loading...</p>}
    </div>
  );
};

export default Game;
