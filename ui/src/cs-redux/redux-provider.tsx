import React, { useRef } from 'react'
import { Provider } from 'react-redux'
import { CsCreateStore } from './store'


export default function CsReduxProvider({ children }: Readonly<{ children: React.ReactNode }>) {
    const storeRef = useRef<ReturnType<typeof CsCreateStore> | null>(null)

    if (!storeRef.current) {
        storeRef.current = CsCreateStore()
    }

    return <Provider store={storeRef.current}>{children}</Provider>
}