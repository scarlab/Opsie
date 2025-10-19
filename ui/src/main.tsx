import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'
import { ThemeProvider } from '@/components/theme'
import { Toaster } from 'sonner'
import CsReduxProvider from '@/cs-redux/redux-provider'


createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <ThemeProvider>
      <CsReduxProvider>
        <App />
      </CsReduxProvider>
      <Toaster richColors />
    </ThemeProvider>
  </StrictMode>,
)
