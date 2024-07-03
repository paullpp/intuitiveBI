import { usePreviewStore } from "../zustand/store";

/*
 * interface Props {
 *   name: string
 * }
 */
export default function Table(props: { name: string, idx: number }) {
  const { name, idx } = props;
  const show = usePreviewStore((state) => state.show);
  const curName = usePreviewStore((state) => state.name);

  const fetchTableData = async () => {
    if (name === curName) {
      // return early to prevent re-querying the same table
      return
    }
    const res = await fetch(`http://0.0.0.0:3000/connections/${idx}/tables/${name}/preview`);
    const data = await res.json();
    show(data, name);
  }

  return (
    <li>
      <button className="" onClick={async () => await fetchTableData()}> {name} </button>
    </li>
  );
}