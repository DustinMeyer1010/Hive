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
