import logo from "./logo.svg";
import "./App.css";

function App() {
  let socket = new WebSocket("ws://localhost:8080/ws");
  console.log("Attempting Websocket Connection");

  socket.onopen = () => {
    console.log("Successfully connected");
    socket.send("Hi from the client");
  };

  socket.onclose = (event) => {
    console.log("Socket closed connection", event);
  };

  socket.onerror = (error) => {
    console.log("Socket error: ", error);
  };

  socket.onmessage = (msg) => {
    console.log(msg);
  };

  
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;
