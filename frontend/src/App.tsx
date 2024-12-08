import React from 'react';
import { BrowserRouter as Router, Routes, Route, useNavigate } from "react-router-dom";
import Game from './components/Game';
import './App.css';

export const BASE_URL = process.env.REACT_APP_MODE === "development" ? "http://localhost:8080/api" : "/api";

const Home: React.FC = () => {
  const navigate = useNavigate();
  return (
    <div className="App">
      <header className="min-h-screen flex flex-col">
        <div className="bg-gray-100 flex flex-col items-center justify-center p-4 flex-1">
          <div className="text-center flex flex-col items-center space-y-6 w-full">
            <h1 className="text-6xl font-wild-west text-gray-800">Suspect Recall</h1>
            <img
              src="/suspects.png"
              alt="Suspect Recall"
              className="h-auto w-auto max-h-[50vh] object-contain"
            />
            <p className="bg-blue-100 border border-blue-300 text-blue-800 p-4 rounded-md shadow-md max-w-xl">
              In this game, you’ll see an image of a suspect for a few seconds.
              Your challenge is to remember as many attributes as possible. Test
              your memory and detective skills!
            </p>
            <button
              className="bg-blue-600 text-white py-3 px-6 rounded-md text-lg hover:bg-blue-700 transition duration-300"
              onClick={() => navigate("/game")}
            >
              Start Game
            </button>
          </div>
        </div>
      </header>
    </div>
  );
};

const App: React.FC = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/game" element={<Game />} />
      </Routes>
    </Router>
  );
};

export default App;