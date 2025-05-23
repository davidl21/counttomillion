import "@testing-library/jest-dom";
import { render, screen, fireEvent, act } from "@testing-library/react";
import WS from "jest-websocket-mock";
import Chat from "../Chat";

describe("Chat Component Integration Tests", () => {
  let ws: WS;

  beforeEach(() => {
    ws = new WS("ws://localhost:8080/ws");
  });

  afterEach(() => {
    WS.clean();
  });

  it("establishes WebSocket connection and handles messages", async () => {
    render(<Chat />);

    // Wait for client connection
    await ws.connected;

    // Test sending message
    const input = screen.getByPlaceholderText("message");
    const sendButton = screen.getByText("Send!!");

    fireEvent.change(input, { target: { value: "test message" } });
    fireEvent.click(sendButton);

    // Verify message was sent to server
    await expect(ws).toReceiveMessage(
      JSON.stringify({ message: "test message" })
    );

    // Simulate server response
    act(() => {
      ws.send(JSON.stringify({ message: "test message" }));
    });

    // Verify message appears in UI
    expect(await screen.findByText("test message")).toBeInTheDocument();
  });

  it("handles connection errors gracefully", async () => {
    render(<Chat />);

    await ws.connected;

    act(() => {
      ws.error();
    });
  });
});
