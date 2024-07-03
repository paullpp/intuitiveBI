import { useQuery } from "@tanstack/react-query";
import { useState } from "react";
import Preview from "./dataPreview";

/*
 * interface Props {
 *   name: string
 * }
 */
export default function Table(props: { name: string }) {
  const { name } = props;
  const [ showData, setShowData ] = useState(false);

  const { data, error, refetch } = useQuery({
    queryKey: [ 'tableData' ],
    enabled: false,
    queryFn: async () => {
      // TODO: remove conn hardcoding
      const res = await fetch(`http://0.0.0.0:3000/connections/${0}/tables/${name}/preview`);

      return res.json();
    }
  });

  const fetchTableData = (stateFunc: React.Dispatch<React.SetStateAction<boolean>>) => {
    refetch();
    stateFunc(prev => !prev);
  }

  return (
    <>
      <p>{error?.message}</p>
      <button className="btn rounded w-52" onClick={() => fetchTableData(setShowData)}> {name} </button>
      {showData && (
        <Preview data={data}/>
      )}
    </>
  );
}