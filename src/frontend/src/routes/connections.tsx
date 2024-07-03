import { createFileRoute } from '@tanstack/react-router';
import { useEffect, useState } from "react";
import { useQuery } from "@tanstack/react-query";
import Connection from '../components/connection';
import Preview from '../components/dataPreview';

export const Route = createFileRoute('/connections')({
  component: Connections
});

function Connections() {
  const [ connStr, setConnStr ] = useState("");
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
      <div className='flex flex-row gap-5 m-5'>
        <div className='flex flex-col gap-10'>
          <p className='font-bold text-2xl'>Connections</p>
          {error && (
            <p>{error.message}</p>
          )}
          <ul className="menu bg-base-200 rounded-box w-56">
            {data && Object.keys(data).map((key) => (
              <Connection connection={data[key]} idx={parseInt(key)} key={key}/>
            ))}
          </ul>
        </div>
        <div className='block'>
          <Preview />
        </div>
      </div>
      <div className='flex flex-col gap-10 m-5'>
        <p className='font-bold text-2xl'>Add a new connection</p>
        <form method="post" onSubmit={async (e) => {
          e.preventDefault();
          await sendData(connStr);
          setConnStr("");
          refetch();
        }}>
          <label htmlFor="connStr"> Connection String </label>
          <input className="input input-bordered rounded" type="text" name="connStr" value={connStr} onChange={(e) => setConnStr(e.target.value)}/>
          <button className="btn rounded" type="submit"> Submit </button>
        </form>
      </div>
    </>
  );
}

const sendData = async (conn: string) => {
  await fetch('http://0.0.0.0:3000/connections', {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(conn),
  });
}
