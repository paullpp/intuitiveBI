import { create } from 'zustand'
import { devtools, persist } from 'zustand/middleware'

interface PreviewState {
  name: string,
  data: object[],
  show: (data: object[], name: string) => void
}

export const usePreviewStore = create<PreviewState>()(
  devtools(
    persist(
      (set) => ({
        name: "",
        data: [],
        show: (data, name) => set((state) => ({ data: state.data = data, name: state.name = name})),
      }),
      {
        name: 'preview-store',
      },
    ),
  ),
)