import React, { useEffect } from 'react';
import DrawingApp from './components/DrawingApp';
import './reset.css'
import './app.css'

const App: React.FC = () => {
    const players = ["Top", "Michou"];

    useEffect(() => {
        new DrawingApp(players);
    }, []);

    return (
        <div className='wrapper'>
          <div className='scoring'>
            <span id="currentPlayer"></span>
          </div>
          <div className='canvas_drawing_wrapper'>
            <span id="currentWord"></span>

            <canvas id="canvas" width="1000" height="400" style={{ border: '1px solid black' }}></canvas>
            <div className='drawing' id='palette'>
              <div>
              <button className='colors black'></button>
              <button className='colors blue'></button>
              <button className='colors yellow'></button>
              <button className='colors red'></button>
              </div>
              <div>
              <button className='colors white'></button>
              <button className='colors red'></button>
              <button className='colors red'></button>
              <button className='colors red'></button>
              </div>
              <button id="clear">Clear</button>
            </div>
          </div>
          <div className='chat'>
            <form id="guess">
              <input type="text" id="guessInput" placeholder="Enter your guess" />
              <button id="guess" type="submit">Guess</button>
            </form>
          </div>
        </div>
    );
}

export default App;
