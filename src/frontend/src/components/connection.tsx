import { useState } from "react";
import Table from "./table";
/*
 * interface Connection {
 *   Name: string,
 *   Tables: string[]
 * }
 * interface Props {
 *   connection: Connection,
 *   key: number
 * }
 */
export default function Connection(props: { connection: { Name: string, Tables: string[] }, idx: number }) {
  const { connection, idx } = props;
  const [ showTables, setShowTables ] = useState(false);

  return (
    <>
      <button className="btn rounded w-64" key={idx} onClick={() => setShowTables(prev => !prev)}> {idx+1}: {connection.Name.split("/")[1]} </button>
      <div className="flex flex-wrap gap-5">
        {showTables && connection.Tables.map((table) => (
          <Table name={table} key={table}/>
        ))}
      </div>
    </>
  );
}