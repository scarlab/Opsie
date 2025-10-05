import { BrowserRouter, Routes, Route } from "react-router-dom";
import DashboardView from "@/views/workspace/dashboard/DashboardView";
import AuthLayout from "@/layouts/AuthLayout";
import LoginView from "@/views/auth/login/LoginView";
import WorkspaceLayout from "@/layouts/WorkspaceLayout";
import RootLayout from "./layouts/RootLayout";
import OnboardingView from "@/views/onboarding/OnboardingView";

export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route element={<RootLayout />}>

          {/* Onboarding routes */}
          <Route path="onboarding" element={<OnboardingView />} />


          {/* Auth routes */}
          <Route path="auth" element={<AuthLayout />}>
            <Route index element={<LoginView />} /> {/* /login */}
            <Route path="login" element={<LoginView />} /> {/* /login */}
          </Route>

          {/* Workspace routes */}
          <Route element={<WorkspaceLayout />}>
            <Route path="dashboard" element={<DashboardView />} /> {/* /dashboard */}
          </Route>

        </Route>
      </Routes>
    </BrowserRouter>
  );
}
