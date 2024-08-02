import React, { useEffect } from 'react';
import DrawingApp from './components/DrawingApp';
import './reset.css'
import './app.css'

const App: React.FC = () => {

    useEffect(() => {
        new DrawingApp();
    }, []);

    return (
        <div className='wrapper'>
          <div className='scoring'>

          </div>
          <div className='canvas_drawing_wrapper'>
            <canvas id="canvas" width="1000" height="400" style={{ border: '1px solid black' }}></canvas>
            <div className='drawing'>
              <button id="clear">Clear</button>
            </div>
          </div>
          <div className='chat'>

          </div>
        </div>
    );
}

export default App;
