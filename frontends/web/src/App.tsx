import { BrowserRouter as Router, Routes, Route} from "react-router-dom";
import Chat from "./Chat";
import ProtectedRoute from "./providers/ProtectedRoute";
import Groups from "./Groups";

function App() {


  return (

      <Router>
        <Groups/>
        <Routes>
            <Route path="/" element={ 
              <ProtectedRoute>
                <Chat /> 
              </ProtectedRoute>
              } />

            <Route path="/login" element={ <div>Login</div> }/>
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