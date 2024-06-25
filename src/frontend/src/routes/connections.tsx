import { createFileRoute } from '@tanstack/react-router';
import { useEffect, useState } from "react";
import { useQuery } from "@tanstack/react-query";

const sendData = async (conn: string) => {
  
  const res = await fetch('http://0.0.0.0:3000/connections', {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(conn),
  });

  console.log(res);
}

function Tables(props: { conn: number }) {
  const conn = props.conn;
  const { data, error } = useQuery({
    queryKey: [ 'tables' ],
    queryFn: async () => {
      const res = await fetch(`http://0.0.0.0:3000/connections/${conn}/tables`);

      return res.json();
    }
  });

  return (
    <div>
      <p>
        {error?.message}
      </p>
      <h3> Tables: </h3>
      <ul>
        {data && data?.map((table: string) => (
          <li>{table}</li>
        ))}
      </ul>
    </div>
  );
}

export const Route = createFileRoute('/connections')({
  component: Connections
});

function Connections() {
  const [ connStr, setConnStr ] = useState("");
  const [ showTables, setShowTables] = useState(false);

  const { data, error, refetch } = useQuery({
    queryKey: [ 'connections' ],
    enabled: false,
    queryFn: async () => {
      const res = await fetch('http://0.0.0.0:3000/connections');

      return res.json();
    }
  });

  useEffect(() => {
    refetch();
  }, []);

  return (
    <>
      <div>
        <p>
          {error?.message}
        </p>
        <h1>Here are your tables</h1>
        <div>
          {data && Object.keys(data).map((key) => (
            <>
              <button key={key} onClick={() => setShowTables(prev => !prev)}> Connection {parseInt(key)+1}: {data[key].Name} </button>
              {showTables && (<Tables conn={parseInt(key)} />)}
            </>
          ))}
        </div>
      </div>
      <div>
        <h3> Add a new Connection </h3>
        <form method="post" onSubmit={async (e) => {
          e.preventDefault();
          await sendData(connStr);
          setConnStr("");
          refetch();
        }}>
          <label htmlFor="connStr"> Connection String </label>
          <input type="text" name="connStr" value={connStr} onChange={(e) => setConnStr(e.target.value)}/>
          <button type="submit"> Submit </button>
        </form>
      </div>
    </>
  );
}



