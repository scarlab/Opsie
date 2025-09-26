import { BrowserRouter, Routes, Route } from "react-router-dom";
import DashboardView from "@/views/DashboardView";
import AuthLayout from "@/layouts/AuthLayout";
import LoginView from "@/views/auth/login/LoginView";
import DashboardLayout from "@/layouts/DashboardLayout";

export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        {/* Auth routes */}
        <Route element={<AuthLayout />}>
          <Route path="/login" element={<LoginView />} />
        </Route>

        {/* Dashboard routes */}
        <Route element={<DashboardLayout />}>
          <Route path="/dashboard" element={<DashboardView />} />
        </Route>
      </Routes>
    </BrowserRouter>
  )
}

