import { BrowserRouter as Router, Routes, Route} from "react-router-dom";
import Chat from "./chat";

function App() {


  return (

      <Router>
        <Routes>
          <Route path="/" element={ <Chat /> } />
          {/*<Route path="/login" element={ <Login/> }/>*/}
        </Routes>
      </Router>

  );
}

export default App;


/*
    <div>
      <h2>Room: {1}</h2>
      <div>
        {messages.map((msg, i) => (
          <div key={i}>{msg}</div>
        ))}
      </div>
      <input
        value={input}
        onChange={(e) => setInput(e.target.value)}
        placeholder="Type message..."
      />
      <button onClick={sendMessage}>Send</button>
    </div>
 */