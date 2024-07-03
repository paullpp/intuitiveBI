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
    <li>
      <button className="" key={idx} onClick={() => setShowTables(prev => !prev)}> {idx+1}: {connection.Name.split("/")[1]} </button>
      <ul>
        {showTables && (connection.Tables ? connection.Tables.map((table) => (
          <Table name={table} idx={idx} key={table}/>
        )) : <li>No tables available</li>)}
      </ul>
    </li>
  );
}