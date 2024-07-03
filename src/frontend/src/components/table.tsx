import { useQuery } from "@tanstack/react-query";
import { usePreviewStore } from "../zustand/store";

/*
 * interface Props {
 *   name: string
 * }
 */
export default function Table(props: { name: string, idx: number }) {
  const { name, idx } = props;
  const { show } = usePreviewStore();

  const { data, error, refetch } = useQuery({
    queryKey: [ 'tableData' ],
    enabled: false,
    queryFn: async () => {
      const res = await fetch(`http://0.0.0.0:3000/connections/${idx}/tables/${name}/preview`);

      return res.json();
    }
  });

  const fetchTableData = () => {
    refetch();
    show(data, name);
  }

  return (
    <li>
      <p>{error?.message}</p>
      <button className="" onClick={() => fetchTableData()}> {name} </button>
    </li>
  );
}