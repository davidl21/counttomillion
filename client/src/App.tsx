import React, { useState } from "react";

function App() {
  const [inputValue, setInputValue] = useState("");
  const [numbers, setNumbers] = useState([3, 2, 1]);
  const [currentNumber, setCurrentNumber] = useState(3);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (parseInt(inputValue) === currentNumber + 1) {
      setNumbers((prev) => [parseInt(inputValue), ...prev.slice(0, 3)]);
      setCurrentNumber((prev) => prev + 1);
      setInputValue("");
    }
  };

  return (
    <div className="flex flex-col items-center justify-center min-h-screen">
      <form onSubmit={handleSubmit} className="flex flex-col items-center mb-8">
        <input
          type="number"
          value={inputValue}
          onChange={(e) => setInputValue(e.target.value)}
          placeholder="Enter the next number..."
          className="p-2 text-xl border rounded-lg w-48 text-center"
        />
      </form>

      <div className="flex flex-col items-center gap-4">
        {numbers.map((num, index) => (
          <div
            key={num}
            className="number text-center"
            style={{
              fontSize: `${Math.max(16, 120 - index * 30)}px`,
              opacity: 1 - index * 0.2,
            }}
          >
            {num}
          </div>
        ))}
      </div>
    </div>
  );
}

export default App;
