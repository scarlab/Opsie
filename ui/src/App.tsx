import { BrowserRouter, Routes, Route } from "react-router-dom";
import DashboardView from "@/views/workspace/dashboard/DashboardView";
import AuthLayout from "@/layouts/AuthLayout";
import LoginView from "@/views/auth/LoginView";
import WorkspaceLayout from "@/layouts/WorkspaceLayout";
import RootLayout from "./layouts/RootLayout";
import OnboardingView from "@/views/onboarding/OnboardingView";
import NotFoundView from "./views/NotFoundView";

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

          {/* 404 - Not Found */}
          <Route path="*" element={<NotFoundView />} />

        </Route>
      </Routes>
    </BrowserRouter>
  );
}
