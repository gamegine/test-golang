const WebSocket = require("ws");

const ws = new WebSocket("ws://localhost:8080/ws");

ws.on("error", console.error);

ws.on("open", () => {
  console.log("WebSocket Client Connected");
  setInterval(() => {
    ws.send("something");
  }, 1000);
});

ws.on("message", (data) => console.log("received: %s", data));
