import { useState } from "react";

const sendData = async (conn: string) => {
  
  const res = await fetch('http://0.0.0.0:3000/connect', {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(conn),
  });

  console.log(res);
}

export default function App() {
  const [ connStr, setConnStr ] = useState("");

  return (
    <div>
      <form method="post" onSubmit={async (e) => {
        e.preventDefault();
        await sendData(connStr)
      }}>
        <label htmlFor="connStr"> Connection String </label>
        <input type="text" name="connStr" value={connStr} onChange={(e) => setConnStr(e.target.value)}/>
        <button type="submit"> Submit </button>
      </form>
    </div>
  );
}
