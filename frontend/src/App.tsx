import { useEffect, useState } from "react";
import "./App.css";
import { Recognition } from "./Recognition";

function App() {
  const [ws, setWs] = useState<WebSocket>();
  useEffect(() => {
    const ws = new WebSocket("ws://localhost:8080/ws");
    setWs(ws);
    return () => {
      ws.close();
    };
  }, []);

  return <>{ws && <Recognition ws={ws} />}</>;
}

export default App;
