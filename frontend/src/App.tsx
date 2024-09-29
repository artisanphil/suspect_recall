import React from 'react';
import Game from './components/Game';

export const BASE_URL = process.env.REACT_APP_MODE === "development" ? "http://localhost:8080/api" : "/api";

const App: React.FC = () => {
  return (
    <div className="App">
      <header className="App-header">
        <Game />
      </header>
    </div>
  );
};

export default App;
