import { create } from 'zustand'
import { createJSONStorage, persist } from 'zustand/middleware'

import { useState, useEffect } from 'react'

export const useStore = <T, F>(
    store: (callback: (state: T) => unknown) => unknown,
    callback: (state: T) => F
) => {
    const result = store(callback) as F
    const [data, setData] = useState<F>()

    useEffect(() => {
        setData(result)
    }, [result])

    return data
}

interface TokenState {
    token: string
    setToken: (token: string) => void
    removeToken: () => void
}

export const useTokenStore = create<TokenState>()(
    persist(
        (set) => ({
            token: "",
            setToken: (newToken) => set(() => ({ token: newToken })),
            removeToken: () => set(() => ({ token: "" })),
        }),
        {
            name: 'token-storage',
            storage: createJSONStorage(() => localStorage),
        }
    )
)